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

	trophies, err := c.Trophies()

	if err != nil {
		log.Fatal("Error in trophies function")
	}

	for _, trophy := range trophies {
		fmt.Println(trophy.Name)
	}

	// pref, _ := c.GetMyPreferences()

	// fmt.Println(pref.VideoAutoplay)
	fmt.Println("Success.........")

}
