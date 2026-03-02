package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id        uuid.UUID   `json:"id"`
	Picture   pgtype.Text `json:"picture"`
	FullName  string      `json:"full_name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Address   string      `json:"address"`
	Phone     string      `json:"phone"`
	RoleId    int         `json:"role_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// ========================================================================= REQUEST

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateInput struct {
	Email    *string `json:"email"`
	Picture  *string `json:"picture"`
	FullName *string `json:"full_name"`
	Password *string `json:"password"`
	Address  *string `json:"address"`
	RoleId   *int    `json:"role_id"`
	Phone    *string `json:"phone"`
}

// ============================================================================= RESPONSE

// Picture pgtype.Text

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Picture  string    `json:"picture"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	RoleId   int       `json:"role_id"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
}
