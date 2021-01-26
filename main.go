package main

import (
	"Atlantis-Backend/config"
	"Atlantis-Backend/controller"
	"Atlantis-Backend/repository"
	"Atlantis-Backend/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.InitConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJwtService()
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(jwtService, userService)
)

func main() {
	fmt.Println("Starting apps...")
	r := gin.Default()

	authRoutes := r.Group("api/v1/user")
	{
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/register", userController.Register)
	}

	userRoutes := r.Group("api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/get", userController.Get)
		userRoutes.PUT("/update", userController.Update)
	}

	r.Run()
}
