package futures

import "time"

type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

type FutAccountsResult struct {
	Error    []string `json:"error"`
	Result   string
	Accounts struct {
		Fi_xbtusd struct {
			Type      string
			Auxiliary struct {
				USD     float64
				PV      float64
				Pnl     float64
				Af      float64
				Funding float64
			}
			MarginRequirements struct {
				Im float64
				Mm float64
				Lt float64
				Tt float64
			}
			TriggerEstimates struct {
				Im float64
				Mm float64
				Lt float64
				Tt float64
			}
			Balances struct {
				XBT float64
			}
			Currency string
		}
	}
	ServerTime time.Time
}
