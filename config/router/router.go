package router

import (
	"github.com/gofiber/fiber/v2"
)

func AppRouter(a *fiber.App) {
	swagger := a.Group("/swagger")
	{
		swaggerRouter(swagger)
	}
	api := a.Group("/api")
	{
		/*api.Use(middleware.JWTProtected())*/

		apiRouter(api)
	}
}
