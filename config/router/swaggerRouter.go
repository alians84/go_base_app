package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func swaggerRouter(g fiber.Router) {
	//baseAuth := basicauth.New(basicauth.Config{
	//	Users: map[string]string{
	//		os.Getenv("SWAGGER_UI_USERNAME"): os.Getenv("SWAGGER_UI_PASSWORD"),
	//	},
	//})
	g.All("*", swagger.HandlerDefault)
}
