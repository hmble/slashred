package slashred

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const (	
	AuthURL =   "https://www.reddit.com/api/v1/authorize"
	TokenURL = "https://www.reddit.com/api/v1/access_token"

)

type Auth struct {
	config  *oauth2.Config
	context context.Context
}

type authTransport struct{
	config *oauth2.Config
	useragent string
}

func (t authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.useragent)
	return http.DefaultTransport.RoundTrip(req)
}

// Return new Authencticator
func NewAuthenticatior(redirectURL string, userAgent string, scopes ...string, ) Auth {

	if redirectURL == "" {
		log.Fatal("Redirect URL cannot be empty")
	}
	if userAgent == "" {
		log.Fatal("User Agent cannot be empty")
	}
	cfg := &oauth2.Config{
		ClientID: os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: scopes,
		RedirectURL: redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL: AuthURL,
			TokenURL: TokenURL,
		},
	}

	tr := authTransport{config: cfg, useragent: userAgent}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: tr})

	return Auth{config: cfg, context: ctx}
}

// SetAuthInfo overwrites the client ID and secret key used by the authenticator.
// You can use this if you don't want to store this information in environment variables.
func (a *Auth) SetAuthInfo(clientID, clientSecret string) {
	a.config.ClientID = clientID
	a.config.ClientSecret = clientSecret
}

func (a Auth) AuthURL(state string, isPermanent bool) string {
	if isPermanent {

		return a.config.AuthCodeURL(state, oauth2.SetAuthURLParam("duration", "permanent"))
	}else {
		return a.config.AuthCodeURL(state, oauth2.SetAuthURLParam("duration", "temporary"))
	}
}

// Gets Token from server ...
func (a Auth) Token(state string, r *http.Request) (*oauth2.Token, error) {
	values := r.URL.Query()

	code := values.Get("code")

	if code == "" {
		return nil, errors.New("slashred: didn't get access code")
	}

	actualState := values.Get("state")

	if actualState != state {
		return nil, errors.New("slashred: redirect state parameter doesn't match")
	}

	return a.config.Exchange(a.context, code)
}

// Exchange is like Token, except it allows you to manually specify
// access code instead of pulling it from HTTP request
func (a Auth) Exchange(code string) (*oauth2.Token, error) {
	return a.config.Exchange(a.context, code)
}

// returns the client
func (a Auth) NewClient(token *oauth2.Token, userAgent string) *Client {
	client := a.config.Client(a.context, token)

	return NewClient(client,userAgent)
}

func (c *Client) Token() (*oauth2.Token, error) {
	tr, ok := c.client.Transport.(*oauth2.Transport)

	if !ok {
		return nil, errors.New("slashred: oauth2 transport type is not correct")
	}
	
	t, err := tr.Source.Token()

	if err != nil {
		return nil, err
	}

	return t, err
}