package controller

import (
	"fmt"
	"log"
	"perpus/entity"

	"gorm.io/gorm"
)

type AksesBook struct {
	DB *gorm.DB
}

func (ab *AksesBook) AddNewBook(userID uint,newBook entity.Book) entity.Book {
	err := ab.DB.Create(&newBook).Where("user_id = ?", userID).Error
	if err != nil {
		log.Println(err)
		return entity.Book{}
	}
	return newBook
}

func (ab *AksesBook) GetAllBook() []entity.Book {
	var daftarBook []entity.Book
	err := ab.DB.Find(&daftarBook)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarBook
}

func (ab *AksesBook) GetBookbyName(judul string) entity.Book {
	var bookname = entity.Book{}
	err := ab.DB.Select([]string{"id", "user_id", "genre_id", "title", "isbn", "author", "penerbit", "jumlah", "deskripsi"}).Where("title = ?", judul).Find(&bookname)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return entity.Book{}
	}
	fmt.Println(" Id buku   :", bookname.Id)
	fmt.Println(" Id user   :", bookname.User_id)
	fmt.Println(" Id genre  :", bookname.Genre_id)
	fmt.Println(" Judul     :", bookname.Title)
	fmt.Println(" Isbn      :", bookname.Isbn)
	fmt.Println(" Pengarang :", bookname.Author)
	fmt.Println(" Penerbit  :", bookname.Penerbit)
	fmt.Println(" Jumlah    :", bookname.Jumlah)
	fmt.Println(" Deskripsi :", bookname.Deskripsi)
	return bookname 
}

func (ab *AksesBook) UpdateDataBook(id uint, input entity.Book) bool {
	err := 	ab.DB.Model(entity.Book{}).Where("id = ?", id).Updates(&input)
	if err.Error != nil {
		log.Fatal(err.Error)
		return false
	}
	if aff := err.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang diupdate")
		return false
	}
	return true
}

func (ab *AksesBook) DeleteBook(idBook uint) bool {
	delete := ab.DB.Where("id = ?", idBook).Delete(&entity.Book{})
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