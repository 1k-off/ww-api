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
	"ww-api-gateway/pkg/auth"
	"ww-api-gateway/pkg/metrics"
	"ww-api-gateway/pkg/target"
	"ww-api-gateway/pkg/user"
)

type Config struct {
	DatabaseConnectionString string
	Port                     string
	Cancel                   context.CancelFunc
	UserService              user.Service
	AuthService              auth.Service
	AccessTokenPublicKey     string
	TargetService            target.Service
	MetricsService           metrics.Service
}

func Start(c *Config) error {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Wazzup, man!")
	})

	api := app.Group("/api")
	router.UserRouter(api, c.UserService, c.AccessTokenPublicKey)
	router.AuthRouter(api, c.AuthService, c.AccessTokenPublicKey)
	router.TargetRouter(api, c.TargetService, c.AccessTokenPublicKey)
	router.CheckerRouter(api, c.TargetService, c.AccessTokenPublicKey)
	router.MetricsRouter(api, c.MetricsService, c.AccessTokenPublicKey)
	defer c.Cancel()
	return app.Listen(":" + c.Port)
}

func NewConfig(
	databaseConnectionString string,
	port string,
	accessTokenPrivateKey, accessTokenPublicKey, refreshTokenPrivateKey, refreshTokenPublicKey string,
	accessTokenExpiresIn, refreshTokenExpiresIn int,
) (*Config, error) {
	db, cancel, err := databaseConnection(databaseConnectionString)
	if err != nil {
		cancel()
		return nil, err
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
		cancel()
		return nil, err
	}
	userRepository := user.NewRepository(userCollection)
	userService := user.NewService(userRepository)

	authService := auth.NewService(userService, accessTokenPrivateKey, accessTokenPublicKey, refreshTokenPrivateKey, refreshTokenPublicKey, accessTokenExpiresIn, refreshTokenExpiresIn)

	targetCollection := db.Collection("targets")
	_, err = targetCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "url", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		cancel()
		return nil, err
	}
	targetRepository := target.NewRepository(targetCollection)
	targetService := target.NewService(targetRepository)

	metricsUptimeCollection := db.Collection("metrics_uptime")
	metricsSslCollection := db.Collection("metrics_ssl")
	metricsDomainExpirationCollection := db.Collection("metrics_domain_expiration")
	metricsRepository := metrics.NewRepository(metricsUptimeCollection, metricsSslCollection, metricsDomainExpirationCollection)
	metricsService := metrics.NewService(metricsRepository)

	return &Config{
		DatabaseConnectionString: databaseConnectionString,
		Port:                     port,
		UserService:              userService,
		Cancel:                   cancel,
		AuthService:              authService,
		AccessTokenPublicKey:     accessTokenPublicKey,
		TargetService:            targetService,
		MetricsService:           metricsService,
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

	opts := options.CreateCollection().
		SetTimeSeriesOptions(options.TimeSeries().
			SetGranularity("seconds").
			SetMetaField("metadata").
			SetTimeField("timestamp")).SetExpireAfterSeconds(60 * 60 * 24 * 30)
	err = client.Database("ww").CreateCollection(context.Background(), "metrics_uptime", opts)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	opts = options.CreateCollection().
		SetTimeSeriesOptions(options.TimeSeries().
			SetGranularity("hours").
			SetMetaField("metadata").
			SetTimeField("timestamp")).SetExpireAfterSeconds(60 * 60 * 24 * 30)
	err = client.Database("ww").CreateCollection(context.Background(), "metrics_ssl", opts)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	opts = options.CreateCollection().
		SetTimeSeriesOptions(options.TimeSeries().
			SetGranularity("hours").
			SetMetaField("metadata").
			SetTimeField("timestamp")).SetExpireAfterSeconds(60 * 60 * 24 * 30)
	err = client.Database("ww").CreateCollection(context.Background(), "metrics_domain_expiration", opts)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	return client.Database("ww"), cancel, nil
}
