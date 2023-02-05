package util

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateJwtToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	key, err := GetJwtPrivateKey(privateKey)
	if err != nil {
		return "", err
	}
	now := time.Now()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}
	return token, nil
}

func ValidateJwtToken(token string, publicKey string) (interface{}, error) {
	key, err := GetJwtPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims["sub"], nil
}

func GetJwtPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create: parse key: %w", err)
	}
	return key, nil
}

func GetJwtPublicKey(publicKey string) (*rsa.PublicKey, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}
	return key, nil
}
