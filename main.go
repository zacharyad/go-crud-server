package main

import (
	"user/config"
	"user/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  config.ConnectToDB()

  r.POST("/api/users", controllers.CreateUser)
  r.POST("/api/users/batch", controllers.CreateUsers)
  r.GET("/api/users/:id", controllers.GetUser)
  r.GET("/api/users", controllers.GetUsers)
  r.PUT("/api/users/:id", controllers.UpdateUser)
  r.DELETE("/api/users/:id", controllers.DeleteUser)

  r.Run("localhost:3000")
}
