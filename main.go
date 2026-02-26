package main

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}
type Users struct {
	Id       int    `json:"id form:id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var listener []Users
var currentID int = 0

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "backend is running well",
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		var data Users

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
		} else {
			for _ , user := range listener{
				if user.Email == data.Email{
					ctx.JSON(400, Response{
						Success: false,
						Message: "email, is already exist",
					})
					return
				}
			}
			currentID++
			data.Id = currentID
			listener = append(listener, data)

			ctx.JSON(201, Response{
				Success: true,
				Message: "User created successfully",
				Results: data,
			})
		}

	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "List of users",
			Results: listener,
		})
	})
	// r.GET("/users:id",func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	if id == "5"{
	// 		ctx.JSON(200, Response{
	// 			Success: true,

	// 		})

	// 	}
	// })
	r.Run("localhost:8888")
}
