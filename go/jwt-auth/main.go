package main

import (
	"go/jwt-auth/controllers"
	"go/jwt-auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncData()
}

func main() {

	r := gin.Default()

	// this endpoint is for health check
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	// signup user
	r.POST("/signup", controllers.SignUp)
	// signin user
	r.POST("/signin", controllers.SignIn)

	r.Run()
}
