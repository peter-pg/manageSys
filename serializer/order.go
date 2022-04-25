package serializer

import (
	"car/model"
	"time"
)

//订单序列化器
type Order struct {
	Id          int       `json:"id" `
	UserName    string    `json:"user_name" `
	UserPhone   string    `json:"user_phone" `
	CarLicense  string    `json:"car_license" `
	CarBrand    string    `json:"car_brand" `
	ServiceItem string    `json:"service_item" `
	Cost        int       `json:"cost" `
	Material    string    `json:"material" `
	Employee    string    `json:"employee" `
	Operator    string    `json:"operator" `
	CreatedAt   time.Time `json:"created_at" `
}

type ServiceCost struct {
	Id           int       `json:"id" `
	CreatedAt    time.Time `json:"created_at" `
	CarLicense   string    `json:"car_license" `
	ServiceItem  string    `json:"service_item" `
	Material     string    `json:"material" `
	MaterialCost int       `json:"material_cost" `
}

// BuildOrder 订单序列化
func BuildOrder(g model.Order) Order {
	return Order{
		Id:          g.Id,
		UserName:    g.UserName,
		UserPhone:   g.UserPhone,
		CarLicense:  g.CarLicense,
		CarBrand:    g.CarBrand,
		ServiceItem: g.ServiceItem,
		Cost:        g.Cost,
		Material:    g.Material,
		Employee:    g.Employee,
		Operator:    g.Operator,
		CreatedAt:   time.Now(),
	}
}

// BuildOrders 订单集合序列化
func BuildOrders(items []model.Order) []Order {
	orders := []Order{}
	for _, item := range items {
		g := BuildOrder(item)
		orders = append(orders, g)
	}

	return orders
}

// BuildServiceCost 订单序列化
func BuildServiceCost(g model.ServiceCost) ServiceCost {
	return ServiceCost{
		Id:           g.Id,
		CarLicense:   g.CarLicense,
		ServiceItem:  g.ServiceItem,
		MaterialCost: g.MaterialCost,
		Material:     g.Material,
		CreatedAt:    time.Now(),
	}
}

// BuildServiceCosts 订单集合序列化
func BuildServiceCosts(items []model.ServiceCost) []ServiceCost {
	costs := []ServiceCost{}
	for _, item := range items {
		g := BuildServiceCost(item)
		costs = append(costs, g)
	}

	return costs
}
