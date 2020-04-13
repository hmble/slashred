package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	data := c.Users.Upvoted("astar0n", slashred.Option{
		"t":     "week",
		"type":  "links",
		"sort":  "new",
		"limit": "100",
	})

	for _, d := range data {
		if d.Data.Subreddit == "rust" {

			if !d.Data.IsSelf {
				fmt.Printf("* [%s](%s).\n", d.Data.Title, d.Data.URL)
				t := time.Unix(int64(d.Data.CreatedUtc), 0)
				fmt.Println(t.Format("02-01-2006"))

			}

		}
	}

}
