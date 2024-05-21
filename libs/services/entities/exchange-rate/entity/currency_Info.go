package exchangerateentity

import (
	"encoding/json"
	"errors"
	"fmt"
	gouuid "libs/shared/go-uuid"
	"log"
	"time"
)

var (
	errIDRequired        = errors.New("id is required")
	errBidRequired       = errors.New("bid is required")
	errCodeRequired      = errors.New("code is required")
	errCodeInRequired    = errors.New("codein is required")
	errTimestampRequired = errors.New("timestamp is required")
)

// CurrencyInfo represents the exchange rate information for a currency pair.
type CurrencyInfo struct {
	ID         gouuid.ID `json:"_id"`
	Code       string    `json:"code"`
	CodeIn     string    `json:"codeIn"`
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

// NewExchangeRate creates a new CurrencyInfo entity from string inputs.
// It converts string inputs to appropriate data types, validates them,
// and initializes a new CurrencyInfo object.
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
	timestamp string,
	createDate string,
) (*CurrencyInfo, error) {
	highFloat, err := StringToFloat64(high)
	if err != nil {
		return nil, err
	}
	lowFloat, err := StringToFloat64(low)
	if err != nil {
		return nil, err
	}
	varBidFloat, err := StringToFloat64(varBid)
	if err != nil {
		return nil, err
	}
	pctChangeFloat, err := StringToFloat64(pctChange)
	if err != nil {
		return nil, err
	}
	bidFloat, err := StringToFloat64(bid)
	if err != nil {
		return nil, err
	}
	askFloat, err := StringToFloat64(ask)
	if err != nil {
		return nil, err
	}
	timestampInt, err := StringToInt64(timestamp)
	if err != nil {
		return nil, err
	}

	currencyInfo := &CurrencyInfo{
		Code:      code,
		CodeIn:    codeIn,
		Name:      name,
		High:      highFloat,
		Low:       lowFloat,
		VarBid:    varBidFloat,
		PctChange: pctChangeFloat,
		Bid:       bidFloat,
		Ask:       askFloat,
		Timestamp: timestampInt,
	}

	err = currencyInfo.setEntityID(code, codeIn, timestamp)
	if err != nil {
		return nil, err
	}

	if err := currencyInfo.isValid(); err != nil {
		return nil, err
	}

	err = currencyInfo.setCreateDate(createDate)
	if err != nil {
		return nil, err
	}

	return currencyInfo, nil
}

// setCreateDate sets the create date of the CurrencyInfo object
// by parsing the input string to a time.Time object.
func (e *CurrencyInfo) setCreateDate(createDate string) error {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, createDate)
	if err != nil {
		return fmt.Errorf("error parsing create date: %w", err)
	}
	e.CreateDate = parsedTime
	return nil
}

// setEntityID sets the ID of the CurrencyInfo object
// using the code, codeIn, and timestamp properties.
func (e *CurrencyInfo) setEntityID(code string, codeIn string, timestamp string) error {
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

// isValid validates the essential fields of the CurrencyInfo object.
func (e *CurrencyInfo) isValid() error {
	if e.ID == "" {
		return errIDRequired
	}
	if e.Code == "" {
		return errCodeRequired
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

// ToMap converts the CurrencyInfo object to a map representation.
func (e *CurrencyInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":         string(e.ID),
		"code":        e.Code,
		"codeIn":      e.CodeIn,
		"name":        e.Name,
		"high":        e.High,
		"low":         e.Low,
		"varBid":      e.VarBid,
		"pctChange":   e.PctChange,
		"bid":         e.Bid,
		"ask":         e.Ask,
		"timestamp":   e.Timestamp,
		"create_date": e.CreateDate,
	}
}

// MapToCurrencyInfoEntity converts a map representation of a CurrencyInfo object
// back to a CurrencyInfo entity.
func MapToCurrencyInfoEntity(document map[string]interface{}) (*CurrencyInfo, error) {
	var documentEntity CurrencyInfo
	documentBytes, err := json.Marshal(document)
	if err != nil {
		log.Printf("Error marshalling document: %v", err)
		return &CurrencyInfo{}, err
	}
	err = json.Unmarshal(documentBytes, &documentEntity)
	if err != nil {
		log.Printf("Error unmarshalling document: %v", err)
		return &CurrencyInfo{}, err
	}

	if err = documentEntity.isValid(); err != nil {
		return &CurrencyInfo{}, err
	}
	return &documentEntity, nil
}
