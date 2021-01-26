package service

import (
	"log"

	"Atlantis-Backend/dto"
	"Atlantis-Backend/models"
	"Atlantis-Backend/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// UserService interface that cover all function needed
type UserService interface {
	LoginUser(email string, password string) interface{}
	RegisterUser(user dto.UserCreateDTO) models.User
	UpdateUser(user dto.UserUpdateDTO) models.User
	GetUser(userID string) models.User
	FindByField(tipe string, val string) models.User
	IsDuplicateField(tipe string, val string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService is used to create new Instance
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func comparePassword(hashed string, pass []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, pass)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (service *userService) LoginUser(email string, pass string) interface{} {
	res := service.userRepository.VerifyCredential(email, pass)
	if val, ok := res.(models.User); ok {
		comparedPass := comparePassword(val.Password, []byte(pass))
		if val.Email == email && comparedPass {
			return res
		}

		return false
	}

	return false
}

func (service *userService) RegisterUser(user dto.UserCreateDTO) models.User {
	newUser := models.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.userRepository.InsertUser(newUser)
	return res
}

func (service *userService) FindByField(tipe string, val string) models.User {
	return service.userRepository.FindField(tipe, val)
}

func (service *userService) IsDuplicateField(tipe string, val string) bool {
	res := service.userRepository.IsDuplicate(tipe, val)

	return !(res.Error == nil)
}

func (service *userService) UpdateUser(user dto.UserUpdateDTO) models.User {
	newUser := models.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	updatedUser := service.userRepository.UpdateUser(newUser)
	return updatedUser
}

func (service *userService) GetUser(id string) models.User {
	return service.userRepository.ProfileUser(id)
}
