package handlers

import (
	"backend/models"
	"backend/utils"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var Users []models.User
var currentID int
var mu sync.Mutex

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

// =============================================================================================================== GET ALL USERS
func GetUsers(ctx *gin.Context) {

	defer mu.Unlock()
	mu.Lock()

	var result []models.User

	for _, user := range Users {
		result = append(result, models.User{
			Id:        user.Id,
			Picture:   user.Picture,
			FullName:  user.FullName,
			Email:     user.Email,
			Role:      user.Role,
			Address:   user.Address,
			Phone:     user.Phone,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		})
	}

	ctx.JSON(200, Response{true, "List of users", result})
}

// ============================================================================================================== REGISTER
func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid email or password", nil})
		return
	}

	mu.Lock()
	defer mu.Unlock()

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
		Id:        currentID,
		Email:     input.Email,
		Password:  hash,
		FullName:  input.FullName,
		Role:      "user",
		Address:   input.Address,
		Phone:     input.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Users = append(Users, newUser)

	ctx.JSON(201, Response{
		true,
		"Register successfully",
		models.UserResponse{
			Id:       newUser.Id,
			Picture:  newUser.Picture,
			FullName: newUser.FullName,
			Email:    newUser.Email,
			Role:     newUser.Role,
			Address:  newUser.Address,
			Phone:    newUser.Phone,
		},
	})
}

// ================================================================================================================ LOGIN
func Login(ctx *gin.Context) {
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid request body", nil})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, user := range Users {
		if user.Email == input.Email {

			match, err := utils.VerifyPassword(input.Password, user.Password)
			if err != nil {
				ctx.JSON(500, Response{false, "Failed to verify password", nil})
				return
			}

			if match {
				ctx.JSON(200, Response{
					true,
					"Login successfully",
					models.UserResponse{
						Id:       user.Id,
						Picture:  user.Picture,
						FullName: user.FullName,
						Email:    user.Email,
						Role:     user.Role,
						Address:  user.Address,
						Phone:    user.Phone,
					},
				})
				return
			}
		}
	}

	ctx.JSON(401, Response{false, "Email or password incorrect", nil})
}

// ======================================================================================================= GET USER BY ID
func GetUserByID(ctx *gin.Context) {

	defer mu.Unlock()
	mu.Lock()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}
	for _, user := range Users {
		if user.Id == id {
			ctx.JSON(200, Response{
				true,
				"User found",
				models.UserResponse{
					Id:       user.Id,
					Picture:  user.Picture,
					FullName: user.FullName,
					Email:    user.Email,
					Role:     user.Role,
				},
			})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

// ============================================================================================================= UPDATE USER
func UpdateUser(ctx *gin.Context) {

	defer mu.Unlock()
	mu.Lock()

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

			if input.Email != nil {
				for _, u := range Users {
					if u.Email == *input.Email && u.Id != id {
						ctx.JSON(400, Response{false, "Email already exists", nil})
						return
					}
				}
				Users[i].Email = *input.Email
			}

			if input.Password != nil {
				hash, err := utils.HashPassword(*input.Password)
				if err != nil {
					ctx.JSON(500, Response{false, "Failed to hash password", nil})
					return
				}
				Users[i].Password = hash
			}

			if input.Picture != nil {
				Users[i].Picture = *input.Picture
			}

			if input.FullName != nil {
				Users[i].FullName = *input.FullName
			}

			if input.Address != nil {
				Users[i].Address = *input.Address
			}

			if input.Phone != nil {
				Users[i].Phone = *input.Phone
			}

			if input.Role != nil {
				Users[i].Role = *input.Role
			}

			Users[i].UpdatedAt = time.Now()

			ctx.JSON(200, Response{
				true,
				"User updated successfully",
				models.UserResponse{
					Id:       Users[i].Id,
					Picture:  Users[i].Picture,
					FullName: Users[i].FullName,
					Email:    Users[i].Email,
					Role:     Users[i].Role,
					Address:  Users[i].Address,
					Phone:    Users[i].Phone,
				},
			})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

// ======================================================================================================= DELETE USER
func DeleteUser(ctx *gin.Context) {

	defer mu.Unlock()
	mu.Lock()

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
	ctx.JSON(200, Response{true, "User deleted successfully", nil})
}
