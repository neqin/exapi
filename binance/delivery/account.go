package binance_delivery

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (api *RestApiClient) AccountBalance() ([]AccountBalance, error) {
	resp, err := api.queryPrivate("GET", "/balance", url.Values{})
	if err != nil {
		return nil, err
	}
	var jsonData []AccountBalance
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	return jsonData, nil
}

func (api *RestApiClient) GetBalance(asset string) (*AccountBalance, error) {
	res, err := api.AccountBalance()
	if err != nil {
		return nil, err
	}

	for _, i := range res {
		if i.Asset == asset {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("Not found asset %s", asset)
}

func (api *RestApiClient) AccountInformation() (*AccountInformation, error) {
	resp, err := api.queryPrivate("GET", "/account", url.Values{})
	if err != nil {
		return nil, err
	}
	var jsonData AccountInformation
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	return &jsonData, nil
}
