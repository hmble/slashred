package main

import (
	"fmt"
	"log"

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

	links := c.New("golang", slashred.Option{
		"limit": "10",
	})

	count := 0
	for _, link := range links {
		fmt.Printf("---------BODY--------\n\n")
		fmt.Println(link.Data.Title)
		fmt.Println("=====================")
		fmt.Printf("Author : %s\t Score : %d\n\n", link.Data.Author, link.Data.Score)
		count++
		fmt.Println(link.Data.Edited)
		fmt.Printf("-----------END--  %d  ---------\n", count)

	}

}

