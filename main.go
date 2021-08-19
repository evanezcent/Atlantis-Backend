package main

import (
	"Atlantis-Backend/config"
	"Atlantis-Backend/controller"
	"Atlantis-Backend/middleware"
	"Atlantis-Backend/repository"
	"Atlantis-Backend/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.InitConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	itemRepository repository.ItemRepository = repository.NewItemRepository(db)
	jwtService     service.JWTService        = service.NewJwtService()
	userService    service.UserService       = service.NewUserService(userRepository)
	itemService    service.ItemService       = service.NewItemService(itemRepository)
	userController controller.UserController = controller.NewUserController(jwtService, userService)
	itemController controller.ItemController = controller.NewItemController(jwtService, itemService)
)

func main() {
	fmt.Println("Starting apps...")
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	authRoutes := r.Group("atlantis-api/v1/user")
	{
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/register", userController.Register)
	}

	userRoutes := r.Group("atlantis-api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/get", userController.Get)
		userRoutes.PUT("/update", userController.Update)
	}

	itemRoutes := r.Group("atlantis-api/v1/item", middleware.AuthorizeJWT(jwtService))
	{
		itemRoutes.POST("/insert", itemController.Add)
		itemRoutes.GET("/all", itemController.All)
		itemRoutes.PUT("/confirm/:id", itemController.Confirm)
		itemRoutes.get("/:id", itemController.Get)
	}

	r.Run()
}
