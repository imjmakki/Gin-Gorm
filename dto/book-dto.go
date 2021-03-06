package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" from:"id" binding:"required"`
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"description" from:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" from:"title" binding:"required"`
	Description string `json:"description" from:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}
