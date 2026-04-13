package handlers

import (
	"backend/models"
	"backend/utils"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Users []models.User
var mu sync.Mutex

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

// =============================== GET ALL USERS
// @Summary Get all users
// @Description Get list of users (in-memory)
// @Tags users
// @Produce json
// @Success 200 {object} Response
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, Response{true, "List of users", Users})
}

// =============================== REGISTER
// @Summary Register new user
// @Description Create new user (in-memory)
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.RegisterInput true "Register Data"
// @Success 201 {object} Response
// @Router /register [post]
func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, Response{false, "Invalid request", nil})
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

	newUser := models.User{
		Id:        uuid.New(),
		Email:     input.Email,
		Password:  hash,
		FullName:  input.FullName,
		RoleId:    2,
		Address:   input.Address,
		Phone:     input.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Users = append(Users, newUser)

	ctx.JSON(201, Response{true, "Register successfully", models.UserResponse{
		Id:       newUser.Id,
		FullName: newUser.FullName,
		Email:    newUser.Email,
		RoleId:   newUser.RoleId,
		Address:  newUser.Address,
		Phone:    newUser.Phone,
	}})
}

// =============================== LOGIN
// @Summary Login user
// @Description Login using email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.LoginInput true "Login Data"
// @Success 200 {object} Response
// @Router /login [post]
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
			match, _ := utils.VerifyPassword(input.Password, user.Password)
			if match {
				ctx.JSON(200, Response{true, "Login successfully", models.UserResponse{
					Id:       user.Id,
					FullName: user.FullName,
					Email:    user.Email,
					RoleId:   user.RoleId,
					Address:  user.Address,
					Phone:    user.Phone,
				}})
				return
			}
		}
	}

	ctx.JSON(401, Response{false, "Email or password incorrect", nil})
}

// =============================== GET USER BY ID
// @Summary Get user by ID
// @Description Get single user by UUID (in-memory)
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Router /users/{id} [get]
func GetUserByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}

	for _, user := range Users {
		if user.Id == id {
			ctx.JSON(200, Response{true, "User found", user})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

// =============================== UPDATE USER
// @Summary Update user
// @Description Update user by UUID (in-memory)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param input body models.UpdateInput true "Update Data"
// @Success 200 {object} Response
// @Router /updateuser/{id} [patch]
func UpdateUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := uuid.Parse(ctx.Param("id"))
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
				hash, _ := utils.HashPassword(*input.Password)
				Users[i].Password = hash
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

			if input.RoleId != nil {
				Users[i].RoleId = *input.RoleId
			}

			Users[i].UpdatedAt = time.Now()

			ctx.JSON(200, Response{true, "User updated successfully", models.UserResponse{
				Id:       Users[i].Id,
				FullName: Users[i].FullName,
				Email:    Users[i].Email,
				RoleId:   Users[i].RoleId,
				Address:  Users[i].Address,
				Phone:    Users[i].Phone,
			}})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}

// =============================== DELETE USER
// @Summary Delete user
// @Description Delete user by UUID (in-memory)
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Router /deleteuser/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, Response{false, "Invalid ID", nil})
		return
	}

	for i, user := range Users {
		if user.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			ctx.JSON(200, Response{true, "User deleted successfully", nil})
			return
		}
	}

	ctx.JSON(404, Response{false, "User not found", nil})
}
