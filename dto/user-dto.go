package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" bindding:"required"`
	Name     string `json:"name" form:"name" bindding:"required"`
	Email    string `json:"email" form:"email" bindding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

type UserCreateDTO struct {
	ID       uint64 `json:"id" form:"id" bindding:"required"`
	Name     string `json:"name" form:"name" bindding:"required"`
	Email    string `json:"email" form:"email" bindding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" bindding:"required"`
}
