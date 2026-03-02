package main

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/users", handlers.GetUsers)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PATCH("/updateuser/:id", handlers.UpdateUser)
	r.DELETE("/deleteuser/:id", handlers.DeleteUser)

	// ================================================================ product order
	r.GET("/products", handlers.GetProducts)
	r.POST("/addcart", handlers.AddChart)
	r.POST("/checkout", handlers.Checkout)
	r.Run("localhost:8888")
}
