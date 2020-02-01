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

	// comments = append(comments, test.Comment)

	// if test.Comment != nil {
	// 	replies := test.Comment.Replies.Data.Children

	// 	for _, reply := range replies {
	// 		if reply.Comment != nil {
	// 			comments = append(comments, reply.Comment)

	// 		} else {

	// 			moreReplies := c.Comment.ReplaceMore(reply.More, "ev67eq", "best")

	// 			// for _, more := range moreReplies {
	// 			// 	fmt.Println("From more ", more.Author)
	// 			// }

	// 			comments = append(comments, moreReplies...)

	// 		}
	// 	}
	// }

	// for _, comment := range comments {
	// 	fmt.Println(comment.Author)
	// }

	if test.Comment != nil {
		replies := test.Comment.Replies.Data.Children

		//parent := slashred.CommentPrefix + test.Comment.ID
		for _, reply := range replies {
			if reply.Comment != nil {
				fmt.Printf("\t %s Parent[%s], ID[%s]\n", reply.Comment.Author,
					reply.Comment.Parent, reply.Comment.ID)

			} else {
				moreReplies := c.Comment.ReplaceMore(reply.More, "ev67eq", "best", test.Comment.ID)
				for _, more := range moreReplies {
					//if more.Parent == parent {
					fmt.Printf("\t %s Parent[%s], ID[%s]\n", more.Author,
						more.Parent, more.ID)
					//	}
				}

			}
		}

	}

	for _, comment := range comments {
		fmt.Println(comment.Author)
	}

	fmt.Println("Comments length ", len(comments))

}

