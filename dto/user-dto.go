package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" bindding:"required"`
	Email    string `json:"email" form:"email" bindding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
