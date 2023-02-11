package target

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ww-api/pkg/entities"
)

type Repository interface {
	Get(id string) (*entities.Target, error)
	GetByUrl(name string) (*entities.Target, error)
	Create(t *entities.Target) (*entities.Target, error)
	Delete(id string) error
	Update(t *entities.Target) (*entities.Target, error)
	GetAll() ([]*entities.Target, error)
	Count() (int64, error)
	GetTargetsForChecker(checker string) ([]*entities.Target, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepository(c *mongo.Collection) Repository {
	return &repository{
		Collection: c,
	}
}

func (r *repository) Get(id string) (*entities.Target, error) {
	var target *entities.Target
	tid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err = r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyId: tid}).Decode(&target); err != nil {
		return nil, err
	}
	return target, nil
}

func (r *repository) GetByUrl(name string) (*entities.Target, error) {
	var target *entities.Target
	if err := r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyUrl: name}).Decode(&target); err != nil {
		return nil, err
	}
	return target, nil
}

func (r *repository) Create(t *entities.Target) (*entities.Target, error) {
	t.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repository) Delete(id string) error {
	tid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{entities.MongoKeyId: tid})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(t *entities.Target) (*entities.Target, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{entities.MongoKeyId: t.ID}, bson.M{"$set": t})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repository) GetAll() ([]*entities.Target, error) {
	result, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var targets []*entities.Target
	if err = result.All(context.Background(), &targets); err != nil {
		return nil, err
	}
	return targets, nil
}

func (r *repository) Count() (int64, error) {
	return r.Collection.CountDocuments(context.Background(), bson.D{})
}

// GetTargetsForChecker returns all targets that are enabled globally and for the given checker
func (r *repository) GetTargetsForChecker(checker string) ([]*entities.Target, error) {
	result, err := r.Collection.Find(context.Background(), bson.M{entities.MongoKeyIsActive: true, checker: true})
	if err != nil {
		return nil, err
	}
	var targets []*entities.Target
	if err = result.All(context.Background(), &targets); err != nil {
		return nil, err
	}
	return targets, nil
}

func (r *repository) DeleteAll() error {
	_, err := r.Collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
