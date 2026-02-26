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
type User struct {
	Id       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var Users []User
var currentID int = 0

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "List of users",
			Results: Users,
		})
	})

	// ================================================================================== Register User
	r.POST("/register", func(ctx *gin.Context) {
		var data User

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
		} else {
			for _, user := range Users {
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
			Users = append(Users, data)

			ctx.JSON(201, Response{
				Success: true,
				Message: "Register successfully",
				Results: data,
			})
		}

	})

	// ======================================================================================================= Login

	r.POST("/login", func(ctx *gin.Context) {
		var userInput User
		var foundUser *User

		if err := ctx.ShouldBindJSON(&userInput); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		for _, data := range Users {
			if data.Email == userInput.Email && data.Password == userInput.Password {
				foundUser = &data
				break
			}
		}

		if foundUser == nil {
			ctx.JSON(401, Response{
				Success: false,
				Message: "Email or password is incorrect",
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: "Login successfully",
			Results: foundUser,
		})
	})

	// ================================================================================================ GET User
	r.POST("/users/", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}
		for _, user := range Users {
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
		var updateUser User
		if err := ctx.ShouldBindJSON(&updateUser); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
		}

		for i, user := range Users {
			if user.Id == id {
				// -----------------------------------------------------------------------------------------------------
				if updateUser.Email != "" && updateUser.Email != user.Email {
					for _, u := range Users {
						if u.Email == updateUser.Email {
							ctx.JSON(400, Response{
								Success: false,
								Message: "Email already exists",
							})
							return
						}
					}
					Users[i].Email = updateUser.Email
				}
				// ---------------------------------------------------------------------------------------------------------
				if updateUser.Password != "" {
					Users[i].Password = updateUser.Password
				}

				ctx.JSON(200, Response{
					Success: true,
					Message: "User updated successfully",
					Results: Users[i],
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

	r.DELETE("user/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}
		var newData []User
		var found bool = false
		for _, user := range Users {
			if user.Id == id {
				found = true
				continue
			}
			newData = append(newData, user)
		}
		if !found {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		Users = newData

		ctx.JSON(200, Response{
			Success: true,
			Message: "User deleted successfully",
		})

	})

	r.Run("localhost:8888")
}
