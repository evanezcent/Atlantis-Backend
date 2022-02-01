package controller

import (
	"Atlantis-Backend/models"
	"fmt"
	"net/http"
	"strconv"

	"Atlantis-Backend/dto"
	"Atlantis-Backend/helper"
	"Atlantis-Backend/service"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ItemController interface for login, register, read, and update user
type ItemController interface {
	Add(ctx *gin.Context)
	Confirm(ctx *gin.Context)
	Update(ctx *gin.Context)
	All(ctx *gin.Context)
	Get(ctx *gin.Context)
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
	// validate request input using DTO
	var newItem dto.ItemCreateDTO
	err := ctx.ShouldBind(&newItem)

	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	// get header for authorization
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)

	// validate user token
	if errToken != nil {
		panic(errToken.Error())
	}

	// get id from token
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userID"])

	// insert data into database
	newItem.UserID = id
	successItem := c.itemService.Insert(newItem)

	// init for multipart form
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		panic(errForm.Error())
	}

	// check is there is an images
	files, errFile := form.File["images"]
	if !errFile {
		res := helper.ResponseFailed("Null images", "Failed to upload images", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	var listImage []string

	// Loop for images upload
	for _, file := range files {
		// get file extension
		var extension = filepath.Ext(file.Filename)

		// generate random name
		filename := helper.RandomString(11) + extension

		// make the path for upload
		name := "uploads/" + filename
		path := name

		// save file name into folder path
		if err := ctx.SaveUploadedFile(file, path); err != nil {
			panic(err.Error())
		}

		// validate image
		var image dto.ItemImageCreateDTO
		image.ItemID = strconv.FormatUint(successItem.ID, 10)
		image.URL = "localhost:8080/" + name

		// upload the image
		res := c.itemService.UploadImage(image)
		if (res == models.ImageItem{}) {
			res := helper.ResponseFailed("Failed to upload image", "Failed", nil)
			ctx.JSON(http.StatusConflict, res)
		}

		// save the result path of the image
		listImage = append(listImage, path)
	}

	// send success response data
	var data Result
	data.obj = successItem
	data.images = listImage

	response := helper.ResponseSucces(true, "success", data)
	ctx.JSON(http.StatusOK, response)
}

func (c *itemController) Confirm(ctx *gin.Context) {
	// get header for authorization
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)

	// validate user token
	if errToken != nil && token == nil {
		panic(errToken.Error())
	}

	// get params
	itemID := ctx.Param("id")
	res := c.itemService.ConfirmItem(itemID)

	response := helper.ResponseSucces(true, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *itemController) Update(ctx *gin.Context) {

}

func (c *itemController) All(ctx *gin.Context) {
	var items []models.Combined

	// get header for authorization
	authHeader := ctx.GetHeader("Authorization")

	// validate user token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	// check is request has query
	userID := ctx.Query("user_id")
	query := ctx.Query("q")

	if userID != "" {
		// parse id into uint 64
		uid, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
		items = c.itemService.GetByUser(uid)
	} else if query != "" {
		items = c.itemService.GetByQuery(query)
	} else {
		items = c.itemService.GetAll()
	}

	res := helper.ResponseSucces(true, "success", items)
	ctx.JSON(http.StatusOK, res)
}

func (c *itemController) Get(ctx *gin.Context) {
	// get header for authorization
	authHeader := ctx.GetHeader("Authorization")

	// validate user token
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	// parse id into uint 64
	itemID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	items := c.itemService.Get(itemID)
	fmt.Println(items)
	res := helper.ResponseSucces(true, "success", items)
	ctx.JSON(http.StatusOK, res)
}
