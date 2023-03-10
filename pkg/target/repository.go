package target

import (
	"context"
	"github.com/rs/zerolog/log"
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
	GetTargetsForSslChecker() ([]*entities.SslTarget, error)
	GetTargetsForUptimeChecker() ([]*entities.UptimeTarget, error)
	GetTargetsForDomainExpirationChecker() ([]*entities.DomainExpirationTarget, error)
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
		log.Debug().Err(err).Msgf("error while parsing target id %s", id)
		return nil, err
	}
	if err = r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyId: tid}).Decode(&target); err != nil {
		log.Debug().Err(err).Msgf("error while getting target with id %s", id)
		return nil, err
	}
	return target, nil
}

func (r *repository) GetByUrl(url string) (*entities.Target, error) {
	var target *entities.Target
	if err := r.Collection.FindOne(context.Background(), bson.M{entities.MongoKeyUrl: url}).Decode(&target); err != nil {
		log.Debug().Err(err).Msgf("error while getting target with url %s", url)
		return nil, err
	}
	return target, nil
}

func (r *repository) Create(t *entities.Target) (*entities.Target, error) {
	t.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), t)
	if err != nil {
		log.Debug().Err(err).Msgf("error while creating target %s", t.URL)
		return nil, err
	}
	log.Debug().Msgf("created target %s", t.URL)
	return t, nil
}

func (r *repository) Delete(id string) error {
	tid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Debug().Err(err).Msgf("error while parsing target id %s", id)
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{entities.MongoKeyId: tid})
	if err != nil {
		log.Debug().Err(err).Msgf("error while deleting target with id %s", id)
		return err
	}
	log.Debug().Msgf("deleted target with id %s", id)
	return nil
}

func (r *repository) Update(t *entities.Target) (*entities.Target, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{entities.MongoKeyId: t.ID}, bson.M{"$set": t})
	if err != nil {
		log.Debug().Err(err).Msgf("error while updating target %s", t.URL)
		return nil, err
	}
	log.Debug().Msgf("updated target %s", t.URL)
	return t, nil
}

func (r *repository) GetAll() ([]*entities.Target, error) {
	result, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Debug().Err(err).Msg("error while getting all targets")
		return nil, err
	}
	var targets []*entities.Target
	if err = result.All(context.Background(), &targets); err != nil {
		log.Debug().Err(err).Msg("error while mapping data for all targets")
		return nil, err
	}
	return targets, nil
}

func (r *repository) Count() (int64, error) {
	return r.Collection.CountDocuments(context.Background(), bson.D{})
}

// GetTargetsForChecker returns all targets that are enabled globally and for the given checker with the whole target information
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

// GetTargetsForSslChecker returns all targets that are enabled globally and for the ssl checker
func (r *repository) GetTargetsForSslChecker() ([]*entities.SslTarget, error) {
	result, err := r.Collection.Find(context.Background(), bson.M{entities.MongoKeyIsActive: true, entities.MongoKeySsl: true})
	if err != nil {
		log.Debug().Err(err).Msg("error while getting all targets for ssl checker")
		return nil, err
	}
	var targets []*entities.SslTarget
	if err = result.All(context.Background(), &targets); err != nil {
		log.Debug().Err(err).Msg("error while mapping all targets data for ssl checker")
		return nil, err
	}
	return targets, nil
}

// GetTargetsForUptimeChecker returns all targets that are enabled globally and for the ssl checker
func (r *repository) GetTargetsForUptimeChecker() ([]*entities.UptimeTarget, error) {
	result, err := r.Collection.Find(context.Background(), bson.M{entities.MongoKeyIsActive: true, entities.MongoKeySsl: true})
	if err != nil {
		log.Debug().Err(err).Msg("error while getting all targets for uptime checker")
		return nil, err
	}
	var targets []*entities.Target
	if err = result.All(context.Background(), &targets); err != nil {
		log.Debug().Err(err).Msg("error while mapping all targets data for uptime checker")
		return nil, err
	}
	var resultTargets []*entities.UptimeTarget
	for _, target := range targets {
		resultTargets = append(resultTargets, &entities.UptimeTarget{
			URL:    target.URL,
			Config: target.Config.Uptime,
		})
	}
	return resultTargets, nil
}

// GetTargetsForDomainExpirationChecker returns all targets that are enabled globally and for the ssl checker
func (r *repository) GetTargetsForDomainExpirationChecker() ([]*entities.DomainExpirationTarget, error) {
	result, err := r.Collection.Find(context.Background(), bson.M{entities.MongoKeyIsActive: true, entities.MongoKeySsl: true})
	if err != nil {
		log.Debug().Err(err).Msg("error while getting all targets for domain expiration checker")
		return nil, err
	}
	var targets []*entities.DomainExpirationTarget
	if err = result.All(context.Background(), &targets); err != nil {
		log.Debug().Err(err).Msg("error while mapping all targets data for domain expiration checker")
		return nil, err
	}
	return targets, nil
}

func (r *repository) DeleteAll() error {
	_, err := r.Collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Debug().Err(err).Msg("error while deleting all targets")
		return err
	}
	return nil
}
