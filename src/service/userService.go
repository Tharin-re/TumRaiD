package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tharin-re/TumRaiD/src/auth"
	"github.com/Tharin-re/TumRaiD/src/dto"
	"github.com/Tharin-re/TumRaiD/src/queries"
	"github.com/Tharin-re/TumRaiD/src/util"
	"github.com/gin-gonic/gin"
)

func RegisterUserPassEndpoint(c *gin.Context) {
	// Bind the JSON payload to the request body struct
	var registerUserPassReqBody dto.RegisterUserPassReqBody
	if err := c.ShouldBind(&registerUserPassReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the length of the username
	if !util.ValidUserLength(registerUserPassReqBody.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid length of User!"})
		return
	}

	// Validate the length of the password
	if !util.ValidPasswordLength(registerUserPassReqBody.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid length of Password!"})
		return
	}

	// Check for unacceptable characters in the username or password
	if util.ContainUnacceptableChar(registerUserPassReqBody.Username) || util.ContainUnacceptableChar(registerUserPassReqBody.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Illegal char"})
		return
	}

	// Check for duplicate username
	dup, err := queries.CheckDupUser(registerUserPassReqBody.Username, context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if dup {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Found duplicate username"})
		return
	}

	// Register the user
	err = queries.RegisterUserPass(registerUserPassReqBody.Username, registerUserPassReqBody.Password, context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Register user %s successfully", registerUserPassReqBody.Username)})
}

// LoginUserPassOrJWTEndPoint handles login requests using either username/password or JWT token.
// It validates the input, authenticates the user, and returns a new or refreshed JWT token.
func LoginUserPassOrJWTEndPoint(c *gin.Context) {
    // Define a struct to hold the login data from the request.
    var loginData dto.LoginData

    // Bind the JSON request body to the loginData struct.
    if err := c.ShouldBind(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Login Data"})
        return
    }

    // Check if a JWT token is provided.
    if loginData.Token != "" {
        // Authenticate the provided JWT token.
        err := auth.JWTAuthenticate(loginData.Token)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token, proceed to user/pass login"})
            return
        } else {
            // Refresh the JWT token if authentication is successful.
            tokenString, err := auth.JWTRefreshToken(loginData.Token)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "error during refreshing token"})
                return
            }
            // Return the refreshed token in the response.
            c.JSON(http.StatusOK, gin.H{"authToken": tokenString})
        }
    } else {
        // Check if username and password are provided.
        if loginData.Username == "" || loginData.Password == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "username and password cannot be empty"})
            return
        } else {
            // Check for unacceptable characters in username and password.
            if util.ContainUnacceptableChar(loginData.Username) || util.ContainUnacceptableChar(loginData.Password) {
                c.JSON(http.StatusBadRequest, gin.H{"error": "username or password contain unacceptable characters"})
                return
            } else {
                // Authenticate the user with username and password.
                if queries.LoginUserPass(loginData.Username, loginData.Password, context.Background()) != nil {
                    c.JSON(http.StatusBadRequest, gin.H{"error": "username or password incorrect"})
                    return
                }
            }
        }
        // Generate a new JWT token for the authenticated user.
        tokenString, err := auth.GenerateJWTClaim(loginData.Username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed creating JWT claim tokenString"})
            return
        }
        // Return the new token in the response.
        c.JSON(http.StatusOK, gin.H{"authToken": tokenString})
    }
}

