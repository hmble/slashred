package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hmble/slashred"
	_ "github.com/joho/godotenv/autoload"
)

const redirectURI = "http://localhost:8080/callback"

var (
	userAgent = os.Getenv("USER_AGENT")
	auth  = slashred.NewAuthenticatior(redirectURI, userAgent,slashred.Scopes...)
	ch    = make(chan *slashred.Client)
	state = "abc123"
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state, true)
	fmt.Println("Please log in to slashred by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	t, err := client.Token()

	slashred.SaveToken("newtoken.json", t)
	if err != nil {
		log.Println("Error in getting token")
	}
	fmt.Printf("%+v\n", t)
	
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok, userAgent)
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}