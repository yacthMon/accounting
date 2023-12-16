package models

import (
	"time"
)

type(
  Transaction struct {
    TransactionID  *string `bson:"_id,omitempty" json:"_id,omitempty"`
    AccountID     string `bson:"accountId" json:"accountId"`
    Total float32 `bson:"total" json:"total"`
    TransactionDate time.Time `bson:"transactionDate" json:"transactionDate"`
    TransactionLabel string `bson:"transactionLabel" json:"transactionLabel"`
    AccountType AccountType `bson:"accountType" json:"accountType"`
  }

  TransactionFilterDTO struct {
    FilterDay *int    `json:"filterDay,omitempty" bson:"filterDay,omitempty"`
    AccountID *string `json:"accountId,omitempty" bson:"accountId,omitempty"`
  }

  TransactionCreateDTO struct {
    AccountID     string `json:"accountId"`
		Total float32 `json:"total"`
		TransactionLabel string `json:"transactionLabel"`
		AccountType AccountType `json:"accountType"`
  }
)