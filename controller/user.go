package controller

import (
	"fmt"
	"log"
	"perpus/entity"

	"gorm.io/gorm"
)

type AksesUser struct {
	DB *gorm.DB
}

func (au *AksesUser) Register(newUser entity.User) entity.User {
	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return entity.User{}
	}
	return newUser
}

func (au *AksesUser) Login(hp, password string) entity.User {
	var user = entity.User{}
	err := au.DB.Where("hp = ? AND password = ?", hp, password).First(&user)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return entity.User{}
	}
	return user
}

func (au *AksesUser) GetAllUser() []entity.User {
	var daftarUser = []entity.User{}
	err := au.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	return daftarUser
}

func (au *AksesUser) GetUserbyName(nama string) entity.User{
	var username = entity.User{}
	err := au.DB.Select([]string{"id", "nama", "address", "hp"}).Where("nama = ?", nama).Find(&username)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return entity.User{}
	}
	fmt.Println(" ID User : ", username.Id)
	fmt.Println(" Nama : ", username.Nama)
	fmt.Println(" Alamat : ", username.Address)
	fmt.Println(" Nomer HP : ", username.Hp)
	return username
}

func (au *AksesUser) UpdateDataUser(hp string, input entity.User) bool {
	err := 	au.DB.Model(entity.User{}).Where("hp = ?", hp).Updates(&input)
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

func (as *AksesUser) DeleteUser(IdUser uint) bool {
	delete := as.DB.Where("id = ?", IdUser).Delete(&entity.User{})
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