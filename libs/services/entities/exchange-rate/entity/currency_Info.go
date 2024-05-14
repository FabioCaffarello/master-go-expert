package exchangerateentity

import (
	"errors"
	"fmt"
	gouuid "libs/shared/go-uuid"
	"time"
)

var (
	errBidRequired       = errors.New("bid is required")
	errCodeRequired      = errors.New("code is required")
	errCodeInRequired    = errors.New("codein is required")
	errTimestampRequired = errors.New("timestamp is required")
)

type CurrencyInfo struct {
	ID         gouuid.ID `json:"exchange_rate_id"`
	Code       string    `json:"code"`
	CodeIn     string    `json:"codein"`
	Name       string    `json:"name"`
	High       float64   `json:"high"`
	Low        float64   `json:"low"`
	VarBid     float64   `json:"varBid"`
	PctChange  float64   `json:"pctChange"`
	Bid        float64   `json:"bid"`
	Ask        float64   `json:"ask"`
	Timestamp  int64     `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func NewExchangeRate(
	code string,
	codeIn string,
	name string,
	high float64,
	low float64,
	varBid float64,
	pctChange float64,
	bid float64,
	ask float64,
	timestamp int64,
	createDate string,
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
	}

	if err := currencyInfo.isValid(); err != nil {
		return nil, err
	}

	err := currencyInfo.setEntityID(code, codeIn, timestamp)
	if err != nil {
		return nil, err
	}

	err = currencyInfo.setCreateDate(createDate)
	if err != nil {
		return nil, err
	}

	return currencyInfo, nil
}

func (e *CurrencyInfo) setCreateDate(createDate string) error {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, createDate)
	if err != nil {
		return fmt.Errorf("error parsing create date: %w", err)
	}
	e.CreateDate = parsedTime
	return nil
}

func (e *CurrencyInfo) setEntityID(code string, codeIn string, timestamp int64) error {
	propertiesID := map[string]interface{}{
		"code":      code,
		"codeIn":    codeIn,
		"timestamp": timestamp,
	}
	entityID, err := gouuid.GetID(propertiesID)
	if err != nil {
		return err
	}
	e.ID = entityID
	return nil
}

func (e *CurrencyInfo) isValid() error {
	if e.Code == "" {
		return errCodeInRequired
	}
	if e.CodeIn == "" {
		return errCodeInRequired
	}
	if e.Timestamp == 0 {
		return errTimestampRequired
	}
	if e.Bid == 0 {
		return errBidRequired
	}
	return nil
}
