package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	bookService BookService
}

func NewBookController(bookService BookService) *bookController {
	return &bookController{bookService}
}

func (handler *bookController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (handler *bookController) FindAll(c *gin.Context) {
	book, err := handler.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
