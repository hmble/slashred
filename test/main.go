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
	//commentsList := c.Comment.GetComments("golang", "ex18cx", "best")
	commentsList := c.Comment.GetComments("AnimalsBeingBros", "ex4gzu", "best")

	comments := c.Comment.List(commentsList, "best")

	fmt.Println(len(comments))

	for _, item := range comments {
		fmt.Println(item.Author)
	}

}
