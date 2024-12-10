package auth

import (
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/Tharin-re/TumRaiD/src/config"
)

var jwtKey = []byte(config.Cfg.JWT.JWTSecretKey)

type Claim struct {
	Username string
	jwt.StandardClaims
}

func GenerateJWTClaim(username string) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.Cfg.JWT.ExpirationTime)*24*time.Hour)
	issueTime := time.Now()
	claims := &Claim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: issueTime.Unix(),
		},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err!=nil {
		return "",err
	}
	return tokenString, nil
}

