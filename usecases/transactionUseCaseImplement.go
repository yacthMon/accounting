package usecases

import (
	"accounting/models"
	"accounting/repositories"
)

type TransactionUseCaseImplement struct {
	transactionRepository repositories.TransactionRepository
}

func CreateTransactionUseCase(transactionRepository repositories.TransactionRepository) TransactionUseCase {
	return &TransactionUseCaseImplement{
		transactionRepository: transactionRepository,
	}
}

func (useCase *TransactionUseCaseImplement) GetTransaction(filter *models.TransactionFilterDTO) []*models.Transaction {
	return useCase.transactionRepository.GetTransaction(filter)
}

func (useCase *TransactionUseCaseImplement) InsertTransaction(transaction *models.Transaction) {
	useCase.transactionRepository.InsertTransaction(transaction)
}