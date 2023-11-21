package handlers

import (
	"accounting/database"
	"accounting/models"
	"time"

	"github.com/gofiber/fiber/v2"
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
func TransactionCreate(c *fiber.Ctx) error {
	payload := struct {
		AccountID     string `json:"accountId"`
		TransactionID  string `json:"transactionId"`
		Total float32 `json:"total"`
		TransactionLabel string `json:"transactionLabel"`
		AccountType models.AccountType `json:"accountType"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"msg": err,
		})
	}
	newTransaction := &models.Transaction{
		AccountID: payload.AccountID,
		TransactionID: payload.TransactionID,
		Total: payload.Total,
		TransactionLabel: payload.TransactionLabel,
		AccountType: payload.AccountType,
		TransactionDate: time.Now(),
	}
	database.InsertTransaction(newTransaction)

	return c.JSON(fiber.Map{
		"success": true,
		"transaction":    newTransaction,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
