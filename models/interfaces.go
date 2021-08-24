package models

type Websocket interface {
	Connect() error
	Close() error
	ReadMessage() (interface{}, error)
	WriteMessage(interface{}) error
}

type MatchProcessor interface {
	Subscribe([]string) error
	Run()
}

type DataStore interface {
	UpdateDataAndGetVWAP(Match) float64
	GetProductID() string
}
