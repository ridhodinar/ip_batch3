package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" faker:"-"`
	CreatedAt time.Time      `faker:"-"`
	UpdatedAt time.Time      `faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
