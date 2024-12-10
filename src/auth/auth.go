package auth

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt"
    "github.com/Tharin-re/TumRaiD/src/config"
)


type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func JWTAuthenticate(tokenString string) error {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return fmt.Errorf("error: signature invalid")
        }
        return fmt.Errorf("error while parsing token: %v", err)
    }

    if !token.Valid {
        return fmt.Errorf("error: token invalid")
    }

    if claims.ExpiresAt < time.Now().Unix() {
        return fmt.Errorf("error: token expired")
    }

    return nil
}

func JWTRefreshToken(tokenString string) (string, error) {
    if err := JWTAuthenticate(tokenString); err != nil {
        return "", fmt.Errorf("error: cannot authenticate JWT claim: %v", err)
    }

    claims := &Claims{}
    _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return "", fmt.Errorf("error while parsing token: %v", err)
    }

    expireTime := time.Now().Add(time.Duration(config.Cfg.JWT.ExpirationTime) * 24 * time.Hour)
    claims.ExpiresAt = expireTime.Unix()

    refreshedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    refreshedTokenString, err := refreshedToken.SignedString(jwtKey)
    if err != nil {
        return "", fmt.Errorf("error on refreshing token: %v", err)
    }

    return refreshedTokenString, nil
}
