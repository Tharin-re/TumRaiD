package queries

import (
    "context"
    "fmt"
    "log"
    "github.com/Tharin-re/TumRaiD/src/config"
    "github.com/jackc/pgx/v5/pgxpool"
)

// Pool is a global variable that holds the connection pool for the database.
var Pool *pgxpool.Pool

// InitDB initializes the database connection pool.
// It constructs the connection string using the configuration values and establishes a connection to the database.
// If the connection fails, it logs a fatal error and terminates the program.
func InitDB() {
    // Construct the connection string using the configuration values.
    connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Dbname)
    
    // Establish a connection to the database using the connection string.
    var err error
    Pool, err = pgxpool.New(context.Background(), connString)
    if err != nil {
        // Log a fatal error and terminate the program if the connection fails.
        log.Fatalf("Cannot establish DB connection with error: %s at Pool", err)
        log.Fatalln(Pool)
    }
    
    // Print a message indicating that the database has been successfully initialized.
    fmt.Printf("Initiated DB %s\n", config.Cfg.Database.Dbname)
}
