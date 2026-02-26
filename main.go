package main

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

type User struct {
	Id       int
	Email    string
	Password string
}

// ============================================================================ REQUEST DTO
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ======================================================================= RESPONSE DTO
type UserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

var Users []User
var currentID int
var argon = argon2.DefaultConfig()

func main() {
	r := gin.Default()

	// ================================= GET ALL USERS
	r.GET("/users", func(ctx *gin.Context) {

		var result []UserResponse
		for _, user := range Users {
			result = append(result, UserResponse{
				Id:    user.Id,
				Email: user.Email,
			})
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: "List of users",
			Results: result,
		})
	})

	// ==================================================================================== REGISTER
	r.POST("/register", func(ctx *gin.Context) {

		var input RegisterInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid email or password format",
			})
			return
		}

		for _, user := range Users {
			if user.Email == input.Email {
				ctx.JSON(400, Response{
					Success: false,
					Message: "Email already exists",
				})
				return
			}
		}

		hash, err := argon.HashEncoded([]byte(input.Password))
		if err != nil {
			ctx.JSON(500, Response{
				Success: false,
				Message: "Failed to hash password",
			})
			return
		}

		currentID++

		newUser := User{
			Id:       currentID,
			Email:    input.Email,
			Password: string(hash),
		}

		Users = append(Users, newUser)

		ctx.JSON(201, Response{
			Success: true,
			Message: "Register successfully",
			Results: UserResponse{
				Id:    newUser.Id,
				Email: newUser.Email,
			},
		})
	})

	// ============================================================================================= LOGIN
r.POST("/login", func(ctx *gin.Context) {

	var input LoginInput
	var foundUser *User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{
			Success: false,
			Message: "Invalid email or password format",
		})
		return
	}

	for i, user := range Users {
		if user.Email == input.Email {

			match, err := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
			if err == nil && match {
				foundUser = &Users[i]
				break
			}
		}
	}

	if foundUser == nil {
		ctx.JSON(401, Response{
			Success: false,
			Message: "Email or password incorrect",
		})
		return
	}

	ctx.JSON(200, Response{
		Success: true,
		Message: "Login successfully",
		Results: UserResponse{
			Id:    foundUser.Id,
			Email: foundUser.Email,
		},
	})
})

	// ==================================================================================== GET USER
	r.GET("/users/:id", func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
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
					Message: fmt.Sprintf("Welcome %s", user.Email),
					Results: UserResponse{
						Id:    user.Id,
						Email: user.Email,
					},
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// ================================= UPDATE USER
	r.PATCH("/users/:id", func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}

		var input UpdateInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		for i, user := range Users {
			if user.Id == id {

				if input.Email != "" {
					Users[i].Email = input.Email
				}

				if input.Password != "" {
					hash, err := argon.HashEncoded([]byte(input.Password))
					if err != nil {
						ctx.JSON(500, Response{
							Success: false,
							Message: "Failed to hash password",
						})
						return
					}
					Users[i].Password = string(hash)
				}

				ctx.JSON(200, Response{
					Success: true,
					Message: "User updated successfully",
					Results: UserResponse{
						Id:    Users[i].Id,
						Email: Users[i].Email,
					},
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// ============================================================================= DELETE
	r.DELETE("/users/:id", func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}

		var newUsers []User
		found := false

		for _, user := range Users {
			if user.Id == id {
				found = true
				continue
			}
			newUsers = append(newUsers, user)
		}

		if !found {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User not found",
			})
			return
		}

		Users = newUsers

		ctx.JSON(200, Response{
			Success: true,
			Message: "User deleted successfully",
		})
	})

	r.Run("localhost:8888")
}