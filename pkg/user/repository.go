package user

import (
	"context"
	"github.com/rs/zerolog/log"
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
		log.Debug().Err(err).Msg("error while creating user")
		return nil, err
	}
	log.Debug().Msg("user created")
	return u, nil
}

func (r *repository) Get(id string) (*entities.User, error) {
	var user *entities.User
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Debug().Err(err).Msgf("error while parsing user id %s", id)
		return nil, err
	}
	if err = r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyId: uid}).Decode(&user); err != nil {
		log.Debug().Err(err).Msgf("error while getting user with id %s", id)
		return nil, err
	}
	return user, nil
}

func (r *repository) GetByLogin(login string) (*entities.User, error) {
	var user *entities.User
	if err := r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyLogin: login}).Decode(&user); err != nil {
		log.Debug().Err(err).Msgf("error while getting user with login %s", login)
		return nil, err
	}
	return user, nil
}

func (r *repository) Update(u *entities.User) (*entities.User, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{entities.MongoKeyId: u.ID}, bson.M{"$set": u})
	if err != nil {
		log.Debug().Err(err).Msgf("error while updating user with id %s", u.ID.Hex())
		return nil, err
	}
	log.Debug().Msgf("user with id %s updated", u.ID.Hex())
	return u, nil
}

func (r *repository) Delete(id string) error {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Debug().Err(err).Msgf("error while parsing user id %s", id)
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{entities.MongoKeyId: uid})
	if err != nil {
		log.Debug().Err(err).Msgf("error while deleting user with id %s", id)
		return err
	}
	log.Debug().Msgf("user with id %s deleted", id)
	return nil
}
