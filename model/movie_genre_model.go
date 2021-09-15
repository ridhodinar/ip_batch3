package model

import (
	"gorm.io/gorm"
)

type MovieGenre struct {
	gorm.Model
	Movie   string
	Genre   string
	MovieID int
	GenreID int
}
