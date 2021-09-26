package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" bindding:"required"`
	Title       string `json:"title" form:"title" bindding:"required"`
	Description string `json:"description" form:"description" bindding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" bindding:"required"`
	Description string `json:"description" form:"description" bindding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
