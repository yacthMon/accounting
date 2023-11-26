package models

type TransactionFilter struct {
	FilterDay *int    `json:"filterDay,omitempty" bson:"filterDay,omitempty"`
	AccountID *string `json:"accountId,omitempty" bson:"accountId,omitempty"`
}