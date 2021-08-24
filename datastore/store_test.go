package datastore

import (
	"math/rand"
	"testing"
	"vwap-engine/models"

	"github.com/stretchr/testify/assert"
)

func TestStore_UpdateDataAndGetVWAP(t *testing.T) {
	dataLimit := 2
	productID := "testproductID"
	ds := NewDatastore(dataLimit, productID)
	match1 := models.Match{
		ProductID: productID,
		Size:      rand.Float64(),
		Price:     rand.Float64() * 1000,
	}
	vwap := ds.UpdateDataAndGetVWAP(match1)
	assert.Equal(t, (match1.Price*match1.Size)/match1.Size, vwap)
	match2 := models.Match{
		ProductID: productID,
		Size:      rand.Float64(),
		Price:     rand.Float64() * 1000,
	}
	vwap = ds.UpdateDataAndGetVWAP(match2)
	expectedVwap := ((match1.Price * match1.Size) + (match2.Price * match2.Size)) / (match1.Size + match2.Size)
	assert.Equal(t, expectedVwap, vwap)
	match3 := models.Match{
		ProductID: productID,
		Size:      rand.Float64(),
		Price:     rand.Float64() * 1000,
	}
	vwap = ds.UpdateDataAndGetVWAP(match3)
	expectedVwap = ((match3.Price * match3.Size) + (match2.Price * match2.Size)) / (match3.Size + match2.Size)
	assert.Equal(t, expectedVwap, vwap)

}
