package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/hmble/slashred"
)

func main() {

	// url := internal.RedditClient.GetUrl()

	// fmt.Printf("\nUrl is: \n%s", url)

	// fmt.Println("\nEnter code here: ")
	// var code string

	// fmt.Scan(&code)

	// token, err := internal.RedditClient.GetToken(code)
	// if err != nil {

	// 	log.Fatal("Error in getting token")
	// }

	// internal.SaveToken("token.json", token)

	// fmt.Println("token saved....")

	u := &slashred.User{
		Name:        "Samyak",
		ProfileUrl:  "",
		IsPermanent: true,
	}
	token, er := slashred.TokenFromFile("token.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	//	var a *internal.Authenticator = internal.RedditClient

	slashred.UpdateToken(token)

	// internal.SaveToken("token.json", token)

	// fmt.Println("Token saved")
	var c *slashred.Client = u.UserClient(token)

	trophies, err := c.Trophies()

	if err != nil {
		log.Fatal("Error in trophies function")
	}

	for _, trophy := range trophies {
		fmt.Println(trophy.Name)
	}

	// pref, _ := c.GetMyPreferences()

	// fmt.Println(pref.VideoAutoplay)
	fmt.Println("Success.........")

}
