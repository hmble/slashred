package main

import (
	"log"
	"os"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {

	authenticator := &slashred.Authenticator{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       slashred.Scopes,
			Endpoint:     slashred.Endpoint,
			RedirectURL:  "https://example.com/auth",
		},
		Useragent: os.Getenv("USER_AGENT"),
	}
	u := &slashred.User{
		Name:          "Samyak",
		ProfileUrl:    "",
		IsPermanent:   true,
		Authenticator: authenticator,
	}

	token, er := slashred.TokenFromFile("token.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	u.UpdateToken(token)

	c := u.UserClient(token)

	// Use this to get list of flairs if you want to add flair to post
	// flairs := c.Flair.UserFlair("test")

	// for _, flair := range flairs {
	// 	fmt.Printf("Flair Name: %s \t Flaird ID: %s\n", flair.Type, flair.ID)
	// }
	c.Listing.LinkSubmit(slashred.PostData{
		"kind": "image",
		"nsfw": "false",
		"sr": "SuperModelIndia",
		"link": "https://i.imgur.com/lskdjf",

	})

}
