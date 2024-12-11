package queries

import (
    "context"
    "fmt"
    "time"
    "github.com/google/uuid"
    "github.com/Tharin-re/TumRaiD/src/util"
)

// GetCurrentDatabase returns the name of the current database.
// It executes a SQL query to fetch the current database name and returns it.
func GetCurrentDatabase(ctx context.Context) (string, error) {
    query := `SELECT current_database()`
    var currentDatabase string
    err := Pool.QueryRow(ctx, query).Scan(&currentDatabase)
    if err != nil {
        return "", fmt.Errorf("query failed: %w", err)
    }

    fmt.Printf("%s", currentDatabase)

    return currentDatabase, nil
}

// CheckDupUser checks if a username already exists in the database.
// It executes a SQL query to check for the existence of the username in the user_pass table.
func CheckDupUser(username string, ctx context.Context) (bool, error) {
    fmt.Printf("Checking duplicate for %s\n", username)
    checkDupQuery := `SELECT username FROM tumraid.user_pass WHERE username = $1;`
    rows, err := Pool.Query(ctx, checkDupQuery, username)

    if err != nil {
        return false, err
    }
    defer rows.Close()

    // If no rows are returned, the username does not exist.
    if !rows.Next() {
        return false, nil
    }

    // If a row is returned, the username exists.
    return true, nil
}

// RegisterUserPass registers a new user with a hashed password.
// It generates a UUID for the user and inserts the username, hashed password, and creation timestamp into the user_pass table.
func RegisterUserPass(username string, password string, ctx context.Context) error {
    hash_ := util.MakePassHash(password)
    uuid_ := uuid.New()
    registerUserPassQuery := "INSERT INTO tumraid.user_pass (username,password_hashed,created_dt,user_id) values ($1,$2,current_timestamp,$3)"
    _, err := Pool.Exec(ctx, registerUserPassQuery, username, hash_, uuid_.String())
    if err != nil {
        return err
    }
    fmt.Printf("Register user %s successfully", username)
    return nil
}

// CreateSessionUser creates a new session for a user with a specified session length.
// It generates a UUID for the session and inserts the username, auth token, creation timestamp, and expiration timestamp into the current_session table.
func CreateSessionUser(username string, ctx context.Context, session_length int) error {
    uuid_ := uuid.New().String()
    createSessionUserQuery := "INSERT INTO tumraid.current_session (username,auth_token,created_dt,expire_dt) values ($1,$2,current_timestamp,current_timestamp + interval '$3 minutes')"
    _, err := Pool.Exec(ctx, createSessionUserQuery, username, uuid_, session_length)
    if err != nil {
        return err
    }
    fmt.Printf("Login user %s successfully", username)
    return nil
}

// LoginUserPass checks if the provided username and password are correct.
// It hashes the provided password and executes a SQL query to check for a matching username and hashed password in the user_pass table.
func LoginUserPass(username string, password string, ctx context.Context) error {
    hash_ := util.MakePassHash(password)
    loginUserPassQuery := "SELECT username, password_hashed from tumraid.user_pass where username = $1 and password_hashed = $2"
    rows, err := Pool.Query(ctx, loginUserPassQuery, username, hash_)
    if err != nil {
        return err
    }
    defer rows.Close()
    if !rows.Next() {
        return fmt.Errorf("username or password incorrect")
    }
    return nil
}

// LogOutAllPurpose logs out a user by deleting their session.
// It executes a SQL query to delete the user's session from the current_session table.
func LogOutAllPurpose(username string, ctx context.Context) error {
    logoutUserQuery := "DELETE FROM tumraid.current_session where username = $1"
    _, err := Pool.Exec(ctx, logoutUserQuery, username)
    if err != nil {
        return err
    }
    fmt.Printf("Logout user %s successfully", username)
    return nil
}

// CheckIfLogonOrExpiredSession checks if a user is logged on and if their session is expired.
// It executes a SQL query to fetch the expiration timestamp of the user's session and compares it with the current time.
func CheckIfLogonOrExpiredSession(username string, ctx context.Context) (bool, bool, error) {
    checkLogonQuery := "SELECT expire_dt from tumraid.current_session where username = $1"
    rows, err := Pool.Query(ctx, checkLogonQuery, username)
    if err != nil {
        return false, false, err
    }
    defer rows.Close()
    if rows.Next() {
        var expire_dt time.Time
        if err := rows.Scan(&expire_dt); err != nil {
            return false, false, err
        }
        // Check if the session is expired.
        if expire_dt.Before(time.Now()) {
            return true, true, nil
        } else {
            return true, false, nil
        }
    }
    // If no rows are returned, the user is not logged on.
    return false, false, nil
}
