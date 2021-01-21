package binance_delivery

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"
	NewOrderRespTypeFULL   NewOrderRespType = "FULL"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"

	SymbolTypeFuture SymbolType = "FUTURE"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize       SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePrice         SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice  SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMarketLotSize SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders  SymbolFilterType = "MAX_NUM_ORDERS"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCrossed  MarginType = "CROSSED"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// SideEffectType define side effect type for orders
type SideEffectType string

// WorkingType define working type
type WorkingType string

// MarginType define margin type
type MarginType string

type AccountBalance struct {
	AccountAlias         string
	Asset                string
	Balance              float64 `json:",string"`
	WithdrawAvailableStr float64 `json:",string"`
	CrossWalletBalance   float64 `json:",string"`
	CrossUnPnl           float64 `json:",string"`
	AvailableBalance     float64 `json:",string"`
	UpdateTime           int64
}

type Asset struct {
	Asset                  string
	WalletBalance          float64 `json:",string"`
	UnrealizedProfit       float64 `json:",string"`
	MarginBalance          float64 `json:",string"`
	MaintMargin            float64 `json:",string"`
	InitialMargin          float64 `json:",string"`
	PositionInitialMargin  float64 `json:",string"`
	OpenOrderInitialMargin float64 `json:",string"`
	MaxWithdrawAmount      float64 `json:",string"`
	CrossWalletBalance     float64 `json:",string"`
	CrossUnPnl             float64 `json:",string"`
	AvailableBalance       float64 `json:",string"`
}

type Position struct {
	Symbol                 string
	InitialMargin          float64 `json:",string"`
	MaintMargin            float64 `json:",string"`
	UnrealizedProfit       float64 `json:",string"`
	PositionInitialMargin  float64 `json:",string"`
	OpenOrderInitialMargin float64 `json:",string"`
	Leverage               float64 `json:",string"`
	Isolated               bool
	PositionSide           PositionSideType
	EntryPrice             float64 `json:",string"`
	MaxQty                 float64 `json:",string"`
}

type AccountInformation struct {
	Assets      []Asset
	Positions   []Position
	CanDeposit  bool
	CanTrade    bool
	CanWithdraw bool
	FeeTier     int
	UpdateTime  int
}
