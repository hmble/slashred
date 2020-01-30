package main

import (
	"fmt"
	"log"
	"time"

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

	//	c.Users.AddFriend("astar0n", "workspc04")

	//eva9qo

	start := time.Now()
	//comments := c.Comment.GetComments("golang", "eva9qo")
	comments := c.Comment.GetComments("AskReddit", "ev67eq")
	////c.Comment.GetComments("AskReddit", "ev67eq")

	for _, comment := range comments {
		fmt.Printf("%s ID : [%s]\n", comment.Author, comment.ID)

		//replies := comment.Data.Replies
		//if comment.Data.Replies != {
		replies := comment.Replies.Data.Children
		//	replies := comment.Replies.ReplyArray

		fmt.Printf("Replies length := %d\n", len(replies))
		for _, reply := range replies {
			if reply.Kind != "more" {
				fmt.Printf("\t >> %s  Parent[%s]\n", reply.Comment.Author, reply.Comment.Parent)
			} else {

				//fmt.Println(reply.More.ParentID)
				comments := c.Comment.ReplaceMore(reply.More, slashred.LinkPrefix+"ev67eq")

				for _, comment := range comments {

					fmt.Printf("\t >> %s  Parent: [%s]\n", comment.Author, comment.Parent)

				}
				// fmt.Println("Index: ", i, children[:4])
				// fmt.Println(strings.Join(children[:4], ","))
			}

		}
	}

	//str := "ffuv8u8,ffv2iku,ffvlk7l,ffv3mdv"

	//// c.Comment.ReplaceMore(str, slashred.NoOptions, "ev67eq")

	//c.Comment.ReplaceMore(, slashred.PostData{}, slashred.LinkPrefix+"ev67eq")
	////		break

	//c.Comment.ReplaceMore(replies[)

	elapesed := time.Since(start)

	fmt.Printf("Time elapsed since making request ==> %s\n", elapesed)

}
