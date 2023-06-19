package userController

import (
	"github.com/gofiber/fiber/v2"
	"pet-projectGoApi/app/model"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// GetItem
// @Summary Get an item
// @Description Get an item by its ID
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} GetItemRequest
// @Failure 400 {object} model.HTTPError
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /api/users/{id} [get]
func (controller *UserController) GetItem(c *fiber.Ctx) error {
	var User GetItemRequest
	User.Id = 1
	User.Name = "name"
	return c.JSON(User)
}

func (controller *UserController) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(model.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
