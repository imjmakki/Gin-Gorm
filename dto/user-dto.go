package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" from:"id" binding:"required"`
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:8"`
}

type UserCreateDTO struct {
	Name     string `json:"name" from:"name" binding:"required"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" binding:"required" validate:"min:8"`
}
