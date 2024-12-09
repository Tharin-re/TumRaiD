package queries

import (
	"context"
	"fmt"
	"log"
	"github.com/Tharin-re/TumRaiD/src/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
    connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Dbname)
	var err error
	Pool, err = pgxpool.New(context.Background(),connString)
	if err != nil {
		log.Fatalf("Cannot establish DB connection with error :%s at Pool",err)
		log.Fatalln(Pool)
	}
	fmt.Printf("Initiated DB %s\n",config.Cfg.Database.Dbname)
}



