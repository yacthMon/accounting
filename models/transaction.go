package models

import (
	"time"
)

type Transaction struct {
	AccountID     string `json:"accountId"`
	TransactionID  string `json:"transactionId"`
  Total float32 `json:"total"`//10000
  TransactionDate time.Time `json:"transactionDate"`
  TransactionLabel string `json:"transactionLabel"`
  AccountType AccountType `json:"accountType"`
}