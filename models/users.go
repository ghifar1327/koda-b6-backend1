package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Picture   string    `json:"picture"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Role     *string `json:"role"`
	Phone    *string `json:"phone"`
}

// ============================================================================= RESPONSE

type UserResponse struct {
	Id       int    `json:"id"`
	Picture  string `json:"picture"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}