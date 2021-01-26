package repository

import (
	"Atlantis-Backend/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository as interface that cover all function
type UserRepository interface {
	InsertUser(iser model.User) model.User
	UpdateUser(iser model.User) model.User
	VerifyCredential(email string, pass string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindEmail(email string) model.User
	ProfileUser(id string) model.User
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepo used to create new Instance of user repository
func NewUserRepo(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.User) model.User {
	user.Password = hashPassword([]byte(user.Password))
	db.connection.Save(&user)

	return user
}

func (db *userConnection) UpdateUser(user model.User) model.User {
	if user.Password != "" {
		user.Password = hashPassword([]byte(user.Password))
	} else {
		var tempUser model.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)

	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user model.User

	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User

	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)

	return user
}

func (db *userConnection) ProfileUser(id string) model.User {
	var user model.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, id)

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
