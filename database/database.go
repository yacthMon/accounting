package database

import (
	"accounting/helper"
	"accounting/models"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/gofiber/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	mu sync.Mutex
	store *mongodb.Storage
	isLogVerbose bool
)

// Connect with database
func Connect() {
	logVerbose, err := strconv.ParseBool(os.Getenv("LOG_VERBOSE"))
	if err != nil {
		isLogVerbose = false
	} else {
		isLogVerbose = logVerbose
	}
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

func GetTransaction() []*models.Transaction {
	filter := bson.D{{}}
	cursor, err := store.Conn().Collection("transactions").Find(context.TODO(), filter)
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
		if isLogVerbose {
			helper.PrintJSON(result)
		}
	}

	return results
}