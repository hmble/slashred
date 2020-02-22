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

	//path :=https://www.reddit.com/r/redditdev/comments/avvl7u/nodejs_snoowrap_usage_find_number_of_comment/"
	path := "https://www.reddit.com/r/redditdev/comments/avvl7u/nodejs_snoowrap_usage_find_number_of_comment/"

	var c *slashred.Client = u.UserClient(token)

	// commentsList := c.Comment.GetComments(path, "best", true)
	// comments := c.Comment.List(commentsList, 8, "best", path, true)

	// count := 0

	// for _, comment := range comments {
	// 	if comment.Author != "[deleted]" {
	// 		count++
	// 	}
	// 	fmt.Printf("Author : %s ==> Parent [%s], Id [%s]\n", comment.Author, comment.Parent, comment.Id)
	// }

	// fmt.Println(len(comments))
	// fmt.Println("Count of non deleted comments ", count)

	sub := c.Link.GetSubmission(path)

	fmt.Println(sub.Selftext)
}
