package binance_delivery

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	APIURL = "https://dapi.binance.com/dapi/v1"
)

type RestApiClient struct {
	key       string
	secret    string
	client    *http.Client
	window    int64
	userAgent string
}

func NewRestApi(key, secret string) *RestApiClient {
	return &RestApiClient{
		key:       key,
		secret:    secret,
		client:    http.DefaultClient,
		window:    155000,
		userAgent: "Binance/golang",
	}
}

func NewRestApiWithClient(key, secret string, httpClient *http.Client) *RestApiClient {
	kraken := NewRestApi(key, secret)
	kraken.client = httpClient
	return kraken
}

func (api *RestApiClient) Time() error {
	_, err := api.queryPublic("GET", "/time", nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *RestApiClient) queryPublic(method string, endpoint string, values url.Values) ([]byte, error) {
	url := fmt.Sprintf("%s%s", APIURL, endpoint)
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   api.userAgent,
	}
	resp, err := api.request(method, url, values, headers)

	return resp, err
}

func (api *RestApiClient) queryPrivate(method string, endpoint string, values url.Values) ([]byte, error) {
	values.Add("timestamp", fmt.Sprint(time.Now().UnixNano()/(1000*1000)))
	payload := fmt.Sprintf("%s&recvWindow=%d", values.Encode(), api.window)
	mac := hmac.New(sha256.New, []byte(api.secret))
	_, err := mac.Write([]byte(payload))
	if err != nil {
		return nil, err
	}
	payload = fmt.Sprintf("%s&signature=%s", payload, hex.EncodeToString(mac.Sum(nil)))
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   api.userAgent,
		"X-MBX-APIKEY": api.key,
	}
	url := fmt.Sprintf("%s%s?%s", APIURL, endpoint, payload)
	resp, err := api.request(method, url, nil, headers)
	return resp, err
}

func (api *RestApiClient) request(method string, url string, values url.Values, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! NewRequest (%s)", err.Error())
	}
	if values == nil {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, fmt.Errorf("Could not execute request! NewRequest (%s)", err.Error())
		}
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! Do (%s)", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! ReadAll(%s)", err.Error())
	}
	return body, nil
}

func Marshal(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return "{marshal error}"
	}
	return string(res)
}
