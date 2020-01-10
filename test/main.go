package main

import (
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

	//c.Delete(slashred.CommentPrefix + "fdpzqgs")

	text := `
	Edited this posts data by api endpoint
	From user /u/anyrandomuserwhichIdontknow

	Edit :

	Text is edited please IGNORE

	**API ENDPOINT OF /api/editusertext
	`
	//commenttext := "Its a comment text"
	//c.PostComment(slashred.LinkPrefix+"emmcy8", commenttext)
	//c.EditUserText(slashred.CommentPrefix+"fdq0uck", text)
	c.EditUserText(slashred.LinkPrefix+"emn0zh", text)
	// emn0zh
	//c.Upvote(slashred.LinkPrefix + "emh2un")

}
