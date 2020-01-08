package slashred

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/hmble/slashred/internal"
	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	Name        string
	ProfileUrl  string
	IsPermanent bool
}

const (
	BaseAuthURL = "https://oauth.reddit.com"
	BaseURL     = "http://reddit.com"
)

const (
	CommentPrefix   = "t1_"
	AccountPrefix   = "t2_"
	LinkPrefix      = "t3_"
	MessagePrefix   = "t4_"
	SubredditPrefix = "t5_"
	AwardPrefix     = "t6_"
)

// TODO : Make token a member of client so that we don't need
// to pass token as parameter for every request we make.
type Client struct {
	Http      *http.Client
	Useragent string
	Token     *oauth2.Token
}

var NoAuthClient = &Client{
	Http: new(http.Client),
}

var auth *internal.Authenticator = internal.DefaultClient

func (u *User) UserClient(token *oauth2.Token) *Client {
	return &Client{
		Http:      auth.Config.Client(oauth2.NoContext, token),
		Useragent: auth.Useragent,
		Token:     token,
	}
}

func (u *User) Authenticate() (*oauth2.Token, error) {

	fmt.Println("Authentication starts from here:  ")
	fmt.Printf("Visit the url given below and paste the code given in url : \n %s", internal.AuthUrl(u.IsPermanent))

	fmt.Println("\n Enter the code here : ")

	var code string
	fmt.Scan(&code)

	token, err := internal.GetToken(code)

	if err != nil {
		log.Fatal("Error in getting token")
		return nil, err
	}

	SaveToken("token.json", token)

	return token, nil

}

func SaveToken(path string, token *oauth2.Token) {
	internal.SaveToken(path, token)
}

func UpdateToken(token *oauth2.Token) {
	internal.UpdateToken(token)
}

func TokenFromFile(filepath string) (*oauth2.Token, error) {
	return internal.TokenFromFile(filepath)
}

type Option map[string]string

var NoOptions Option = Option{}

func (c *Client) Get(endpoint string, opts Option) (res *http.Response, err error) {

	temp := BaseAuthURL + endpoint
	u, _ := url.Parse(temp)
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("raw_json", "1")

	for k, v := range opts {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	path := u.String()
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal("Error getting request")
		return nil, err
	}

	req.Header.Add("User-Agent", c.Useragent)

	str := fmt.Sprintf("bearer %s", c.Token.AccessToken)
	req.Header.Add("Authorization", str)

	fmt.Println(req.Header)

	return c.Http.Do(req)

}
