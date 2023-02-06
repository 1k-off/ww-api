package target

import "ww-api-gateway/pkg/entities"

type Service interface {
	Get(id string) (*entities.Target, error)
	GetByUrl(name string) (*entities.Target, error)
	Create(t *entities.Target) (*entities.Target, error)
	Delete(id string) error
	Update(t *entities.Target) (*entities.Target, error)
	GetAll() ([]*entities.Target, error)
	Count() (int64, error)
	GetTargetsForChecker(checker string) ([]*entities.Target, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(id string) (*entities.Target, error) {
	return s.repository.Get(id)
}

func (s *service) GetByUrl(name string) (*entities.Target, error) {
	return s.repository.GetByUrl(name)
}

func (s *service) Create(t *entities.Target) (*entities.Target, error) {
	return s.repository.Create(t)
}

func (s *service) Delete(id string) error {
	_, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}

func (s *service) Update(t *entities.Target) (*entities.Target, error) {
	return s.repository.Update(t)
}

func (s *service) GetAll() ([]*entities.Target, error) {
	return s.repository.GetAll()
}

func (s *service) Count() (int64, error) {
	return s.repository.Count()
}

func (s *service) GetTargetsForChecker(checker string) ([]*entities.Target, error) {
	return s.repository.GetTargetsForChecker(checker)
}