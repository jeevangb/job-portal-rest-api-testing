package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type Jobs struct {
	gorm.Model
	Company Company `json:"-" gorm:"ForeignKey:cid"`
	Cid     uint    `json:"cid"`
	Title   string  `json:"name"`
	Salary  string  `json:"salary"`
}
