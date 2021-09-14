package review

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Review  string
	Rate    int
	UserID  int
	MovieID int
}
