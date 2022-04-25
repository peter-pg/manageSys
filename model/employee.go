package model

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Id         int
	Name       string
	Gender     string
	IdCard     string
	Birthday   string
	Phone      string
	Address    string
	IsHead     string
	Department string
	Salary     int
}

func (e *Employee) TableName() string {
	return "t_employee_inf"
}

type Performance struct {
	gorm.Model
	Id          int
	Name        string
	ServiceItem string
	Performance int
	CreatedAt   time.Time
}

func (s *Performance) TableName() string {
	return "t_employee_performance"
}
