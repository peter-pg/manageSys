package model

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Id          int
	UserName    string
	UserPhone   string
	CarLicense  string
	CarBrand    string
	ServiceNums int
	Consumption int
	IsVip       string
}

func (e *Guest) TableName() string {
	return "t_guest"
}
