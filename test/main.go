package main

import (
	"fmt"
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

	path := "https://www.reddit.com/r/learnprogramming/comments/bs6466/why_study_programming_when_you_can_just_play_an/"
	comments := c.Comment.GetComments(path, "")

	for _, comment := range comments {
		fmt.Println("-----------------  ", comment.Author)
		if len(comment.Body) < 50 {
			fmt.Println(comment.Body)
		} else {
			fmt.Println(comment.Body[:50])
		}

		fmt.Println("-----------------")
	}

	fmt.Printf("Comments length is %d\n", len(comments))

}
