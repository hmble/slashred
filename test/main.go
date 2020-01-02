package main

import (
	"fmt"
	"log"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	u := &slashred.User{
		Name:        "Samyak",
		ProfileUrl:  "",
		IsPermanent: true,
	}
	token, er := slashred.TokenFromFile("token.json")
	//	token, er := u.Authenticate()
	if er != nil {
		log.Fatal("Error in reading token")
	}

	//	var a *internal.Authenticator = internal.RedditClient

	slashred.UpdateToken(token)

	// internal.SaveToken("token.json", token)

	// fmt.Println("Token saved")
	var c *slashred.Client = u.UserClient(token)

	c.PrefsBlocked()
	c.PrefsTrusted()
	c.PrefsMessaging()
	c.MeBlocked()
	c.MeFriends()

	fmt.Println("Success.........")

}
