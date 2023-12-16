package usecases

import "accounting/models"

type (
		TransactionUseCase interface {
			GetTransaction(filter *models.TransactionFilterDTO) []*models.Transaction
			InsertTransaction(transaction *models.Transaction)
		}
)
