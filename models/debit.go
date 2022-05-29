package models

import (
	"time"

	"gorm.io/gorm"
)

type DebitCard struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	NameCard      string  `json:"name_Card"`
	LastFourDigit string  `json:"last_Four_digit"`
	value         float64 `json:"value"`

	CreatedAt time.Time      `json:created`
	UpdatedAt time.Time      `json:updated`
	DeletedAt gorm.DeletedAt `json:"index" json:"deleted"`
}
