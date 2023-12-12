package repositories

import (
	"accounting/database"
	"accounting/helper"
	"accounting/models"

	"context"
	"fmt"
	"sync"

	"github.com/gofiber/storage/mongodb"
)

type TransactionMongoRepository struct {
	db *mongodb.Storage
	isLogVerbose bool
}

var (
	mu sync.Mutex
)

func CreateTransactionMongoRepository(db *mongodb.Storage, isLogVerbose bool) TransactionRepository {
	return &TransactionMongoRepository {db: db, isLogVerbose: isLogVerbose}
}

func (r *TransactionMongoRepository) GetTransaction(transactionFilter *models.TransactionFilter) []*models.Transaction {
	filter := database.CreateFilterBSON(transactionFilter, r.isLogVerbose)
	cursor, err := r.db.Conn().Collection("transactions").Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	var findResult []*models.Transaction
	var results []*models.Transaction
	if err = cursor.All(context.TODO(), &findResult); err != nil {
		panic(err)
	}
	for _, result := range findResult {
		cursor.Decode(&result)
		results = append(results, result)
		if r.isLogVerbose {
			helper.PrintJSON(result)
		}
	}

	return results
}

func (r *TransactionMongoRepository) InsertTransaction(transaction *models.Transaction) {
	mu.Lock()
	_, err :=  r.db.Conn().Collection("transactions").InsertOne(context.TODO(), transaction)
	if err != nil {
		panic(err)
	}
	mu.Unlock()
}