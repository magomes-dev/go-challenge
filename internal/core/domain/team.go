package domain

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	ID      int
	Name    string
	Country string
	City    string
}
