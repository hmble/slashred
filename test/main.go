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

	var c *slashred.Client = u.UserClient(token)

	path := "https://www.reddit.com/r/memes/comments/exkw6j/its_the_thought_that_counts/"

	//path := "https://www.reddit.com/r/dailyprogrammer/comments/dv0231/20191111_challenge_381_easy_yahtzee_upper_section"
	//	path := "https://www.reddit.com/r/golang/comments/7pnw2e/fun_golang_projects/"
	commentsList := c.Comment.GetComments(path, "best")

	for _, list := range commentsList {

		if list.Comment != nil {
			fmt.Printf("%s \n", list.Comment.Author)

			replies := c.Comment.Replies(0, list.Comment, "best", path)

			for _, reply := range replies {

				fmt.Printf("\t%s\n", reply.Author)
			}

		} else {
			fmt.Println("More count is ", list.More.Count)
		}
	}
}
