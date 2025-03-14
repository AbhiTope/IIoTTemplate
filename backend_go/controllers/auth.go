package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "iiot_template/go/cache_repo"
)

func HandleLogin(c *gin.Context) {

    var json repo.LoginModel
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if !json.Validate() {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}


func HandleRegister(c *gin.Context) {

    var json repo.RegisterModel
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := json.Add(); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "user created successfuly"})
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, repo.GetUsers())
}

func GetUser(c *gin.Context) {

	var userName string

	if err := c.BindJSON(&userName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	result, err := repo.GetUser(userName)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

