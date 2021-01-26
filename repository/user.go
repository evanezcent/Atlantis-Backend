package repository

import (
	"Atlantis-Backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository as interface that cover all function
type UserRepository interface {
	InsertUser(iser models.User) models.User
	UpdateUser(iser models.User) models.User
	VerifyCredential(email string, pass string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindEmail(email string) models.User
	ProfileUser(id string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepository used to create new Instance of user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = hashPassword([]byte(user.Password))
	db.connection.Save(&user)

	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashPassword([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)

	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User

	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User

	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)

	return user
}

func (db *userConnection) ProfileUser(id string) models.User {
	var user models.User
	db.connection.Find(&user, id)

	return user
}

func hashPassword(pwd []byte) string {
	hashh, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash the password")
	}

	return string(hashh)
}
