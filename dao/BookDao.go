package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hdas_task/model"
	"log"
)

type Manager2 interface {
	AddBook(book *model.Book)
	UpdateBook(book *model.Book)
    DeleteBook(id int)
	GetAllBooks() []*model.Book
}

type manager2 struct {
	db2 *gorm.DB
}

var Mgr2 Manager2

func init() {
	dsn2:="root:cb123456@tcp(127.0.0.1:3306)/hdas_task?charset=utf8mb4&parseTime=true"
	db2, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err!= nil {
		log.Fatal(err)
	}
	Mgr2 = &manager2{db2: db2}
	db2.AutoMigrate(&model.Book{})
}


func (mgr2 *manager2) AddBook(book *model.Book) {
	mgr2.db2.Create(book)
}

func (mgr2 *manager2) UpdateBook(book *model.Book) {

}

func (mgr2 *manager2) DeleteBook(id int) {

}

func (mgr2 *manager2) GetAllBooks() []*model.Book {
	var books = make([]*model.Book, 0)
    mgr2.db2.Find(&books)
	return books
}