package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
)

const redirectURI = "http://localhost:8080/callback"

var (
	userAgent = os.Getenv("USER_AGENT")
	auth  = slashred.NewAuthenticatior(redirectURI, userAgent,slashred.Scopes...)
	ch    = make(chan *slashred.Client)
	state = "abc123"
)

func main() {
	// first start an HTTP server
		token, er := slashred.TokenFromFile("newtoken.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	client := auth.NewClient(token, userAgent)
	
	account, _, err := client.Account.GetMe()

	
	if err != nil {
		log.Println("from aboutme ", err)
	}

	//PANIC RECOVER------------------------------
    defer func() { //catch or finally
        if r := recover(); r != nil { //catch
            fmt.Println("Recover Triggered: ", r)
        }
    }()

	fmt.Println(account.Name)
}
