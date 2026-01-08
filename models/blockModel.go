package models

import "time"

type Block struct {
	Timestamp time.Time `gorm:"type:datetime;not null"`
	Data      []byte    `gorm:"type:varchar(255);not null"`
	PrevHash  []byte    `gorm:"type:varchar(255);not null"`
	Hash      []byte    `gorm:"type:varchar(255);not null"`
}
