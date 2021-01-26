package dto

// UserUpdateDTO is used to catch body json from client
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required" validate:"min:6"`
	Image    string `json:"image" form:"image"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

// UserCreateDTO is used to catch body json from client
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required" validate:"min:6"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
	Image    string `json:"image" form:"image"`
}

// LoginDTO is used to catch body json from client
type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}
