package main

import (
	"accounting/database"
	"accounting/handlers"
	"accounting/repositories"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

type FiberApp struct {
	app *fiber.App
	db database.Database
}

func NewFiberApp(db database.Database) FiberApp {
	return FiberApp{
		app: fiber.New(fiber.Config{
			Prefork: *prod, // go run app.go -prod
		}),
		db: db,
	}
}

func (a *FiberApp) Start() {
	// Middleware
	a.app.Use(recover.New())
	a.app.Use(logger.New())

	a.InitializeTransactionHandler()

	// Order Matter
	// Setup static files
	a.app.Static("/", "./static/public")

	// Handle not founds
	a.app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(a.app.Listen(*port)) // go run app.go -port=:3000
}

func (a *FiberApp) InitializeTransactionHandler() {
	v1 := a.app.Group("/api/v1")
	transactionRepository := repositories.CreateTransactionMongoRepository(a.db.GetDB(), a.db.IsLogVerbose())
	transactionHandler := handlers.CreateTransactionHandler(transactionRepository)
	transactionHandler.Mount(v1)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	// Parse command-line flags
	flag.Parse()

	db := database.CreateMongoDB()

	fiberApp := NewFiberApp(db)
	fiberApp.Start()

}
