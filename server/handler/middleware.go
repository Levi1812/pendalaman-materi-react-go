package handler

import (
	"backendGolang/auth"
	"backendGolang/user"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		token, err := authService.ValidateToken(header)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		userID := int(claim["user_id"].(float64))

		userIDstr := strconv.Itoa(userID)

		// handle error nya gimana

		user, err := userService.GetUserByID(userIDstr)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		if user.ID == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize user"})
			return
		}

		c.Set("currentUser", user) // set id
	}
}

func InvestorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
