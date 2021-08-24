package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var AppConfig Config

type Config struct {
	ProductIDs      []string `json:"product_ids"`
	DataPointsLimit int      `json:"data_points_limit"`
	WebsocketURL    string   `json:"websocket_url"`
}

func LoadConfig() {
	raw, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		panic("Unable to read configFile")
	}
	err = json.Unmarshal(raw, &AppConfig)
	if err != nil {
		panic(fmt.Sprintf("Error while unmarshaling conf.json :%s", err.Error()))
	}
}
