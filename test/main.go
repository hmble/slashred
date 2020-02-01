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
	if er != nil {
		log.Fatal("Error in reading token")
	}

	slashred.UpdateToken(token)

	var c *slashred.Client = u.UserClient(token)
	//commentsList := c.Comment.GetComments("golang", "ex18cx", "best")
	commentsList := c.Comment.GetComments("AnimalsBeingBros", "ex4gzu", "best")

	comments := c.Comment.List(commentsList, "best")

	fmt.Println(len(comments))

	for _, item := range comments {
		fmt.Println(item.Author)
	}

}
