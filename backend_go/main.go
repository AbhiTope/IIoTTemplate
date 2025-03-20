package main

import (
	con "iiot_template/go/controllers"
	mw "iiot_template/go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.POST("/login", con.HandleLogin)
    r.POST("/validate", con.Validate)
    r.POST("/register", con.HandleRegister)
    r.GET("/getuser", mw.TokenAuthMiddleware([]string{"admin", "user"}), con.GetUser)
    r.GET("/getusers", mw.TokenAuthMiddleware([]string{"admin"}), con.GetUsers)

    r.Run() 
}

