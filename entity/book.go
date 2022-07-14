package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey;autoIncrement:true"`
	User_id   uint 
	Genre_id  uint
	Title     string `gorm:"type:varchar(50)"`
	Isbn      string `gorm:"type:varchar(13)"`
	Author    string `gorm:"type:varchar(10)"`
	Penerbit  string `gorm:"type:varchar(10)"`
	Jumlah    uint   
	Deskripsi string `gorm:"type:varchar(10)"`
	Rents Rent `gorm:"foreignKey:Book_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}