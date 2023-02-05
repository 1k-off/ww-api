package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"ww-api-gateway/api/router"
	"ww-api-gateway/pkg/user"
)

type Config struct {
	DatabaseConnectionString string
	Port                     string
	Cancel                   context.CancelFunc
	UserService              user.Service
}

func Start(c *Config) error {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Wazzup, man!")
	})

	api := app.Group("/api")
	router.UserRouter(api, c.UserService)
	defer c.Cancel()
	return app.Listen(":" + c.Port)
}

func NewConfig(databaseConnectionString string, port string) (*Config, error) {
	db, cancel, err := databaseConnection(databaseConnectionString)
	if err != nil {
		return nil, err
		cancel()
	}
	log.Println("Connected to database")

	userCollection := db.Collection("users")
	_, err = userCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "login", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return nil, err
		cancel()
	}
	userRepository := user.NewRepository(userCollection)
	userService := user.NewService(userRepository)

	return &Config{
		DatabaseConnectionString: databaseConnectionString,
		Port:                     port,
		UserService:              userService,
		Cancel:                   cancel,
	}, nil
}

func databaseConnection(connectionString string) (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString).SetServerSelectionTimeout(time.Second*10))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return client.Database("ww"), cancel, nil
}
