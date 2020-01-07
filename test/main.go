package main

import (
	"log"
	"strings"

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

	//c.Listings()
	//names := strings.Split("t3_eiuti6", "_")
	names := strings.Split("t3_eioanh", "_")

	//	c.ListingByID(names)
	article := names[1]
	comment := "fcsnxlf"

	c.GetCommentsID(article, comment)

	// comments := c.GetComments(article)

	// for _, comment := range comments {

	// 	fmt.Printf("Comment: %s %d\n", comment.Data.Author, comment.Data.Depth)

	// 	replies := comment.Data.Replies.Data.Children

	// 	for _, reply := range replies {

	// 		if reply.Kind == "t1" {
	// 			fmt.Printf("\tReply: %s  --> %d\n", reply.Comment.Author, reply.Comment.Depth)

	// 			isEmpty := reply.Comment.Replies
	// 			tempstruct := &slashred.Replies{}
	// 			if &isEmpty != tempstruct {
	// 				fmt.Printf("\t\t\t size is %d\n", len(isEmpty.Data.Children))

	// 			}

	// 			fmt.Printf("\t\tID : %s  %s --> url is %s\n", reply.Comment.ID, reply.Comment.Author, reply.Comment.Url)
	// 		}
	// 	}

	// }
	// fmt.Println("Success.........")

}

