package handler

import (
	"backendGolang/book"
	"backendGolang/user"
	"strconv"

	// "os/user"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) GetBookByUserLogin(c *gin.Context) {
	userID := c.MustGet("currentUser").(user.UserFormatter)

	userIDstr := strconv.Itoa(userID.ID)

	books, err := h.service.GetAllBookByUser(userIDstr)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if userID.ID == 0 {
		c.JSON(401, gin.H{"error": "unauthorize user"})
		return
	}

	c.JSON(200, books)
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var input book.CreateBook

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	userID := c.MustGet("currentUser").(user.UserFormatter)

	// userIDstr := strconv.Itoa(userID.ID)

	book, err := h.service.Createbook(userID.ID, input)

	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(201, book)
}

func (h *bookHandler) GetByID(c *gin.Context) {
	bookID := c.Params.ByName("book_id")

	book, err := h.service.GetBookByID(bookID)

	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, book)
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	bookID := c.Params.ByName("book_id")

	var input book.UpdateBook

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	book, err := h.service.UpdateBook(bookID, input)

	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, book)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	bookID := c.Params.ByName("book_id")

	message, err := h.service.DeleteBook(bookID)

	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": message})
}
