package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/co-codin/hotel-reservation/api"
	"github.com/co-codin/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)

	apiv1 := app.Group("/api/v1")
	
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
