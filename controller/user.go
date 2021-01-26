package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"Atlantis-Backend/dto"
	"Atlantis-Backend/helper"
	"Atlantis-Backend/models"
	"Atlantis-Backend/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBind(&loginDTO)

	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	authRes := c.userService.LoginUser(loginDTO.Email, loginDTO.Password)
	if val, ok := authRes.(models.User); ok {
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
	var newUser dto.UserCreateDTO
	err := ctx.ShouldBind(&newUser)
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	if !c.userService.IsDuplicateField("email", newUser.Email) {
		res := helper.ResponseFailed("Email has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else if !c.userService.IsDuplicateField("phone", newUser.Phone) {
		res := helper.ResponseFailed("Phone has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else {
		createUser := c.userService.RegisterUser(newUser)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token

		res := helper.ResponseSucces(true, "success", createUser)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var newUser dto.UserUpdateDTO
	err := ctx.ShouldBind(&newUser)
	if err != nil {
		res := helper.ResponseFailed("Failed to process", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, errID := strconv.ParseUint(fmt.Sprintf("%v", claims["userID"]), 10, 64)
	if errID != nil {
		panic(errToken.Error())
	}

	if !c.userService.IsDuplicateField("email", newUser.Email) {
		res := helper.ResponseFailed("Email has ben registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else if !c.userService.IsDuplicateField("phone", newUser.Phone) {
		res := helper.ResponseFailed("Phone has been registered", "Failed", nil)
		ctx.JSON(http.StatusConflict, res)
	} else {
		newUser.ID = id
		updateUser := c.userService.UpdateUser(newUser)
		res := helper.ResponseSucces(true, "success", updateUser)
		ctx.JSON(http.StatusOK, res)
	}

}

func (c *userController) Get(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.GetUser(fmt.Sprintf("%v", claims["userID"]))
	res := helper.ResponseSucces(true, "success", user)
	ctx.JSON(http.StatusOK, res)
}
