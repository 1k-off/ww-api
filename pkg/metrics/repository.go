package metrics

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ww-api/pkg/entities"
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
}

func NewRepository(u *mongo.Collection, s *mongo.Collection, de *mongo.Collection) Repository {
	return &repository{
		uptimeCollection:           u,
		sslCollection:              s,
		domainExpirationCollection: de,
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
	_, _ = r.uptimeCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetadataUrl: url})
	_, _ = r.sslCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetadataUrl: url})
	_, _ = r.domainExpirationCollection.DeleteMany(context.Background(), bson.M{entities.MongoKeyMetadataUrl: url})
}

func (r *repository) GetDownTargets() ([]*entities.TargetDown, error) {
	return nil, errors.New("not implemented")
}

func (r *repository) GetSslExpiringSoon() ([]*entities.SslExpiringSoon, error) {
	return nil, errors.New("not implemented")
}

func (r *repository) GetDomainExpiringSoon() ([]*entities.DomainExpiringSoon, error) {
	return nil, errors.New("not implemented")
}

func (r *repository) GetStats() (*entities.MetricsStats, error) {
	return nil, errors.New("not implemented")
}
