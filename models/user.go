package models

import "github.com/jinzhu/gorm"

import "strings"

import "errors"

import "github.com/badoux/checkmail"

//User model
type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null;" json:"name"`
	Email    string `gorm:"size:100;not_null;unique_index" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}

//GetAllUsers function
func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// Validate user cation
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "add":
		if u.Name == "" {
			return errors.New("Name is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email format")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	default:
		return nil
	}
}

//AddUser func
func (u *User) AddUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}
