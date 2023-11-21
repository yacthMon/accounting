package models

type AccountType string

const (
	BASE        AccountType = "Base"
	SECOND      AccountType = "Second"
	TRUE_WALLET AccountType = "True Wallet"
	WALLET      AccountType = "Wallet"
	CREDIT      AccountType = "Credit"
	RABBIT      AccountType = "Rabbit"
	EASY_PASS   AccountType = "Easy Pass"
)