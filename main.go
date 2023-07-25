package main

import (
	"context"
	"flag"
	"github.com/cawakharkov/hotel-reservation/api"
	"github.com/cawakharkov/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of API server")
	flag.Parse()

	app := fiber.New(config)

	apiv1 := app.Group("api/v1")

	UserHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/user", UserHandler.HandleGetUsers)
	apiv1.Get("/user/:id", UserHandler.HandleGetUser)
	apiv1.Post("/user", UserHandler.HandlePostUser)
	apiv1.Delete("/user/:id", UserHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", UserHandler.HandlePutUser)

	app.Listen(*listenAddr)
}
