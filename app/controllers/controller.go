package controllers

import (
	"pet-projectGoApi/app/controllers/authController"
	"pet-projectGoApi/app/controllers/userController"
)

type Controller struct {
	UserController *userController.UserController
	AuthController *authController.AuthController
}
