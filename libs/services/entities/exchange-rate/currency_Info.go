package exchangerateentity

import (
	"errors"
	"time"
)

type CurrencyInfo struct {
	Code       string    `json:"code"`
	CodeIn     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varBid"`
	PctChange  string    `json:"pctChange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  int64     `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func NewExchangeRate(
	code string,
	codeIn string,
	name string,
	high string,
	low string,
	varBid string,
	pctChange string,
	bid string,
	ask string,
	timestamp int64,
	createDate time.Time,
) (*CurrencyInfo, error) {
	currencyInfo := &CurrencyInfo{
		Code:       code,
		CodeIn:     codeIn,
		Name:       name,
		High:       high,
		Low:        low,
		VarBid:     varBid,
		PctChange:  pctChange,
		Bid:        bid,
		Ask:        ask,
		Timestamp:  timestamp,
		CreateDate: createDate,
	}

	if err := currencyInfo.isValid(); err != nil {
		return nil, err
	}

	return currencyInfo, nil
}

func (e *CurrencyInfo) isValid() error {
	if e.Bid == "" {
		return errors.New("bid is required")
	}
	return nil
}
