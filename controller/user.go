package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"Atlantis-Backend/dto"
	"Atlantis-Backend/helper"
	"Atlantis-Backend/models"
	"Atlantis-Backend/service"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// UserController interface for login, register, read, and update user
type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type userController struct {
	jwtService  service.JWTService
	userService service.UserService
}

// NewUserController is like constructor of the models
func NewUserController(jwtService service.JWTService, userService service.UserService) UserController {
	return &userController{
		jwtService,
		userService,
	}
}

func (c *userController) Login(ctx *gin.Context) {
	// validate request input using DTO
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBind(&loginDTO)

	// if validation error
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	// check data into database
	authRes := c.userService.LoginUser(loginDTO.Email, loginDTO.Password)

	if val, ok := authRes.(models.User); ok {
		// if success generate user token
		// using user id
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(val.ID, 10))
		val.Token = generateToken

		res := helper.ResponseSucces(true, "success", val)

		ctx.JSON(http.StatusOK, res)

		return
	}

	response := helper.ResponseFailed("Invalid Credential", "Invalid Credential", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *userController) Register(ctx *gin.Context) {
	// validate input
	var newUser dto.UserCreateDTO
	err := ctx.ShouldBind(&newUser)

	// if valdiation error
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	if !c.userService.IsDuplicateField("email", newUser.Email) {
		// if email already exist
		res := helper.ResponseFailed("Email has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else if !c.userService.IsDuplicateField("phone", newUser.Phone) {
		// if phone number already exist
		res := helper.ResponseFailed("Phone has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else {
		// register new user
		createUser := c.userService.RegisterUser(newUser)

		// generate the token using USER ID
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token

		res := helper.ResponseSucces(true, "success", createUser)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) Update(ctx *gin.Context) {
	// valdiate input
	var newUser dto.UserUpdateDTO
	err := ctx.ShouldBind(&newUser)

	// if validation error
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	// get authentication header
	authHeader := ctx.GetHeader("Authorization")

	// validate user token
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	// get user id from token
	claims := token.Claims.(jwt.MapClaims)
	id, errID := strconv.ParseUint(fmt.Sprintf("%v", claims["userID"]), 10, 64)
	if errID != nil {
		panic(errToken.Error())
	}

	// check the user email and phone
	userEmail := c.userService.FindByField("email", newUser.Email)
	userPhone := c.userService.FindByField("phone", newUser.Phone)

	if (userEmail.ID != id && userEmail != models.User{}) {
		// if email already used by other user
		fmt.Println(userEmail.ID, id)
		res := helper.ResponseFailed("Email has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else if (userPhone.ID != id && userPhone != models.User{}) {
		// if phone number already used by other user
		fmt.Println(userEmail.ID, id)
		res := helper.ResponseFailed("Phone has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else {
		// if request has images
		file, _, err := ctx.Request.FormFile("images")
		if file != nil && err == nil {
			// upload user iamge
			imgURL := upload(ctx)
			newUser.Image = imgURL
		}

		// update user by id
		newUser.ID = id
		updateUser := c.userService.UpdateUser(newUser)
		res := helper.ResponseSucces(true, "success", updateUser)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) Get(ctx *gin.Context) {
	// get authorization header
	authHeader := ctx.GetHeader("Authorization")

	// validate token
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	// get usr id from token
	claims := token.Claims.(jwt.MapClaims)

	// get data user depend the ID
	user := c.userService.GetUser(fmt.Sprintf("%v", claims["userID"]))
	res := helper.ResponseSucces(true, "success", user)
	ctx.JSON(http.StatusOK, res)
}

func upload(ctx *gin.Context) string {
	// is request has images
	file, errForm := ctx.FormFile("images")
	if errForm != nil {
		panic(errForm.Error())
	}

	// get file extension
	var extension = filepath.Ext(file.Filename)

	// generate random sttring
	filename := helper.RandomString(11) + extension
	name := "uploads/" + filename

	// make path
	path := name

	// save into local folder
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		panic(err.Error())
	}

	return path
	// res := helper.ResponseSucces(true, "success", path)
	// ctx.JSON(http.StatusOK, res)
}

// func (c *userController) Upload(ctx *gin.Context) {
// 	form, errForm := ctx.MultipartForm()
// 	if errForm != nil {
// 		panic(errForm.Error())
// 	}

// 	files := form.File["images"]
// 	fmt.Println(filepath.Dir("/uploads/"))

// 	var listPath []string

// 	for _, file := range files {
// 		var extension = filepath.Ext(file.Filename)
// 		filename := helper.RandomString(11) + extension
// 		name := "uploads/" + filename
// 		fmt.Println(filename)
// 		path := name
// 		if err := ctx.SaveUploadedFile(file, path); err != nil {
// 			panic(err.Error())
// 		}
// 		listPath = append(listPath, path)
// 	}
// 	res := helper.ResponseSucces(true, "success", listPath)
// 	ctx.JSON(http.StatusOK, res)
// }
