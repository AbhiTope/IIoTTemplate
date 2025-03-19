package main

import (
	"iiot_template/go/controllers"
	"iiot_template/go/utils"
	"net/http"
	"strings"
	"slices"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.POST("/login", controller.HandleLogin)
    r.POST("/register", controller.HandleRegister)
    r.GET("/getuser", TokenAuthMiddleware([]string{"admin", "user"}), controller.GetUser)
    r.GET("/getusers", TokenAuthMiddleware([]string{"admin"}), controller.GetUsers)

    r.Run() 
}

func TokenAuthMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // if no "Bearer" prefix
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{ "message": err.Error()})
			c.Abort()
			return
		}

		// Check if the role in the token matches the required role
		//if claims.Role != requiredRole {
		if !slices.Contains(allowedRoles, claims.Role)  {
			c.JSON(http.StatusForbidden, gin.H{"message": "You do not have the required role"})
			c.Abort()
			return
		}

		// Store the claims in the context for later use
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
