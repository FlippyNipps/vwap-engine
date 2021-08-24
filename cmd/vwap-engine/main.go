package main

import (
	"fmt"
	"vwap-engine/config"
	"vwap-engine/datastore"
	match_processor "vwap-engine/match-processor"
	"vwap-engine/models"
	"vwap-engine/websocket"
)

func main() {
	config.LoadConfig()
	webSocket := websocket.NewWebsocket(config.AppConfig.WebsocketURL)
	err := webSocket.Connect()
	//TODO add logging package
	if err != nil {
		panic("Unable to setup websocket connection")
	}
	productIDs := config.AppConfig.ProductIDs
	stores := []models.DataStore{}
	for _, prodID := range productIDs {
		stores = append(stores, datastore.NewDatastore(config.AppConfig.DataPointsLimit, prodID))
	}
	matchProcessor := match_processor.NewMatchProcessor(webSocket, stores)
	err = matchProcessor.Subscribe(productIDs)
	if err != nil {
		panic(fmt.Errorf("Error: Unable to subscribe to the specified Product IDs: %w", err))
	}

	matchProcessor.Run()

}
