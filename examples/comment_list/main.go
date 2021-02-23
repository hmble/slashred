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
	// commentsList := c.Comment.GetComments("golang", "ex18cx", "best")

	path := "https://old.reddit.com/r/typescript/comments/aofcik/38_of_bugs_at_airbnb_could_have_been_prevented_by"
	comments := c.Comment.GetComments(path, "")
	
	fmt.Println(len(comments))

	for _, item := range comments {
		//fmt.Println(item.Author, "  ", item.Parent)
		fmt.Printf("%s\tId[%s]\tParent[%s]\n", item.Author, item.Id, item.Parent)
	}

}
