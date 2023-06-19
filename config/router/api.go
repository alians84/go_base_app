package router

import (
	"github.com/gofiber/fiber/v2"
	"pet-projectGoApi/app/controllers"
	"pet-projectGoApi/app/controllers/authController"
	"pet-projectGoApi/app/controllers/userController"
)

func apiRouter(f fiber.Router) {
	appController := initializationController()

	user := f.Group("/users")
	{
		//user.Post("/", appController.UserController.GetItem)
		user.Get("/:id", appController.UserController.GetItem)
	}

	auth := f.Group("/auth")
	{
		auth.Get("", appController.AuthController.Authorization)
	}

}

func initializationController() controllers.Controller {
	return controllers.Controller{
		UserController: userController.NewUserController(),
		AuthController: authController.NewAuthController(),
	}
}
