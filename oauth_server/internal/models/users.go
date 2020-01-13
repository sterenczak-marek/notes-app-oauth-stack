package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server/config"
	"golang.org/x/crypto/bcrypt"
)

const userTable = "users"

type User struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email" gorm:"unique_index;not null" `
	Password  string    `json:"-" gorm:"not null" `
}

var InvalidCredentialsError = fmt.Errorf("invalid credentials")

func init() {
	db := config.GormDB
	db.Debug().AutoMigrate(&User{})
}

//Validate incoming user details...
func (user *User) Validate() error {

	if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("email address is required")
	}

	if userPasswordLength := len(user.Password); userPasswordLength == 0 {
		return fmt.Errorf("password is required")
	} else if userPasswordLength < 4 {
		return fmt.Errorf("password is too short")
	}

	//Email must be unique
	temp := &User{}
	err := config.GormDB.Table(userTable).Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("connection error. Please retry")
	}
	if temp.Email != "" {
		return fmt.Errorf("email address already in use by another user")
	}
	return nil
}

func (user *User) Create() (err error) {
	if err = user.Validate(); err != nil {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.UUID = uuid.New()

	err = config.GormDB.Create(&user).Error
	if err != nil || user.ID <= 0 {
		return fmt.Errorf("failed to create account, connection error: %s", err)
	}
	return nil
}

func GetUserByEmail(email string) *User {
	user := new(User)
	err := config.GormDB.Table(userTable).Where("Email = ?", email).First(user).Error
	if err != nil || user.Email == "" {
		return nil
	}
	return user
}

func GetUserByID(userID string) *User {
	user := new(User)
	err := config.GormDB.Table(userTable).Where("UUID = ?", userID).First(user).Error
	if err != nil || user.Email == "" {
		return nil
	}
	return user
}

func Authenticate(email, password string) (*User, error) {
	user := GetUserByEmail(email)
	if user == nil {
		return nil, InvalidCredentialsError
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, InvalidCredentialsError
	}

	return user, nil
}
