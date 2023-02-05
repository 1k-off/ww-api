package target

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ww-api-gateway/pkg/entities"
)

type Repository interface {
	Get(id string) (*entities.Target, error)
	Create(t *entities.Target) (*entities.Target, error)
	Delete(id string) error
	Update(t *entities.Target) (*entities.Target, error)
	GetAll() ([]*entities.Target, error)
	//GetDown() ([]*entities.Target, error)
	//GetSslExp() ([]*entities.Target, error)
	//GetDomainExp() ([]*entities.Target, error)
	//GetStats() ([]*entities.Target, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepository(c *mongo.Collection) Service {
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
	if err := r.Collection.FindOne(context.Background(), bson.M{"_id": tid}).Decode(&target); err != nil {
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
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": tid})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(t *entities.Target) (*entities.Target, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": t.ID}, bson.M{"$set": t})
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
	if err := result.All(context.Background(), &targets); err != nil {
		return nil, err
	}
	return targets, nil

}

//func (r *repository) GetDown() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (r *repository) GetSslExp() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (r *repository) GetDomainExp() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (r *repository) GetStats() ([]*entities.Target, error) {
//	return nil, nil
//}

func (r *repository) DeleteAll() error {
	_, err := r.Collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
