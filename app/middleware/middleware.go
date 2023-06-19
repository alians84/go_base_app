package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AppMiddleware(app *fiber.App) {
	fiberMiddleware(app)
}
