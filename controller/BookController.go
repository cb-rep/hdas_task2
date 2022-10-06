package controller

import (
	"github.com/gin-gonic/gin"
	"hdas_task/dao"
	"hdas_task/model"
)

func AddBook(c *gin.Context) {
	kind := c.PostForm("kind")
	name := c.PostForm("name")
	description := c.PostForm("description")

	book :=model.Book{
		Kind:  kind,
        Name:      name,
        Description: description,
	}

	dao.Mgr2.AddBook(&book)
	c.Redirect(301,"/books")
}

func GoAddBook(c *gin.Context) {
	c.HTML(200,"addBook.html",nil)
}

func GetAllBooks(c *gin.Context) {
	books:=dao.Mgr2.GetAllBooks()
	c.HTML(200,"books.html",books)
}

