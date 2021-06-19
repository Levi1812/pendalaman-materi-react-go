package main

import (
	"backendGolang/auth"
	"backendGolang/book"
	"backendGolang/config"
	"backendGolang/handler"
	"backendGolang/user"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Config()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository, authService)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService)
	middleware              = handler.Middleware(userService, authService)

	bookRepository = book.NewRepository(DB)
	bookService    = book.NewService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/users", middleware, userHandler.GetAllUser)
	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/login", userHandler.LoginUser)
	r.GET("/users/:user_id", middleware, userHandler.GetUserByID)

	r.GET("/books", middleware, bookHandler.GetBookByUserLogin)
	r.POST("/books", middleware, bookHandler.CreateBook)
	r.GET("/books/:book_id", middleware, bookHandler.GetByID)
	r.PUT("/books/:book_id", middleware, bookHandler.UpdateBook)    // jika user id nya tidak sama dengan userid di book, maka tidak bisa update
	r.DELETE("/books/:book_id", middleware, bookHandler.DeleteBook) // jika user id nya tidak sama dengan userid di book, maka tidak bisa delete

	r.Run()
}
