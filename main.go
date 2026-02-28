package main

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handlers.GetUsers)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PATCH("/upadteuser/:id", handlers.UpdateUser)
	r.DELETE("/deleteuser/:id", handlers.DeleteUser)

	r.Run("localhost:8888")
}
