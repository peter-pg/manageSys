package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id          int
	UserName    string
	UserPhone   string
	CarLicense  string
	CarBrand    string
	ServiceItem string
	Cost        int
	Material    string
	Employee    string
	Operator    string
	CreatedAt   time.Time
}

type ServiceCost struct {
	gorm.Model
	Id           int
	CreatedAt    time.Time
	CarLicense   string
	ServiceItem  string
	Material     string
	MaterialCost int
}

func (s *ServiceCost) TableName() string {
	return "t_service_cost"
}

func (o *Order) TableName() string {
	return "t_service"
}
