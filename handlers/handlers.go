package handlers

import (
	"accounting/database"
	"accounting/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// TransactionList returns a list of transaction
func TransactionList(c *fiber.Ctx) error {
	transactions := database.GetTransaction()

	return c.JSON(fiber.Map{
		"success": true,
		"transactions":  transactions,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		Name: utils.CopyString(c.FormValue("user")),
	}
	database.Insert(user)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
