package api

import (
	"github.com/cawakharkov/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Aleks",
		LastName:  "ZZZ",
	}
	return c.JSON(u)
}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
