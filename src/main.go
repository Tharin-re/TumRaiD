package main

import (
	// "context"
	// "fmt"
	"github.com/Tharin-re/TumRaiD/src/config"
	"github.com/Tharin-re/TumRaiD/src/queries"
	// "github.com/Tharin-re/TumRaiD/src/util"
    "github.com/Tharin-re/TumRaiD/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	queries.InitDB()
}

func main() {
    engine := gin.Default()

    userRoute := engine.Group("/user")
    {
        userRoute.POST("/register", service.RegisterUserPassEndpoint)
        // body {
        //     Username string
	    //     Password string
        // }
        userRoute.POST("/login",service.LoginUserPassOrJWTEndPoint)
        // body {
        //     Username string
        //     Password string
        //     Token string
        // }
    }

    engine.Run(":8080")
}
