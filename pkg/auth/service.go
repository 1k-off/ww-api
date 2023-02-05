package auth

import (
	"fmt"
	"time"
	"ww-api-gateway/pkg/entities"
	"ww-api-gateway/pkg/user"
	"ww-api-gateway/pkg/util"
)

type Service interface {
	Login(u *entities.User) (*entities.User, string, string, error)
	Refresh(refreshToken string) (*entities.User, string, string, error)
	Logout(accessToken string) error
	AccessTokenExpiresIn() time.Time
	RefreshTokenExpiresIn() time.Time
}

type service struct {
	u                      user.Service
	accessTokenPrivateKey  string
	accessTokenPublicKey   string
	accessTokenExpiresIn   int
	refreshTokenPrivateKey string
	refreshTokenPublicKey  string
	refreshTokenExpiresIn  int
}

func NewService(u user.Service, accessTokenPrivateKey, accessTokenPublicKey, refreshTokenPrivateKey, refreshTokenPublicKey string, accessTokenExpiresIn, refreshTokenExpiresIn int) Service {
	return &service{
		u:                      u,
		accessTokenPrivateKey:  accessTokenPrivateKey,
		accessTokenPublicKey:   accessTokenPublicKey,
		accessTokenExpiresIn:   accessTokenExpiresIn,
		refreshTokenPrivateKey: refreshTokenPrivateKey,
		refreshTokenPublicKey:  refreshTokenPublicKey,
		refreshTokenExpiresIn:  refreshTokenExpiresIn,
	}
}

func (s *service) Login(u *entities.User) (*entities.User, string, string, error) {
	usr, err := s.u.GetByLogin(u.Login)
	if err != nil {
		return nil, "", "", err
	}
	err = util.VerifyPassword(usr.PasswordHash, u.Password)
	if err != nil {
		return nil, "", "", err
	}
	accessTokenExpiresIn := time.Duration(s.accessTokenExpiresIn) * time.Second
	refreshTokenExpiresIn := time.Duration(s.refreshTokenExpiresIn) * time.Second
	accessToken, err := util.CreateJwtToken(accessTokenExpiresIn, usr.ID, s.accessTokenPrivateKey)
	if err != nil {
		return nil, "", "", err
	}
	refreshToken, err := util.CreateJwtToken(refreshTokenExpiresIn, usr.ID, s.refreshTokenPrivateKey)
	if err != nil {
		return nil, "", "", err
	}
	return usr, accessToken, refreshToken, nil
}

func (s *service) Refresh(refreshToken string) (*entities.User, string, string, error) {
	sub, err := util.ValidateJwtToken(refreshToken, s.refreshTokenPublicKey)
	if err != nil {
		return nil, "", "", err
	}
	usr, err := s.u.Get(fmt.Sprint(sub))
	if err != nil {
		return nil, "", "", err
	}
	accessTokenExpiresIn := time.Duration(s.accessTokenExpiresIn) * time.Second
	refreshTokenExpiresIn := time.Duration(s.refreshTokenExpiresIn) * time.Second
	accessToken, err := util.CreateJwtToken(accessTokenExpiresIn, usr.ID, s.accessTokenPrivateKey)
	if err != nil {
		return nil, "", "", err
	}
	refreshToken, err = util.CreateJwtToken(refreshTokenExpiresIn, usr.ID, s.refreshTokenPrivateKey)
	if err != nil {
		return nil, "", "", err
	}
	return usr, accessToken, refreshToken, nil
}

func (s *service) Logout(accessToken string) error {
	sub, err := util.ValidateJwtToken(accessToken, s.accessTokenPublicKey)
	if err != nil {
		return err
	}
	_, err = s.u.Get(fmt.Sprint(sub))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AccessTokenExpiresIn() time.Time {
	return time.Now().Add(time.Duration(s.accessTokenExpiresIn) * time.Second)
}

func (s *service) RefreshTokenExpiresIn() time.Time {
	return time.Now().Add(time.Duration(s.accessTokenExpiresIn) * time.Second)
}
