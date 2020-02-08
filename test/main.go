package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hmble/slashred"
//	"github.com/hmble/slashred/internal"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {

	authenticator := &Authenticator{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       slashred.Scopes,
			Endpoint:     Endpoint,
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

	//	path := "https://www.reddit.com/r/memes/comments/exkw6j/its_the_thought_that_counts/"

	//path := "https://www.reddit.com/r/dailyprogrammer/comments/dv0231/20191111_challenge_381_easy_yahtzee_upper_section"
	path := "https://www.reddit.com/r/golang/comments/7pnw2e/fun_golang_projects/"
	commentsList := c.Comment.GetComments(path, "best")
	comments, usedLimit := c.Comment.List(commentsList, 20, "best", path, true)

	deleteCount := 0
	validCount := 0
	for _, comment := range comments {

		fmt.Println(comment.Author)

		if comment.Author == "[deleted]" {
			deleteCount++
		} else {

			validCount++
		}

	}

	fmt.Println("Got total comments ", len(comments))
	fmt.Println("Total delete count is ", deleteCount)
	fmt.Println("Total valid comment count is ", validCount)
	fmt.Println("Used limit is ", usedLimit)

}
