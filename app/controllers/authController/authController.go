package authController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"pet-projectGoApi/app/model"
	"pet-projectGoApi/config/connectDb"
	"strings"
	"time"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Registration GetCollection
// @Summary Get an item
// @Description Get an item by its ID
// @Accept  json
// @Produce  json
// @Tags Item
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /api/auth/register [get]
func (controller *AuthController) Registration(c *fiber.Ctx) error {
	var payload *model.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": errors})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Password do not match"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newUser := model.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Photo:    &payload.Photo,
	}

	result := connectDb.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterUserRecord(&newUser)}})
}

// Login GetCollection
// @Summary Get an item
// @Description Get an item by its ID
// @Accept  json
// @Produce  json
// @Tags Item
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /api/auth/register [get]
func (controller *AuthController) Login(c *fiber.Ctx) error {
	var payload *model.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": errors})
	}

	var user model.User

	result := connectDb.DB.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "invalid email or password"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "invalid email or  password "})
	}

	config, _ := connectDb.LoadConfig(".")
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMagAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

// LogoutUser GetItem
// @Summary Get an item
// @Description Get an item by its ID
// @Accept  json
// @Produce  json
// @Tags Item
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /api/auth/register [get]
func (controller *AuthController) LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
