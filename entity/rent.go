package entity

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id          uint       `gorm:"primaryKey;autoIncrement:true"`
	User_id     uint 
	Book_id     uint
	Created_at  time.Time  `gorm:"autoCreateTime"`
}