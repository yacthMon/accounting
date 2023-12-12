package repositories

import (
	"accounting/models"
)

type (
	TransactionRepository interface {
		InsertTransaction(transaction *models.Transaction)
		GetTransaction(transactionFilter *models.TransactionFilter) []*models.Transaction
	}
)