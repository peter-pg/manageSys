package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Id           int
	MaterialName string
	Manufacturer string
	Number       int
	UnitPrice    int
	TotalPrice   int
	CreatedAt    time.Time
}

func (s *Store) TableName() string {
	return "t_store"
}
