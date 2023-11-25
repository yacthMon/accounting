package models

import (
	"time"
)

type Transaction struct {
  ID            string `bson:"_id"`
	AccountID     string `bson:"accountId"`
	TransactionID  string `bson:"transactionId"`
  Total float32 `bson:"total"`//10000
  TransactionDate time.Time `bson:"transactionDate"`
  TransactionLabel string `bson:"transactionLabel"`
  AccountType AccountType `bson:"accountType"`
}