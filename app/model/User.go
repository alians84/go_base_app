package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(100);not null"`
	Email    string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string  `gorm:"type:varchar(100);not null"`
	Role     *string `gorm:"type:varchar(50);default:'ROLE_USER';not null"`
	Provider *string `gorm:"type:varchar(50);default:'local';not null"`
	Photo    *string `gorm:"not null;default:default.png"`
	Verified *bool   `gorm:"not null;default:false"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Photo    string `json:"photo,omitempty"`
	Provider string `json:"provider"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		Name:     user.Name,
		Email:    user.Email,
		Role:     *user.Role,
		Photo:    *user.Photo,
		Provider: *user.Provider,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Failed string `json:"failed"`
	Tag    string `json:"tag"`
	Value  string `json:"value"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Failed = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
