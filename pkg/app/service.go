package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"ww-api/pkg/auth"
	"ww-api/pkg/metrics"
	"ww-api/pkg/target"
	"ww-api/pkg/user"
	"ww-api/pkg/util"
)

type Service struct {
	DbCancelFunc         context.CancelFunc
	UserService          user.Service
	AuthService          auth.Service
	AccessTokenPublicKey string
	TargetService        target.Service
	MetricsService       metrics.Service
	ctx                  context.Context
}

func New(dbConnectionString, atPrivateKey, atPublicKey, rtPrivateKey, rtPublicKey string, atExpiresIn, rtExpiresIn int) (*Service, error) {
	db, cancel, err := databaseConnection(dbConnectionString)
	if err != nil {
		cancel()
		return nil, err
	}
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

	authService := auth.NewService(userService, atPrivateKey, atPublicKey, rtPrivateKey, rtPublicKey, atExpiresIn, rtExpiresIn)

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

	return &Service{
		DbCancelFunc:         cancel,
		UserService:          userService,
		AuthService:          authService,
		AccessTokenPublicKey: atPublicKey,
		TargetService:        targetService,
		MetricsService:       metricsService,
		ctx:                  context.Background(),
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

	collectionNames, _ := client.Database("ww").ListCollectionNames(ctx, bson.M{})
	requiredTimeseriesCollections := []string{"metrics_uptime", "metrics_ssl", "metrics_domain_expiration"}
	for _, requiredCollection := range requiredTimeseriesCollections {
		if !util.Contains(collectionNames, requiredCollection) {
			switch requiredCollection {
			case "metrics_uptime":
				opts := options.CreateCollection().
					SetTimeSeriesOptions(options.TimeSeries().
						SetGranularity("seconds").
						SetMetaField("metadata").
						SetTimeField("timestamp")).SetExpireAfterSeconds(60 * 60 * 24 * 30)
				err = client.Database("ww").CreateCollection(context.Background(), requiredCollection, opts)
				if err != nil {
					cancel()
					return nil, nil, err
				}
			case "metrics_ssl", "metrics_domain_expiration":
				opts := options.CreateCollection().
					SetTimeSeriesOptions(options.TimeSeries().
						SetGranularity("hours").
						SetMetaField("metadata").
						SetTimeField("timestamp")).SetExpireAfterSeconds(60 * 60 * 24 * 30)
				err = client.Database("ww").CreateCollection(context.Background(), requiredCollection, opts)
				if err != nil {
					cancel()
					return nil, nil, err
				}
			}
		}
	}

	return client.Database("ww"), cancel, nil
}

func (s *Service) Stop() {
	s.ctx.Done()
	s.DbCancelFunc()
}
