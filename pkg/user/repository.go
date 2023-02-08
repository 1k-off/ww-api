package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ww-api/pkg/entities"
)

type Repository interface {
	Create(u *entities.User) (*entities.User, error)
	Get(id string) (*entities.User, error)
	GetByLogin(login string) (*entities.User, error)
	Update(u *entities.User) (*entities.User, error)
	Delete(id string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepository(c *mongo.Collection) Repository {
	return &repository{
		Collection: c,
	}
}

func (r *repository) Create(u *entities.User) (*entities.User, error) {
	u.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repository) Get(id string) (*entities.User, error) {
	var user *entities.User
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := r.Collection.FindOne(context.Background(), bson.M{"_id": uid}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetByLogin(login string) (*entities.User, error) {
	var user *entities.User
	if err := r.Collection.FindOne(context.Background(), bson.M{"login": login}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) Update(u *entities.User) (*entities.User, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": u})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repository) Delete(id string) error {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": uid})
	if err != nil {
		return err
	}
	return nil
}
