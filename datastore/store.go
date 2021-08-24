package datastore

import (
	"vwap-engine/models"
)

type Store struct {
	ProductID   string
	Matches     []models.Match
	TotalValue  float64
	TotalVolume float64
	DataLimit   int
}

var _ models.DataStore = &Store{}

func NewDatastore(dataLimit int, productID string) *Store {
	return &Store{
		ProductID: productID,
		DataLimit: dataLimit,
		Matches:   []models.Match{},
	}
}

func (store *Store) UpdateDataAndGetVWAP(match models.Match) float64 {
	store.TotalValue += match.Price * match.Size
	store.TotalVolume += match.Size
	store.Matches = append(store.Matches, match)
	if len(store.Matches) > store.DataLimit {
		store.TotalValue -= (store.Matches[0].Price * store.Matches[0].Size)
		store.TotalVolume -= store.Matches[0].Size
		store.Matches = store.Matches[1:]
	}
	return store.TotalValue / store.TotalVolume
}

func (store *Store) GetProductID() string {
	return store.ProductID
}
