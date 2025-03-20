package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "iiot_template/go/cache_repo"
    "iiot_template/go/utils"
)

func HandleLogin(c *gin.Context) {

    var json repo.LoginModel
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := repo.GetPassword(json.UserName)
    if err != nil{
	    c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
	    return
    }

    if !utils.CheckPasswordHash(json.Password, result) {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    }

    token, err := utils.GenerateToken(json.UserName, "admin")

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error Login Please try again."})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Validate(c *gin.Context){

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // if no "Bearer" prefix
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			return
		}

		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{ "message": err.Error()})
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.JSON(http.StatusOK, gin.H{
			"message": "valid token",
			"username": claims.Username,
			"role": claims.Role,
		})
}


func HandleRegister(c *gin.Context) {

    var json repo.RegisterModel
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    passwordHash, err := utils.HashPassword(json.Password);

    if err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
    }

    json.Password = passwordHash
    json.Role = "user"
    json.IsLocked = false

    if err := json.Add(); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "user created successfuly"})
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, repo.GetActiveUsers())
}

func GetUser(c *gin.Context) {

	userName := c.DefaultQuery("userName", "") // Default value is empty string

	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userName query parameter is required"})
		return
	}

	result, err := repo.GetUser(userName)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

