package outputdto

// CurrencyInfoDTO contains detailed information about the exchange rate.
type CurrencyInfoDTO struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// CurrencyInfoMapDTO is a map of currency codes to CurrencyInfo.
type CurrencyInfoMapDTO map[string]CurrencyInfoDTO
