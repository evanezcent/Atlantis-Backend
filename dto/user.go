package dto

// UserUpdateDTO is used to catch body json from client
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Phone    int8   `json:"phone" form:"phone" binding:"required" validate:"min:6"`
	Name     string `json:"name" form:"name" binding:"required"`
	Image    string `json:"image" form:"image" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

// UserCreateDTO is used to catch body json from client
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
	Phone    uint64 `json:"phone" form:"phone" binding:"required" validate:"min:6"`
	Image    string `json:"image" form:"image" binding:"required"`
}

// LoginDTO is used to catch body json from client
type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}
