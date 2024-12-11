package auth

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt"
    "github.com/Tharin-re/TumRaiD/src/config"
)

// Claims struct defines the structure of the JWT claims.
// It includes the username and standard claims such as expiration time.
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// JWTAuthenticate verifies the provided JWT token string.
// It parses the token, validates its signature, checks its validity, and ensures it is not expired.
func JWTAuthenticate(tokenString string) error {
    claims := &Claims{}

    // Parse the token with the provided claims and key function.
    token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return fmt.Errorf("error: signature invalid")
        }
        return fmt.Errorf("error while parsing token: %v", err)
    }

    // Check if the token is valid.
    if !token.Valid {
        return fmt.Errorf("error: token invalid")
    }

    // Check if the token is expired.
    if claims.ExpiresAt < time.Now().Unix() {
        return fmt.Errorf("error: token expired")
    }

    return nil
}

// JWTRefreshToken refreshes the provided JWT token string.
// It authenticates the token, updates its expiration time, and returns a new token string.
func JWTRefreshToken(tokenString string) (string, error) {
    // Authenticate the token.
    if err := JWTAuthenticate(tokenString); err != nil {
        return "", fmt.Errorf("error: cannot authenticate JWT claim: %v", err)
    }

    claims := &Claims{}
    // Parse the token with the provided claims and key function.
    _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return "", fmt.Errorf("error while parsing token: %v", err)
    }

    // Update the expiration time.
    expireTime := time.Now().Add(time.Duration(config.Cfg.JWT.ExpirationTime) * 24 * time.Hour)
    claims.ExpiresAt = expireTime.Unix()

    // Create a new token with the updated claims.
    refreshedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    refreshedTokenString, err := refreshedToken.SignedString(jwtKey)
    if err != nil {
        return "", fmt.Errorf("error on refreshing token: %v", err)
    }

    return refreshedTokenString, nil
}
