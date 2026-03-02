package main

import (
	"backend/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Header("Access-Control-Allow-Headers", "content-type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}

func main() {
	r := gin.Default()

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		r.GET("/users", handlers.GetUsers)
		r.POST("/register", handlers.Register)
		r.POST("/login", handlers.Login)
		r.GET("/users/:id", handlers.GetUserByID)
		r.PATCH("/updateuser/:id", handlers.UpdateUser)
		r.DELETE("/deleteuser/:id", handlers.DeleteUser)
	}

	// ================================================================ product order
	r.GET("/products", handlers.GetProducts)
	r.POST("/addcart", handlers.AddChart)
	r.POST("/checkout", handlers.Checkout)
	r.Run("localhost:8888")
}
