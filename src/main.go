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

// init function is called before the main function.
// It loads the configuration settings and initializes the database connection.
func init() {
    config.LoadConfig() // Load the configuration settings from the config file.
    queries.InitDB()    // Initialize the database connection.
}

// main function is the entry point of the application.
// It sets up the HTTP server and routes using the Gin framework.
func main() {
    engine := gin.Default() // Create a new Gin engine instance.

    // Define a group of routes for user-related endpoints.
    userRoute := engine.Group("/user")
    {
        // Define a POST route for user registration.
        // The request body should contain the following fields:
        // {
        //     Username string
        //     Password string
        // }
        userRoute.POST("/register", service.RegisterUserPassEndpoint)

        // Define a POST route for user login.
        // The request body should contain the following fields:
        // {
        //     Username string
        //     Password string
        //     Token string
        // }
        userRoute.POST("/login", service.LoginUserPassOrJWTEndPoint)
    }

    // Run the HTTP server on port 8080.
    engine.Run(":8080")
}
