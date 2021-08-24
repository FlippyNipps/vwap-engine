package match_processor

import (
	"encoding/json"
	"fmt"
	"log"
	"vwap-engine/models"
)

type MatchProcessor struct {
	wsFeed            models.Websocket
	ProductIDStoreMap map[string]models.DataStore
}

var _ models.MatchProcessor = &MatchProcessor{}

func NewMatchProcessor(ws models.Websocket, datastores []models.DataStore) *MatchProcessor {
	dsMap := map[string]models.DataStore{}
	for _, store := range datastores {
		dsMap[store.GetProductID()] = store
	}
	return &MatchProcessor{
		ws,
		dsMap,
	}
}

func (mp *MatchProcessor) Subscribe(productIDs []string) error {
	if len(productIDs) == 0 {
		return fmt.Errorf("Empty product IDs")
	}
	subMsg := &models.SubscribeMessage{
		"subscribe",
		[]models.Channel{models.Channel{
			"ticker",
			productIDs,
		}},
	}
	err := mp.wsFeed.WriteMessage(subMsg)
	//TODO error logging
	if err != nil {
		return fmt.Errorf("Error while writing ws message: %w", err)
	}

	msg, err := mp.wsFeed.ReadMessage()
	if err != nil {
		return fmt.Errorf("Error while reading ws response: %w", err)
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Error while reading ws response: %w", err)
	}
	var respMessage models.SubscribeMessage
	err = json.Unmarshal(msgBytes, &respMessage)
	if err != nil {
		return fmt.Errorf("Error while reading ws response: %w", err)
	}
	if len(respMessage.Channels) == 0 || len(respMessage.Channels[0].ProductIDs) != len(productIDs) {
		return fmt.Errorf("Failed to subscribe to %s", productIDs)
	}
	return nil
}

func (mp *MatchProcessor) Run() {
	for {
		msg, err := mp.wsFeed.ReadMessage()
		if err != nil {
			log.Printf("Error while reading ws response: %s", err.Error())
			break
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error while reading ws response: %s", err.Error())
			break
		}
		var respMessage models.Match
		err = json.Unmarshal(msgBytes, &respMessage)
		if err != nil {
			fmt.Printf("Error while reading ws response: %s", err.Error())
			break
		}
		vwapValue := mp.ProductIDStoreMap[respMessage.ProductID].UpdateDataAndGetVWAP(respMessage)
		fmt.Println(fmt.Sprintf("%s VWAP: %f", respMessage.ProductID, vwapValue))
	}

	err := mp.wsFeed.Close()
	if err != nil {
		log.Printf("Error while closing websocket connection :%s", err.Error())
	}
}
