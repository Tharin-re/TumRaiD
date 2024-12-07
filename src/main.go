package main

import (
    "context"
    "fmt"
    "github.com/Tharin-re/TumRaiD/src/config"
    "github.com/Tharin-re/TumRaiD/src/script"
    "github.com/Tharin-re/TumRaiD/src/util"
)

func init() {
    config.LoadConfig()
    script.InitDB()
}

func main() {
    username := "testUser1_X"
    pass := "testTESTtest"
    if !util.ContainUnacceptableChar(username) && !util.ContainUnacceptableChar(pass) {
        err := script.RegisterUserPass(username, pass, context.Background())
        if err != nil {
            fmt.Println("Error:", err)
        }    
    } else {
        fmt.Println("contain unacceptable char")
    }
	// script.GetCurrentDatabase(context.Background())
}


