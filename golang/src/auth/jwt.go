package auth

import (
    "time"
    "github.com/golang-jwt/jwt"
    "github.com/Tharin-re/TumRaiD/src/config"
)

// jwtKey is a global variable that holds the secret key used for signing JWTs.
var jwtKey = []byte(config.Cfg.JWT.JWTSecretKey)

// Claim struct defines the structure of the JWT claims.
// It includes the username and standard claims such as expiration time and issued at time.
type Claim struct {
    Username string
    jwt.StandardClaims
}

// GenerateJWTClaim generates a new JWT token for the given username.
// It sets the expiration time and issued at time for the token, signs it with the secret key, and returns the token string.
func GenerateJWTClaim(username string) (string, error) {
    // Calculate the expiration time for the token.
    expireTime := time.Now().Add(time.Duration(config.Cfg.JWT.ExpirationTime) * 24 * time.Hour)
    // Get the current time as the issued at time.
    issueTime := time.Now()
    // Create the claims for the token, including the username, expiration time, and issued at time.
    claims := &Claim{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            IssuedAt:  issueTime.Unix(),
        },
    }
    // Create a new token with the specified claims and signing method.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // Sign the token with the secret key and return the token string.
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}
