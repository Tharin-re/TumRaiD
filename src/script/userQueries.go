package script

import (
    "context"
    "fmt"
    // "log"
    // "github.com/jackc/pgx/v5/pgxpool"
	"github.com/Tharin-re/TumRaiD/src/util"
)

func GetCurrentDatabase(ctx context.Context) (string, error) {
    query := `SELECT current_database()`
    var currentDatabase string

    err := Pool.QueryRow(ctx, query).Scan(&currentDatabase)
    if err != nil {
        return "", fmt.Errorf("query failed: %w", err)
    }

    fmt.Printf("%s",currentDatabase)

    return currentDatabase, nil
}


func CheckDupUserPass(username string, password string, ctx context.Context) (bool, error) {
    idx := util.MakeUserPassHash(username, password)

    fmt.Printf("Checking duplicate for %s\n",idx)

    checkDupQuery := `SELECT idx FROM tumraid.user_pass WHERE idx = $1;`

    rows, err := Pool.Query(ctx, checkDupQuery, idx)
    if err != nil {
        return false, err
    }
    defer rows.Close()

    if !rows.Next() {
        return true, nil
    }

    return false, nil
}

func RegisterUserPass(username string, password string, ctx context.Context) error {
    hash_ := util.MakeUserPassHash(username,password)
    registerUserPassQuery := "INSERT INTO tumraid.user_pass values ($1,$2,current_timestamp,$3)"
    
    rows,err := Pool.Query(ctx,registerUserPassQuery,username,password,hash_)
    if err != nil {
        return err
    }
    defer rows.Close()
    fmt.Printf("Register user %s successfully",username)
    return err
}

// func LoginUser(username string, password string, ctx context.Context) error {
    
// }

