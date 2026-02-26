package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}
type Users struct {
	Id       int    `json:"id" form:"id"`
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
	// ================================================================================== Create User
	r.POST("/users", func(ctx *gin.Context) {
		var data Users

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
		} else {
			for _, user := range listener {
				if user.Email == data.Email {
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

	// ================================================================================================ GET User

	r.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}
		for _, user := range listener {
			if user.Id == id {
				ctx.JSON(200, Response{
					Success: true,
					Message: fmt.Sprintf("Wellcome %s", user.Email),
					Results: user,
				})
				return
			}
		}
		ctx.JSON(400, Response{
			Success: false,
			Message: "User not found",
		})
	})
	// ======================================================================================== Update User

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}
		var updateUser Users
		if err := ctx.ShouldBindJSON(&updateUser); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
		}

		for i, user := range listener {
			if user.Id == id {
				// -----------------------------------------------------------------------------------------------------
				if updateUser.Email != "" && updateUser.Email != user.Email {
					for _, u := range listener {
						if u.Email == updateUser.Email {
							ctx.JSON(400, Response{
								Success: false,
								Message: "Email already exists",
							})
							return
						}
					}
					listener[i].Email = updateUser.Email
				}
				// ---------------------------------------------------------------------------------------------------------
				if updateUser.Password != "" {
					listener[i].Password = updateUser.Password
				}

				ctx.JSON(200, Response{
					Success: true,
					Message: "User updated successfully",
					Results: listener[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// ================================================================================================================================== delete user
	r.Run("localhost:8888")
}
