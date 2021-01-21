package futures

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	APIURL = "https://futures.kraken.com/derivatives"
)

type RestApiClient struct {
	key    string
	secret string
	client *http.Client
}

func NewRestApi(key, secret string) *RestApiClient {
	return &RestApiClient{
		key:    key,
		secret: secret,
		client: http.DefaultClient,
	}
}

func NewRestApiWithClient(key, secret string, httpClient *http.Client) *RestApiClient {
	kraken := NewRestApi(key, secret)
	kraken.client = httpClient
	return kraken
}

func (api *RestApiClient) Instruments() (interface{}, error) {
	resp, err := api.queryPublic("GET", "instruments", url.Values{
		//"asset":  {asset},
		//"key":    {key},
		//"amount": {amount.String()},
	})
	log.Println(string(resp))
	if err != nil {
		return nil, err
	}
	return 1, nil
}

/*
url.Values{
		//"asset":  {asset},
		//"key":    {key},
		//"amount": {amount.String()},
	}
*/

func (api *RestApiClient) Accounts() (interface{}, error) {
	resp, err := api.queryPrivate("GET", "/api/v3/accounts", url.Values{})
	if err != nil {
		return nil, err
	}

	var jsonData FutAccountsResult
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request! (%s)", jsonData.Error)
	}
	return jsonData, nil
}

func (api *RestApiClient) queryPublic(method string, endpoint string, values url.Values) ([]byte, error) {
	url := fmt.Sprintf("%s%s", APIURL, endpoint)
	resp, err := api.request(method, url, values, nil)
	return resp, err
}

func (api *RestApiClient) queryPrivate(method string, endpoint string, values url.Values) ([]byte, error) {
	url := fmt.Sprintf("%s%s", APIURL, endpoint)
	secret, _ := base64.StdEncoding.DecodeString(api.secret)
	nonce := time.Now().UnixNano()
	signature := createSignature(endpoint, values, nonce, secret)
	headers := map[string]string{
		"Accept":  "application/json",
		"APIKey":  api.key,
		"Nonce":   fmt.Sprint(nonce),
		"Authent": signature,
	}
	resp, err := api.request(method, url, values, headers)
	return resp, err
}

func (api *RestApiClient) request(method string, endpoint string, values url.Values, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	return body, nil
}
