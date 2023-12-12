package database

import (
	"github.com/gofiber/storage/mongodb"
)

type Database interface {
	Connect()
	GetDB() *mongodb.Storage
	IsLogVerbose() bool
}