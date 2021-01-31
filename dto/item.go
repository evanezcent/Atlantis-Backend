package dto

// ItemCreateDTO is used to catch body json from client
type ItemCreateDTO struct {
	Title         string `json:"title" form:"title" binding:"required"`
	SpesificDate  string `form:"spesific_date" json:"spesific_date" binding:"required"`
	SpesificPlace string `form:"spesific_place" json:"spesific_place" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
	UserID        string `json:"user_id" form:"user_id"`
	IsDone        bool   `json:"is_done" form:"is_done"`
}

// ItemUpdateDTO is used to catch body json from client
type ItemUpdateDTO struct {
	Title         string `json:"title" form:"title" binding:"required"`
	SpesificDate  string `form:"spesific_date" json:"spesific_date" binding:"required"`
	SpesificPlace string `form:"spesific_place" json:"spesific_place" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
}
