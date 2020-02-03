package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hmble/slashred"
	"github.com/hmble/slashred/internal"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {

	authenticator := &internal.Authenticator{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       slashred.Scopes,
			Endpoint:     internal.Endpoint,
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

	var c *slashred.Client = u.UserClient(token)

	path := "https://www.reddit.com/r/memes/comments/exkw6j/its_the_thought_that_counts/"

	commentsList := c.Comment.GetComments(path, "best")
	comments := c.Comment.List(commentsList, 20, "best", path, true)

	for _, comment := range comments {

		fmt.Println(comment.Author)

	}

	fmt.Println("Got total comments ", len(comments))
}
