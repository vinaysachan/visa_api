package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/base/packages/passport"
	"github.com/vinaysachan/visa_api/data/tasks"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Auth Token provided"})
			return
		}

		// Validate token and get token data
		authToken, err := passport.GetTokenData(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		//Get User Data from User Task:
		user, err := tasks.FindUserByID(authToken.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		}

		// Add user info to context for use in handlers
		c.Set("authUser", user)
		c.Set("userID", authToken.UserID)
		c.Set("accessToken", authToken.ID)

		c.Next()
	}
}
