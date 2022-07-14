package controller

import (
	"log"
	"perpus/entity"

	"gorm.io/gorm"
)

type AksesRent struct {
	DB *gorm.DB
}

func (ar *AksesRent) AddNewRent(idUser uint,newRent entity.Rent) entity.Rent {
	err := ar.DB.Create(&newRent).Where("user_id = ?", idUser).Error
	if err != nil {
		log.Println(err)
		return entity.Rent{}
	}
	return newRent
}

func (ar *AksesRent) GetBookbyUserID(id uint) []entity.Rent{
	var listBook = []entity.Rent{}
	err := ar.DB.Select([]string{"user_id", "book_id"}).Where("user_id = ?", id).Find(&listBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return []entity.Rent{}
	}
	return listBook
}

func (ar *AksesRent) ReturnBook(IdBook uint) bool {
	delete := ar.DB.Where("book_id = ?", IdBook).Delete(&entity.Rent{})
	if err := delete.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := delete.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	return true

}