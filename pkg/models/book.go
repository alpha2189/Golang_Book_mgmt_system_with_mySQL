package models

import (
	"github.com/alpha2189/go_bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)


var db *gorm.DB
type Book struct{
	gorm.Model
	Name string 'gorm:""json:"name"'
	Author string 'json:"author"'
	Publication string 'json:"publication"'
}


func init(){
	config.Connect()
	db = config.GetDB
	db. AutoMigrate(&Book{})
}

func (b * Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b 
}

func GetAllBooks() []Book{
	var Books []Bookdb.Find(&Books)
	return Books 
}


func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book 
	Bookdb.Where("ID=?", ID).Delete(book)
	return book
}
