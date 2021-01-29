package dto

// ItemImageCreateDTO is used to catch body json from client
type ItemImageCreateDTO struct {
	URL    string `json:"url" form:"url" binding:"required"`
	ItemID string `form:"itemID" json:"itemID" binding:"required"`
}
