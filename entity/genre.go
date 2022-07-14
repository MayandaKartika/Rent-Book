package entity

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Id   uint   `gorm:"primaryKey;autoIncrement:true"`
	Nama string `gorm:"type:varchar(20)"`
	Books Book `gorm:"foreignKey:Genre_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
}

