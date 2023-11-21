package database

import (
	"accounting/models"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/gofiber/storage/mongodb"
)

var (
	mu sync.Mutex
	store *mongodb.Storage
)

// Connect with database
func Connect() {
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		mongoPort = 27017
	}
	mongoDatabase := os.Getenv("MONGO_DATABASE")
	store = mongodb.New(mongodb.Config{
    Host:       mongoHost,
    Port:       mongoPort,
    Database:   mongoDatabase,
    Reset:      false,
	})
	fmt.Println("Connected with Database")
}

func Insert(user *models.User) {
	mu.Lock()
	mu.Unlock()
}

func Get() {
}