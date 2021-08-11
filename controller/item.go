package controller

import (
	"Atlantis-Backend/models" 
	"fmt"
	"net/http"

	"Atlantis-Backend/dto"
	"Atlantis-Backend/helper"
	"Atlantis-Backend/service"

	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ItemController interface for login, register, read, and update user
type ItemController interface {
	Add(ctx *gin.Context)
	Confirm(ctx *gin.Context)
	Update(ctx *gin.Context)
	All(ctx *gin.Context)
}

type itemController struct {
	jwtService  service.JWTService
	itemService service.ItemService
}

// Result is to handle insert and update
type Result struct {
	obj    models.Item `json:"item"`
	images []string    `json:"images"`
}

// NewItemController is like constructor of the models
func NewItemController(jwtService service.JWTService, itemService service.ItemService) ItemController {
	return &itemController{
		jwtService,
		itemService,
	}
}

func (c *itemController) Add(ctx *gin.Context) {
	var newItem dto.ItemCreateDTO
	err := ctx.ShouldBind(&newItem)

	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userID"])

	newItem.UserID = id
	successItem := c.itemService.Insert(newItem)

	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		panic(errForm.Error())
	}

	files, errFile := form.File["images"]
	if !errFile {
		res := helper.ResponseFailed("Null images", "Failed to upload images", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}
	var listImage []string

	// Loop for images upload
	for _, file := range files {
		var extension = filepath.Ext(file.Filename)
		filename := helper.RandomString(11) + extension
		name := "uploads/" + filename
		path := name

		if err := ctx.SaveUploadedFile(file, path); err != nil {
			panic(err.Error())
		}

		var image dto.ItemImageCreateDTO
		image.ItemID = id
		image.URL = "localhost:8080/" + name
		res := c.itemService.UploadImage(image)
		if (res == models.ImageItem{}) {
			res := helper.ResponseFailed("Failed to upload image", "Failed", nil)
			ctx.JSON(http.StatusConflict, res)
		}

		listImage = append(listImage, path)
	}

	var data Result
	data.obj = successItem
	data.images = listImage 

	response := helper.ResponseSucces(true, "success", data)
	ctx.JSON(http.StatusOK, response)
}

func (c *itemController) Confirm(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil && token == nil {
		panic(errToken.Error())
	}

	itemID := ctx.Param("id")
	res := c.itemService.ConfirmItem(itemID)
	response := helper.ResponseSucces(true, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *itemController) Update(ctx *gin.Context) {

}

func (c *itemController) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	items := c.itemService.GetAll()  
	res := helper.ResponseSucces(true, "success", items)
	ctx.JSON(http.StatusOK, res)
}
