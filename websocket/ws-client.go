package websocket

import (
	"fmt"
	"log"
	"vwap-engine/models"

	ws "golang.org/x/net/websocket"
)

type Websocket struct {
	conn *ws.Conn
	URL  string
}

var _ models.Websocket = &Websocket{}

func NewWebsocket(url string) *Websocket {
	return &Websocket{
		URL: url,
	}
}

func (websocket *Websocket) Connect() error {
	conn, err := ws.Dial(websocket.URL, "", "http://localhost")
	if err != nil {
		return fmt.Errorf("Error while setting up websocket: %w", err)
	}
	log.Println("Websocket connection setup successfully")
	websocket.conn = conn
	return nil
}

func (websocket *Websocket) Close() error {
	if err := websocket.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (websocket *Websocket) WriteMessage(msg interface{}) error {
	err := ws.JSON.Send(websocket.conn, msg)
	if err != nil {
		return fmt.Errorf("Error while writing websocket message: %w", err)
	}
	return nil
}

func (websocket *Websocket) ReadMessage() (interface{}, error) {
	var message interface{}
	err := ws.JSON.Receive(websocket.conn, &message)
	if err != nil {
		return nil, fmt.Errorf("Error while reading websocket message: %w", err)
	}
	return message, nil
}
