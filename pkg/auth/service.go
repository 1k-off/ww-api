package auth

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
	"ww-api/pkg/entities"
	"ww-api/pkg/user"
	"ww-api/pkg/util"
)

type Service interface {
	Login(u *entities.User) (*entities.User, string, string, error)
	Refresh(refreshToken string) (*entities.User, string, string, error)
	Logout(accessToken string) error
	AccessTokenExpiresIn() time.Time
	RefreshTokenExpiresIn() time.Time
	CheckAccessToken(accessToken string) (bool, error)
	GetCurrentUser(accessToken string) (*entities.User, error)
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
		log.Debug().Err(err).Msgf("user %s not found", u.Login)
		return nil, "", "", err
	}
	err = util.VerifyPassword(usr.PasswordHash, u.Password)
	if err != nil {
		log.Debug().Err(err).Msgf("password for user %s is incorrect", u.Login)
		return nil, "", "", err
	}
	accessTokenExpiresIn := time.Duration(s.accessTokenExpiresIn) * time.Second
	refreshTokenExpiresIn := time.Duration(s.refreshTokenExpiresIn) * time.Second
	accessToken, err := util.CreateJwtToken(accessTokenExpiresIn, usr.ID, s.accessTokenPrivateKey)
	if err != nil {
		log.Debug().Err(err).Msgf("failed to create access token for user %s", u.Login)
		return nil, "", "", err
	}
	refreshToken, err := util.CreateJwtToken(refreshTokenExpiresIn, usr.ID, s.refreshTokenPrivateKey)
	if err != nil {
		log.Debug().Err(err).Msgf("failed to create refresh token for user %s", u.Login)
		return nil, "", "", err
	}
	log.Debug().Msgf("user %s logged in", u.Login)
	return usr, accessToken, refreshToken, nil
}

func (s *service) Refresh(refreshToken string) (*entities.User, string, string, error) {
	sub, err := util.ValidateJwtToken(refreshToken, s.refreshTokenPublicKey)
	if err != nil {
		log.Debug().Err(err).Msg("refresh token is invalid")
		return nil, "", "", err
	}
	usr, err := s.u.Get(fmt.Sprint(sub))
	if err != nil {
		log.Debug().Err(err).Msgf("user %s not found", sub)
		return nil, "", "", err
	}
	accessTokenExpiresIn := time.Duration(s.accessTokenExpiresIn) * time.Second
	refreshTokenExpiresIn := time.Duration(s.refreshTokenExpiresIn) * time.Second
	accessToken, err := util.CreateJwtToken(accessTokenExpiresIn, usr.ID, s.accessTokenPrivateKey)
	if err != nil {
		log.Debug().Err(err).Msgf("failed to create access token for user %s", usr.Login)
		return nil, "", "", err
	}
	refreshToken, err = util.CreateJwtToken(refreshTokenExpiresIn, usr.ID, s.refreshTokenPrivateKey)
	if err != nil {
		log.Debug().Err(err).Msgf("failed to create refresh token for user %s", usr.Login)
		return nil, "", "", err
	}
	log.Debug().Msgf("user %s refreshed tokens", usr.Login)
	return usr, accessToken, refreshToken, nil
}

func (s *service) Logout(accessToken string) error {
	sub, err := util.ValidateJwtToken(accessToken, s.accessTokenPublicKey)
	if err != nil {
		log.Debug().Err(err).Msg("access token is invalid")
		return err
	}
	_, err = s.u.Get(fmt.Sprint(sub))
	if err != nil {
		log.Debug().Err(err).Msgf("user %s not found", sub)
		return err
	}
	log.Debug().Msgf("user %s logged out", sub)
	return nil
}

func (s *service) AccessTokenExpiresIn() time.Time {
	return time.Now().Add(time.Duration(s.accessTokenExpiresIn) * time.Second)
}

func (s *service) RefreshTokenExpiresIn() time.Time {
	return time.Now().Add(time.Duration(s.accessTokenExpiresIn) * time.Second)
}

// CheckAccessToken Checks validity of access token
func (s *service) CheckAccessToken(accessToken string) (bool, error) {
	_, err := util.ValidateJwtToken(accessToken, s.accessTokenPublicKey)
	if err != nil {
		log.Debug().Err(err).Msg("access token is invalid")
		return false, err
	}
	return true, nil
}

// CheckRefreshToken Checks validity of refresh token
func (s *service) CheckRefreshToken(refreshToken string) (bool, error) {
	_, err := util.ValidateJwtToken(refreshToken, s.refreshTokenPublicKey)
	if err != nil {
		log.Debug().Err(err).Msg("refresh token is invalid")
		return false, err
	}
	return true, nil
}

func (s *service) GetCurrentUser(accessToken string) (*entities.User, error) {
	sub, err := util.ValidateJwtToken(accessToken, s.accessTokenPublicKey)
	if err != nil {
		log.Debug().Err(err).Msg("access token is invalid")
		return nil, err
	}
	usr, err := s.u.Get(fmt.Sprint(sub))
	if err != nil {
		log.Debug().Err(err).Msgf("user %s not found", sub)
		return nil, err
	}
	return s.u.GetByLogin(usr.Login)
}
