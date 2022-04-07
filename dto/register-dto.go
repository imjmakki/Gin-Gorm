package dto

type RegisterDTO struct {
	Name     string `json:"name" from:"name" binding:"required" validate:"min:4"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:8"`
}
