package handlers

import (
	"demo/models"
	"demo/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Users []models.User
var currentID int

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

// ================================================================================================================ GET USERS
func GetUsers(ctx *gin.Context) {
	var result []models.UserResponse

	for _, user := range Users {
		result = append(result, models.UserResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}

	ctx.JSON(200, Response{
		Success: true,
		Message: "List of users",
		Results: result,
	})
}

// =========================================================================================================================== REGISTER
func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid email or password format", nil})
		return
	}

	for _, user := range Users {
		if user.Email == input.Email {
			ctx.JSON(400, Response{false, "Email already exists", nil})
			return
		}
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		ctx.JSON(500, Response{false, "Failed to hash password", nil})
		return
	}

	currentID++

	newUser := models.User{
		Id:       currentID,
		Email:    input.Email,
		Password: hash,
	}

	Users = append(Users, newUser)

	ctx.JSON(201, Response{
		Success: true,
		Message: "Register successfully",
		Results: models.UserResponse{
			Id:    newUser.Id,
			Email: newUser.Email,
		},
	})
}

// =========================================================================================================================== LOGIN
func Login(ctx *gin.Context) {
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid email or password format", nil})
		return
	}

	for _, user := range Users {
		if user.Email == input.Email {

			match, _ := utils.VerifyPassword(input.Password, user.Password)
			if match {
				ctx.JSON(200, Response{
					Success: true,
					Message: "Login successfully",
					Results: models.UserResponse{
						Id:    user.Id,
						Email: user.Email,
					},
				})
				return
			}
		}
	}

	ctx.JSON(401, Response{false, "Email or password incorrect", nil})
}

// ====================================================================================================================== GET USER BY ID
func GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}

	for _, user := range Users {
		if user.Id == id {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("Welcome %s", user.Email),
				Results: models.UserResponse{
					Id:    user.Id,
					Email: user.Email,
				},
			})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

// ========================================================================================================================== UPDATE USER
func UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}

	var input models.UpdateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid request body", nil})
		return
	}

	for i, user := range Users {
		if user.Id == id {

			if input.Email != "" {
				for _, u := range Users {
					if u.Email == input.Email && u.Id != id {
						ctx.JSON(400, Response{false, "Email already exists", nil})
						return
					}
				}
				Users[i].Email = input.Email
			}

			if input.Password != "" {
				hash, err := utils.HashPassword(input.Password)
				if err != nil {
					ctx.JSON(500, Response{false, "Failed to hash password", nil})
					return
				}
				Users[i].Password = hash
			}

			ctx.JSON(200, Response{
				Success: true,
				Message: "User updated successfully",
				Results: models.UserResponse{
					Id:    Users[i].Id,
					Email: Users[i].Email,
				},
			})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}

	var newData []models.User
	found := false

	for _, user := range Users {
		if user.Id == id {
			found = true
			continue
		}
		newData = append(newData, user)
	}

	if !found {
		ctx.JSON(404, Response{false, "User not found", nil})
		return
	}

	Users = newData

	ctx.JSON(200, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
