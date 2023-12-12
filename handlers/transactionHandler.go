package handlers

import (
	"accounting/models"
	"accounting/repositories"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct { 
	transactionRepository repositories.TransactionRepository
}

func CreateTransactionHandler(transactionRepository repositories.TransactionRepository) Handler {
	return &TransactionHandler{transactionRepository:transactionRepository}
}

func (h *TransactionHandler) Mount(router fiber.Router) {
	router.Get("/transactions", h.TransactionList)
	router.Post("/transaction", h.TransactionCreate)
	fmt.Println("Transaction mount complete")
}

// TransactionList returns a list of transaction
func (h *TransactionHandler) TransactionList(c *fiber.Ctx) error {
	payload := struct {
		Filter	*models.TransactionFilter `json:"filter,omitempty"`
	}{}
	if err := c.BodyParser(&payload); err != nil && len(c.Body()) > 0  {
		return c.JSON(fiber.Map{
			"success": false,
			"msg": err,
		})
	}
	transactions := h.transactionRepository.GetTransaction(payload.Filter)

	return c.JSON(fiber.Map{
		"success": true,
		"transactions":  transactions,
	})
}

// TransactionCreate create a transaction record
func (h *TransactionHandler) TransactionCreate(c *fiber.Ctx) error {
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
	h.transactionRepository.InsertTransaction(newTransaction)

	return c.JSON(fiber.Map{
		"success": true,
		"transaction":    newTransaction,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
