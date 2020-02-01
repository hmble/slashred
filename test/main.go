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

	//	c.Users.AddFriend("astar0n", "workspc04")

	//eva9qo

	//comments := c.Comment.GetComments("golang", "eva9qo")
	commentsList := c.Comment.GetComments("AskReddit", "ev67eq", "best")
	////c.Comment.GetComments("AskReddit", "ev67eq")

	comments := make([]*slashred.Comment, 0)

	test := commentsList[1]

	if test.Comment != nil {
		replies := test.Comment.Replies.Data.Children

		//parent := slashred.CommentPrefix + test.Comment.ID
		reply := replies[0]

		if reply.Comment != nil {
			fmt.Printf("\t%s Parent[%s] ID[%s]\n", reply.Comment.Author,
				reply.Comment.Parent, reply.Comment.ID)
			reply2 := reply.Comment.Replies.Data.Children

			for _, r := range reply2 {
				if r.Comment != nil {
					fmt.Println(r.Comment.Author)
				} else {
					moreReplies := c.Comment.ReplaceMore(r.More, "ev67eq", "best", reply.Comment.ID)

					for _, m := range moreReplies {
						fmt.Printf("\t%s Parent[%s] ID[%s]\n", m.Author,
							m.Parent, m.ID)
					}
				}

			}

		} else {
			more := reply.More.Count

			fmt.Println(more)
		}

	}

	for _, comment := range comments {
		fmt.Println(comment.Author)
	}

	fmt.Println("Comments length ", len(comments))

}

