package models

import (
	"time"
)

type Transaction struct {
  ID            string `bson:"_id" json:"_id"`
	AccountID     string `bson:"accountId" json:"accountId"`
	TransactionID  string `bson:"transactionId" json:"transactionId"`
  Total float32 `bson:"total" json:"total"`
  TransactionDate time.Time `bson:"transactionDate" json:"transactionDate"`
  TransactionLabel string `bson:"transactionLabel" json:"transactionLabel"`
  AccountType AccountType `bson:"accountType" json:"accountType"`
}