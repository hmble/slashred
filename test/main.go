package main

import (
	"fmt"
	"log"
	"strings"

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

	//c.Listings()
	//names := strings.Split("t3_eiuti6", "_")
	names := strings.Split("t3_eioanh", "_")

	//	c.ListingByID(names)
	article := names[1]
	c.GetComments(article)
	fmt.Println("Success.........")

}
