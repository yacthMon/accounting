package database

import (
	"accounting/helper"
	"accounting/models"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	mongoHost, isFound := os.LookupEnv("MONGO_HOST")
	if !isFound {
		mongoHost = "127.0.0.1"
	}
	mongoPort, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		mongoPort = 27017
	}
	mongoDatabase, isFound := os.LookupEnv("MONGO_DATABASE")
	if !isFound {
		mongoDatabase = "accounting"
	}
	store = mongodb.New(mongodb.Config{
    Host:       mongoHost,
    Port:       mongoPort,
    Database:   mongoDatabase,
    Reset:      false,
	})
	fmt.Println("Connected with Database")
}

func InsertTransaction(transaction *models.Transaction) {
	mu.Lock()
	_, err := store.Conn().Collection("transactions").InsertOne(context.TODO(), transaction)
	if err != nil {
		panic(err)
	}
	mu.Unlock()
}

func Get() {
}

func GetTransaction(transactionFilter *models.TransactionFilter) []*models.Transaction {
	filter := CreateFilterBSON(transactionFilter)
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

func CreateFilterBSON(transactionFilter *models.TransactionFilter) bson.M{
	filter := bson.M{}
	if(transactionFilter == nil ) {
		return filter
	}
	if transactionFilter.FilterDay != nil {
		now := time.Now()
		fmt.Println(now)
		startDate := now.Add(time.Duration(-*transactionFilter.FilterDay * 24)  * time.Hour)
		filter["transactionDate"] = bson.M{"$gte": primitive.NewDateTimeFromTime(startDate) }
	}
	if transactionFilter.AccountID != nil {
		filter["accountId"] = transactionFilter.AccountID
		fmt.Println("Append filter Account ID")
	}
	if isLogVerbose {
		fmt.Println(filter)
	}
	return filter
}