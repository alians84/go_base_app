package main

import (
	// Import Go Fiber
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"pet-projectGoApi/app/middleware"
	"pet-projectGoApi/config/connectDb"
	"pet-projectGoApi/config/router"
	_ "pet-projectGoApi/docs"
)

func Init() {
	config, err := connectDb.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load environment variables \n", err.Error())
	}
	connectDb.ConnectDB(&config)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email youremail@provider.com
// @host localhost:3000
// @BasePath /
func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("Not found .env file, err: ", err)
	}

	app := fiber.New()
	Init()
	middleware.AppMiddleware(app)
	router.AppRouter(app)
	app.Listen(":3000")

	//config.AppConfig()

}
