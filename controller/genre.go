package controller

import (
	"fmt"
	"log"
	"perpus/entity"

	"gorm.io/gorm"
)

type AksesGenre struct {
	DB *gorm.DB
}

func (ag *AksesGenre) AddGenre(newGenre entity.Genre) entity.Genre {
	err := ag.DB.Create(&newGenre).Error
	if err != nil {
		log.Println(err)
		return entity.Genre{}
	}
	return newGenre
}

func (ag *AksesGenre) GetAllGenre() []entity.Genre {
	var ListGenre []entity.Genre
	err := ag.DB.Find(&ListGenre)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return ListGenre
}

func (ag *AksesGenre) GetGenrebyName(nama string) entity.Genre {
	var genre = entity.Genre{}
	err := ag.DB.Where("nama = ?", nama).First(&genre)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return entity.Genre{}
	}
	fmt.Println(" Id : ", genre.Id)
	fmt.Println(" Alamat : ", genre.Nama)
	return genre
}

func (ag *AksesGenre) DeleteGenre(IdGenre int) bool {
	delete := ag.DB.Where("id = ?", IdGenre ).Delete(&entity.Genre{})
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