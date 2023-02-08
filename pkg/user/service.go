package user

import (
	"ww-api/pkg/entities"
	"ww-api/pkg/util"
)

type Service interface {
	Create(u *entities.User) (*entities.User, error)
	Get(id string) (*entities.User, error)
	GetByLogin(login string) (*entities.User, error)
	Update(u *entities.User) (*entities.User, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(u *entities.User) (*entities.User, error) {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.PasswordHash = hashedPassword
	return s.repository.Create(u)
}

func (s *service) Get(id string) (*entities.User, error) {
	return s.repository.Get(id)
}

func (s *service) GetByLogin(login string) (*entities.User, error) {
	return s.repository.GetByLogin(login)
}

func (s *service) Update(u *entities.User) (*entities.User, error) {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.PasswordHash = hashedPassword
	return s.repository.Update(u)
}

func (s *service) Delete(id string) error {
	_, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}
