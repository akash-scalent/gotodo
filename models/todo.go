package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name string `gorm:"not null"`
	Description string 
}