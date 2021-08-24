package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SubscribeMessage struct {
	MsgType  string    `json:"type"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}

type Match struct {
	TradeID   float64 `json:"trade_id"`
	ProductID string  `json:"trade_id"`
	Time      string  `json:"time"`
	Price     float64 `json:"price"`
	Size      float64 `json:"last_size"`
}

func (m *Match) UnmarshalJSON(b []byte) error {
	var temp map[string]interface{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling into Match object: %s", err.Error())
	}
	priceStr, _ := temp["price"].(string)
	m.Price, err = strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return fmt.Errorf("Error while converting price field to float type: %s", err.Error())
	}
	sizeStr, _ := temp["last_size"].(string)
	m.Size, err = strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return fmt.Errorf("Error while converting size field to float type: %s", err.Error())
	}
	m.Time, _ = temp["time"].(string)
	m.ProductID, _ = temp["product_id"].(string)
	m.TradeID, _ = temp["trade_id"].(float64)
	return nil
}
