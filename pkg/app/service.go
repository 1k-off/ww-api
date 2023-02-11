package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"time"
	"ww-api/pkg/auth"
	"ww-api/pkg/entities"
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
		return nil, err
	}
	userCollection := db.Collection(entities.MongoCollectionNameUsers)
	_, err = userCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: entities.MongoKeyLogin, Value: 1}},
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

	targetCollection := db.Collection(entities.MongoCollectionNameTargets)
	_, err = targetCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: entities.MongoKeyUrl, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		cancel()
		return nil, err
	}
	targetRepository := target.NewRepository(targetCollection)
	targetService := target.NewService(targetRepository)

	metricsUptimeCollection := db.Collection(entities.MongoCollectionNameMetricsUptime)
	metricsSslCollection := db.Collection(entities.MongoCollectionNameMetricsSsl)
	metricsDomainExpirationCollection := db.Collection(entities.MongoCollectionNameMetricsDomainExpiration)
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
	cs, err := connstring.ParseAndValidate(connectionString)
	if err != nil {
		return nil, nil, err
	}
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
	collectionNames, _ := client.Database(cs.Database).ListCollectionNames(ctx, bson.M{})
	requiredTimeseriesCollections := []string{
		entities.MongoCollectionNameMetricsUptime,
		entities.MongoCollectionNameMetricsSsl,
		entities.MongoCollectionNameMetricsDomainExpiration,
	}
	for _, requiredCollection := range requiredTimeseriesCollections {
		if !util.Contains(collectionNames, requiredCollection) {
			switch requiredCollection {
			case entities.MongoCollectionNameMetricsUptime:
				opts := options.CreateCollection().
					SetTimeSeriesOptions(options.TimeSeries().
						SetGranularity(entities.MongoTsGranularitySeconds).
						SetMetaField(entities.MongoKeyMetadata).
						SetTimeField(entities.MongoKeyTimestamp)).SetExpireAfterSeconds(60 * 60 * 24 * 30)
				err = client.Database(cs.Database).CreateCollection(context.Background(), requiredCollection, opts)
				if err != nil {
					cancel()
					return nil, nil, err
				}
			case entities.MongoCollectionNameMetricsSsl, entities.MongoCollectionNameMetricsDomainExpiration:
				opts := options.CreateCollection().
					SetTimeSeriesOptions(options.TimeSeries().
						SetGranularity(entities.MongoTsGranularityHours).
						SetMetaField(entities.MongoKeyMetadata).
						SetTimeField(entities.MongoKeyTimestamp)).SetExpireAfterSeconds(60 * 60 * 24 * 30)
				err = client.Database(cs.Database).CreateCollection(context.Background(), requiredCollection, opts)
				if err != nil {
					cancel()
					return nil, nil, err
				}
			}
		}
	}

	return client.Database(cs.Database), cancel, nil
}

func (s *Service) Stop() {
	s.ctx.Done()
	s.DbCancelFunc()
}
