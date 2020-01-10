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

  [RANDOM] test number 5
  `

	postdata := slashred.PostData{
		"kind": "self",
		//"nsfw": "false",
		"sr":    "test",
		"text":  text,
		"title": "[RANODM] test made with some changes",
	}

	c.Listing.LinkSubmit(postdata)

}
