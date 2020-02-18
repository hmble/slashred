package main

import (
  "fmt"
	"log"
	"os"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {

	authenticator := &slashred.Authenticator{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       slashred.Scopes,
			Endpoint:     slashred.Endpoint,
			RedirectURL:  "https://example.com/auth",
		},
		Useragent: os.Getenv("USER_AGENT"),
	}
	u := &slashred.User{
		Name:          "Samyak",
		ProfileUrl:    "",
		IsPermanent:   true,
		Authenticator: authenticator,
	}

	token, er := slashred.TokenFromFile("token.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	u.UpdateToken(token)

	var c *slashred.Client = u.UserClient(token)

  flairs := c.Flair.UserFlair("linux")

  for _, flair := range flairs {
    fmt.Println(flair.Text)
  }


  fmt.Println("-----Testing userflair v2")

  flairs2 := c.Flair.UserFlairV2("linux")

  for _, flair := range flairs2 {
    fmt.Println(flair.BackgroundColor)
  }
	

}
