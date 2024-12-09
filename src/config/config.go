package config

import (
	"fmt"
    "log"
    "github.com/spf13/viper"
)

type Config struct {
    Database struct {
        Host     string
        Port     string
        User     string
        Password string
        Dbname   string
    }
    UserPassConstraints struct {
        UserLengthMin int
        UserLengthMax int
        PasswordLengthMin int
        PasswordLengthMax int
        IllegalChar string
    }
}

var Cfg Config

func LoadConfig() {
	fmt.Println("Loading config...")

    viper.AddConfigPath(".")
    viper.AddConfigPath("./config") // Ensure this matches your directory structure
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Cannot read config file: %v", err)
    }

    if err := viper.Unmarshal(&Cfg); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }

	fmt.Println("Load config successful")
}
