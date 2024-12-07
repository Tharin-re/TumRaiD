package script

import (
	"context"
	"fmt"
	"time"
    "github.com/google/uuid"
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

	fmt.Printf("%s", currentDatabase)

	return currentDatabase, nil
}

func CheckDupUserPass(username string, password string, ctx context.Context) (bool, error) {
	fmt.Printf("Checking duplicate for %s\n", username)
	checkDupQuery := `SELECT username FROM tumraid.user_pass WHERE username = $1;`
	rows, err := Pool.Query(ctx, checkDupQuery, username)

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
	hash_ := util.MakePassHash(password)
	uuid_ := uuid.New()
	registerUserPassQuery := "INSERT INTO tumraid.user_pass (username,password_hashed,created_dt,user_id) values ($1,$2,current_timestamp,$3)"
	_, err := Pool.Exec(ctx, registerUserPassQuery, username, hash_,uuid_.String())
	if err != nil {
		return err
	}
	fmt.Printf("Register user %s successfully", username)
	return err
}

func CreateSessionUser(username string, ctx context.Context, session_length int) error {
	uuid_ := uuid.New().String()
	loginUserQuery := "INSERT INTO tumraid.current_session (username,auth_token,created_dt,expire_dt) values ($1,$2,current_timestamp,current_timestamp + interval '$3 minutes')"
	_,err := Pool.Exec(ctx, loginUserQuery,username,uuid_,session_length)
	if err!=nil {
		return err
	}
	fmt.Printf("Login user %s successfully",username)
	return nil
}

func LogOutAllPurpose(username string, ctx context.Context) error {
	logoutUserQuery := "DELETE FROM tumraid.current_session where username = $1"
	_,err := Pool.Exec(ctx,logoutUserQuery,username)
	if err!=nil {
		return err
	}
	fmt.Printf("Logout user %s successfully", username)
	return nil
}

func CheckIfLogonOrExpiredSession(username string, ctx context.Context) (bool,bool,error) {
	checkLogonQuery := "SELECT expire_dt from tumraid.current_session where username = $1"
	rows,err := Pool.Query(ctx,checkLogonQuery,username)
	if err!= nil {
		return false,false,err
	}
	defer rows.Close()
	if rows.Next(){
		var expire_dt time.Time
		if err := rows.Scan(&expire_dt); err!= nil {
			return false,false,err
		}
		if expire_dt.Before(time.Now()) {
			return true,true,nil
		} else {
			return true,false,nil
		}
	}
	return false,false,nil
}