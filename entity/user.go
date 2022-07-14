package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey;autoIncrement:true"`
	Nama     string `gorm:"type:varchar(100)"`
	Address  string `gorm:"type:varchar(50)"`
	Hp       string `gorm:"type:varchar(13)"`
	Password string `gorm:"type:varchar(7)"`
	Books []Book `gorm:"foreignKey:User_id;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rents []Rent `gorm:"foreignKey:User_id;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}