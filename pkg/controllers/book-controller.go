package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alpha2189/book_mgmt/pkg/utils"
	"github.com/alpha2189/go_bookstore/pkg/models"
	"github.com/alpha2189/go_bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks(){
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookid,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _:= models.GetBookById(ID)
	res, _ := json.Marshall(bookDetails)
	w.Header().Set("Content - Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.Response, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
		
	}
	book := models.DeleteBook(ID)
	res, _ : json.Marshal(book)
	w.Header().Set("Content - Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId : vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	bookDetails, db :=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}

	if updateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content - Type", "pkhlication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}