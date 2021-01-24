package binance_delivery

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (api *RestApiClient) ListAccountBalance() ([]AccountBalance, error) {
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

func (api *RestApiClient) GetAccountBalance(asset string) (*AccountBalance, error) {
	res, err := api.ListAccountBalance()
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

func (api *RestApiClient) GetAccountInformation() (*AccountInformation, error) {
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

func (api *RestApiClient) GetIncomeHistory(symbol string, incomeType IncomeType, startTime int, endTime int, limit int) ([]IncomeHistory, error) {
	params := url.Values{
		"limit": {fmt.Sprint(limit)},
	}
	if len(incomeType) > 0 {
		params.Add("incomeType", string(incomeType))
	}
	if len(symbol) > 0 {
		params.Add("symbol", symbol)
	}
	if startTime > 0 {
		params.Add("startTime", fmt.Sprint(startTime))
	}
	if endTime > 0 {
		params.Add("endTime", fmt.Sprint(endTime))
	}
	resp, err := api.queryPrivate("GET", "/income", params)

	if err != nil {
		return nil, err
	}
	var jsonData []IncomeHistory
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	return jsonData, nil
}
