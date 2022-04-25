package serializer

import (
	"car/model"
	"time"
)

//订单序列化器
type Store struct {
	Id           int       `json:"id" `
	MaterialName string    `json:"material_name" `
	Manufacturer string    `json:"manufacturer" `
	UnitPrice    int       `json:"unit_price" `
	Number       int       `json:"number" `
	TotalPrice   int       `json:"total_price" `
	CreatedAt    time.Time `json:"created_at" `
}

// BuildOrder 订单序列化
func BuildStore(g model.Store) Store {
	return Store{
		Id:           g.Id,
		MaterialName: g.MaterialName,
		Manufacturer: g.Manufacturer,
		UnitPrice:    g.UnitPrice,
		Number:       g.Number,
		TotalPrice:   g.TotalPrice,
		CreatedAt:    g.CreatedAt,
	}
}

// BuildOrders 订单集合序列化
func BuildStores(items []model.Store) []Store {
	stores := []Store{}
	for _, item := range items {
		g := BuildStore(item)
		stores = append(stores, g)
	}

	return stores
}
