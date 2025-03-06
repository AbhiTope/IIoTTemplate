package main

import (
    "github.com/gin-gonic/gin"
    "iiot_template/go/controllers"
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
    r.GET("/get", controller.GetUsers)

    r.Run() 
}

