package config

import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)

// Config struct holds the configuration settings for the application.
// It includes settings for the database, user password constraints, and JWT.
type Config struct {
    Database struct {
        Host     string // Database host
        Port     string // Database port
        User     string // Database user
        Password string // Database password
        Dbname   string // Database name
    }
    UserPassConstraints struct {
        UserLengthMin     int    // Minimum length for username
        UserLengthMax     int    // Maximum length for username
        PasswordLengthMin int    // Minimum length for password
        PasswordLengthMax int    // Maximum length for password
        IllegalChar       string // Illegal characters for username and password
    }
    JWT struct {
        JWTSecretKey   string // Secret key for signing JWT
        ExpirationTime int    // Expiration time for JWT in hours
    }
}

// Cfg is a global variable that holds the loaded configuration settings.
var Cfg Config

// LoadConfig loads the configuration settings from a YAML file.
// It uses the viper package to read the configuration file and unmarshal it into the Config struct.
func LoadConfig() {
    fmt.Println("Loading config...")

    // Add the current directory and the ./config directory to the config paths.
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config") // Ensure this matches your directory structure
    viper.SetConfigName("config")   // Set the config file name (without extension)
    viper.SetConfigType("yaml")     // Set the config file type

    // Read the config file
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Cannot read config file: %v", err)
    }

    // Unmarshal the config file into the Cfg struct
    if err := viper.Unmarshal(&Cfg); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }

    fmt.Println("Load config successful")
}
