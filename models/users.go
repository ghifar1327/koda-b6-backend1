package models

type User struct {
	Id       int
	Email    string
	Password string
}
// ============================================================ REQUEST
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

// ======================================================== RESPONSE

type UserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}