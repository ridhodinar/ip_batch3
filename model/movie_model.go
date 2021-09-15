package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Ratings int
	Title   string
	Year    int
}
