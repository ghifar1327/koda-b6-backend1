package handlers

import (
	"backend/models"
	"backend/utils"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var Users []models.User

// var currentID int
var mu sync.Mutex

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

func textToPtr(t pgtype.Text) *string {
	if t.Valid {
		return &t.String
	}
	return nil
}

// =============================================================================================================== GET ALL USERS
func GetUsers(ctx *gin.Context) {

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL:", dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Println("CONNECT ERROR:", err)
		ctx.JSON(500, Response{false, "Failed to connect database", nil})
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT id, full_name, picture, email,password ,role_id ,phone, address, created_at, updated_at FROM users`)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
		ctx.JSON(500, Response{false, "Failed to query users", nil})
		return
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		fmt.Println("COLLECT ERROR:", err)
		ctx.JSON(500, Response{false, "Failed to collect users", nil})
		return
	}

	ctx.JSON(200, Response{true, "List of users", users})
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
		ctx.JSON(400, Response{false, "Failed to hash password", nil})
		return
	}

	// currentID++

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

	ctx.JSON(201, Response{
		true,
		"Register successfully",
		models.UserResponse{
			Id:       newUser.Id,
			Picture:  *textToPtr(newUser.Picture),
			FullName: newUser.FullName,
			Email:    newUser.Email,
			RoleId:   newUser.RoleId,
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
				ctx.JSON(400, Response{false, "Failed to verify password", nil})
				return
			}

			if match {
				ctx.JSON(200, Response{
					true,
					"Login successfully",
					models.UserResponse{
						Id:       user.Id,
						Picture:  *textToPtr(user.Picture),
						FullName: user.FullName,
						Email:    user.Email,
						RoleId:   user.RoleId,
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

	id, err := uuid.Parse(ctx.Param("id"))
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
					Picture:  *textToPtr(user.Picture),
					FullName: user.FullName,
					Email:    user.Email,
					RoleId:   user.RoleId,
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
				hash, err := utils.HashPassword(*input.Password)
				if err != nil {
					ctx.JSON(500, Response{false, "Failed to hash password", nil})
					return
				}
				Users[i].Password = hash
			}

			if input.Picture != nil {
				Users[i].Picture = pgtype.Text{
					String: *input.Picture,
					Valid:  true,
				}
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

			ctx.JSON(200, Response{
				true,
				"User updated successfully",
				models.UserResponse{
					Id:       Users[i].Id,
					Picture:  *textToPtr(Users[i].Picture),
					FullName: Users[i].FullName,
					Email:    Users[i].Email,
					RoleId:   Users[i].RoleId,
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
	id, err := uuid.Parse(ctx.Param("id"))
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
