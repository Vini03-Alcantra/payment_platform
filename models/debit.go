package models

import (
	"time"

	"gorm.io/gorm"
)

type DebitValue struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	NameCard      string `json:"name_Card"`
	LastFourDigit string `json:"last_Four_digit"`
	ValueDebit    string `json:"value_debit"`

	CreatedAt time.Time      `json:created`
	UpdatedAt time.Time      `json:updated`
	DeletedAt gorm.DeletedAt `json:"index" json:"deleted"`
}
