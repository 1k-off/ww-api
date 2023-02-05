package target

import "ww-api-gateway/pkg/entities"

type Service interface {
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

//func (s *service) GetDown() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (s *service) GetSslExp() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (s *service) GetDomainExp() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (s *service) GetStats() ([]*entities.Target, error) {
//	return nil, nil
//}
//
//func (s *service) DeleteAll() error {
//	return nil
//}
