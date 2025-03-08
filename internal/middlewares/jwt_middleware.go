package middlewares

import (
	"mentalartsapi/internal/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is used for user authentication
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Validate the token format (Bearer token)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			c.Abort()
			return
		}

		// Parse and validate the JWT token
		claims := &utils.JWTClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Attach the claims to the context for later use
		c.Set("user", claims)

		c.Next()
	}
}

// AdminOnly middleware ensures that only admins can access certain routes
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user claims from the context
		userClaims, _ := c.Get("user")
		claims, ok := userClaims.(*utils.JWTClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user claims"})
			c.Abort()
			return
		}

		// Check if the user is an admin
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
