package metrics

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"sync"
	"ww-api/pkg/entities"
	"ww-api/pkg/target"
)

type Repository interface {
	InsertUptime(d interface{}) error
	InsertUptimeBatch(d []interface{}) error
	InsertSsl(d interface{}) error
	InsertSslBatch(d []interface{}) error
	InsertDomainExpiration(d interface{}) error
	InsertDomainExpirationBatch(d []interface{}) error
	Delete(url string)
	GetDownTargets() ([]*entities.TargetDown, error)
	GetSslExpiringSoon() ([]*entities.SslExpiringSoon, error)
	GetDomainExpiringSoon() ([]*entities.DomainExpiringSoon, error)
	GetStats() (*entities.MetricsStats, error)
}

type repository struct {
	uptimeCollection           *mongo.Collection
	sslCollection              *mongo.Collection
	domainExpirationCollection *mongo.Collection
	targetService              target.Service
}

func NewRepository(u *mongo.Collection, s *mongo.Collection, de *mongo.Collection, tSvc target.Service) Repository {
	return &repository{
		uptimeCollection:           u,
		sslCollection:              s,
		domainExpirationCollection: de,
		targetService:              tSvc,
	}
}

func (r *repository) InsertUptime(d interface{}) error {
	_, err := r.uptimeCollection.InsertOne(context.Background(), d)
	return err
}
func (r *repository) InsertUptimeBatch(d []interface{}) error {
	_, err := r.uptimeCollection.InsertMany(context.Background(), d)
	return err
}

func (r *repository) InsertSsl(d interface{}) error {
	_, err := r.sslCollection.InsertOne(context.Background(), d)
	return err
}
func (r *repository) InsertSslBatch(d []interface{}) error {
	_, err := r.sslCollection.InsertMany(context.Background(), d)
	return err
}

func (r *repository) InsertDomainExpiration(d interface{}) error {
	_, err := r.domainExpirationCollection.InsertOne(context.Background(), d)
	return err
}
func (r *repository) InsertDomainExpirationBatch(d []interface{}) error {
	_, err := r.domainExpirationCollection.InsertMany(context.Background(), d)
	return err
}

func (r *repository) Delete(url string) {
	_, _ = r.uptimeCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetricMetadataUrl: url})
	_, _ = r.sslCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetricMetadataUrl: url})
	_, _ = r.domainExpirationCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetricMetadataUrl: url})
}

func (r *repository) GetDownTargets() ([]*entities.TargetDown, error) {
	return nil, errors.New("not implemented")
}

// GetSslExpiringSoon returns a list of ssl targets that have expirationSoon set to true in the metrics collection
func (r *repository) GetSslExpiringSoon() ([]*entities.SslExpiringSoon, error) {
	targets, err := r.targetService.GetAll()
	if err != nil {
		return nil, err
	}
	expTargets, _ := r.getSslExpiringSoon(targets)
	return expTargets, nil
}

func (r *repository) GetDomainExpiringSoon() ([]*entities.DomainExpiringSoon, error) {
	return nil, errors.New("not implemented")
}

func (r *repository) GetStats() (*entities.MetricsStats, error) {
	return nil, errors.New("not implemented")
}

func (r *repository) getSslExpiringSoon(targets []*entities.Target) ([]*entities.SslExpiringSoon, error) {
	var (
		expTargets []*entities.SslExpiringSoon
		wg         sync.WaitGroup
	)
	for _, t := range targets {
		wg.Add(1)
		go func(t *entities.Target) {
			defer wg.Done()
			result := &entities.SslExpiringSoon{}
			result.Url = t.URL
			expiringSoon, date, err := r.targetIsSslExpirationSoonOrHasError(t)
			// !strings.HasPrefix(err.Error(), "no results for target") is a hack to ignore the error when there are no results for a target. TODO: remove this
			if err != nil && !strings.HasPrefix(err.Error(), "no results for target") {
				result.Error = err.Error()
				expTargets = append(expTargets, result)
			}
			if expiringSoon {
				result.Expires = date
				expTargets = append(expTargets, result)
			}
		}(t)
	}
	wg.Wait()
	return expTargets, nil
}

func (r *repository) targetIsSslExpirationSoonOrHasError(target *entities.Target) (bool, string, error) {
	pipeline := []bson.M{
		{"$match": bson.M{entities.MongoKeyMetricMetadataUrl: target.URL}},
		{"$sort": bson.M{entities.MongoKeyTimestamp: -1}},
		{"$limit": 1},
	}
	cur, err := r.sslCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return false, "", err
	}
	var results []*entities.SslData
	if err = cur.All(context.Background(), &results); err != nil {
		return false, "", err
	}
	if len(results) > 1 {
		return false, "", fmt.Errorf("more than one result for target %s", target.URL)
	}
	if len(results) == 0 {
		return false, "", fmt.Errorf("no results for target %s", target.URL)
	}
	if results[0].ExpiringSoon {
		return true, results[0].ExpirationDate, nil
	}
	if results[0].Error != "" {
		return true, "", fmt.Errorf(results[0].Error)
	}
	return false, "", nil
}
