package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" from:"id" binding:"required"`
	Title       uint64 `json:"title" from:"title" binding:"required"`
	Description uint64 `json:"description" from:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}

type BookCreateDTO struct {
}
