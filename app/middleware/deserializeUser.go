package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"pet-projectGoApi/app/model"
	"pet-projectGoApi/config/connectDb"
	"strings"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string

	authoriztion := c.Get("Authorization")

	if strings.HasPrefix(authoriztion, "Bearer ") {
		tokenString = strings.TrimPrefix(authoriztion, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := connectDb.LoadConfig(".")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})
	}

	var user model.User
	connectDb.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if claims["sub"] != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", model.FilterUserRecord(&user))
	return c.Next()
}
