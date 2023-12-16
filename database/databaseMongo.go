package database

import (
	"accounting/models"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDatabase struct {
	store *mongodb.Storage
	logVerbose	bool
}

func CreateMongoDB() Database {
	mongoDB := MongoDatabase {}
	mongoDB.Connect()
	return &mongoDB
}

// Connect with database
func (db *MongoDatabase) Connect() {
	logVerbose, err := strconv.ParseBool(os.Getenv("LOG_VERBOSE"))
	if err != nil {
		db.logVerbose = false
	} else {
		db.logVerbose = logVerbose
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
	db.store = mongodb.New(mongodb.Config{
    Host:       mongoHost,
    Port:       mongoPort,
    Database:   mongoDatabase,
    Reset:      false,
	})
	fmt.Println("Connect with Database ...")
	 _, err = db.store.Conn().ListCollectionNames(context.TODO(), bson.M{})
	 if err != nil {
		panic("Can't connect database")
	 }
	fmt.Println("Database connected!")
}

func (db *MongoDatabase) GetDB() *mongodb.Storage {
	return db.store
}

func (db *MongoDatabase) IsLogVerbose() bool{
	return db.logVerbose
}

func CreateFilterBSON(transactionFilter *models.TransactionFilterDTO, isLogVerbose bool) bson.M {
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