package service


import (
	"fmt"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Tharin-re/TumRaiD/src/queries"
	"github.com/Tharin-re/TumRaiD/src/util"
	"github.com/Tharin-re/TumRaiD/src/dto"
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

