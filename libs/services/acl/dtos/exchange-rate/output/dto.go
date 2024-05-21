package outputdto

// ExchangeRateDTO is a data transfer object that represents the exchange rate api output.
type ExchangeRateDTO struct {
	Code       string  `json:"code"`
	CodeIn     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	VarBid     float64 `json:"varBid"`
	PctChange  float64 `json:"pctChange"`
	Bid        float64 `json:"bid"`
	Ask        float64 `json:"ask"`
	Timestamp  int64   `json:"timestamp"`
	CreateDate string  `json:"create_date"`
}

// ExchangeRatesDTO is a data transfer object that represents a map of exchange rates.
type ExchangeRatesDTO map[string]ExchangeRateDTO
