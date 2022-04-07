package dto

type UserUpdateDTO struct {
	ID    uint64 `json:"id" from:"id" binding:"required"`
	Name  string `json:"name" from:"name" binding:"required"`
	Email string `json:"email" from:"email" binding:"required"`
}
