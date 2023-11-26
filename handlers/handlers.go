package handlers

import (
	"accounting/database"
	"accounting/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TransactionList returns a list of transaction
func TransactionList(c *fiber.Ctx) error {
	payload := struct {
		Filter	*models.TransactionFilter `json:"filter,omitempty"`
	}{}
	if err := c.BodyParser(&payload); err != nil && len(c.Body()) > 0  {
		return c.JSON(fiber.Map{
			"success": false,
			"msg": err,
		})
	}
	transactions := database.GetTransaction(payload.Filter)

	return c.JSON(fiber.Map{
		"success": true,
		"transactions":  transactions,
	})
}

// TransactionCreate create a transaction record
func TransactionCreate(c *fiber.Ctx) error {
	payload := struct {
		AccountID     string `json:"accountId"`
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
