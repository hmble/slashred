package main

import (
	"fmt"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// authenticator := &slashred.Authenticator{
	// 	Config: &oauth2.Config{
	// 		ClientID:     os.Getenv("CLIENT_ID"),
	// 		ClientSecret: os.Getenv("CLIENT_SECRET"),
	// 		Scopes:       slashred.Scopes,
	// 		Endpoint:     slashred.Endpoint,
	// 		RedirectURL:  "https://example.com/auth",
	// 	},
	// 	Useragent: os.Getenv("USER_AGENT"),
	// }
	// u := &slashred.User{
	// 	Name:          "Samyak",
	// 	ProfileUrl:    "",
	// 	IsPermanent:   true,
	// 	Authenticator: authenticator,
	// }

	// token, er := slashred.TokenFromFile("token.json")
	// if er != nil {
	// 	log.Fatal("Error in reading token")
	// }

	// u.UpdateToken(token)

	// c := u.UserClient(token)

	// account, err := c.Account.GetMe()

	// if err != nil {
	// 	log.Println("Error in getting account info")
	// }

	// fmt.Println("Name is ",  account.Name)
	// karma, err := c.Account.GetKarma()

	// if err != nil {
	// 	log.Println("Error in getting karma")
	// }
	// var totalCount int
	// for _, k := range karma {
	// 	totalCount += k.CommentKarma + k.LinkKarma
	// }

	// fmt.Println("Total Karma is ", totalCount)

	auth := slashred.NewAuthenticatior("https://example.com/auth", "read")

	fmt.Println(auth.AuthURL("1234", true))
}
