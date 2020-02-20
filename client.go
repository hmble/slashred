package slashred

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	//_ "github.com/joho/godotenv/autoload"
)

var Scopes = []string{
	"edit",
	"flair",
	"history",
	"identity",
	"modconfig",
	"modflair",
	"modlog",
	"modposts",
	"modwiki",
	"mysubreddits",
	"privatemessages",
	"read",
	"report",
	"save",
	"submit",
	"subscribe",
	"vote",
	"wikiedit",
	"wikiread",
}

type User struct {
	Name          string
	ProfileUrl    string
	IsPermanent   bool
	Authenticator *Authenticator
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

type Client struct {
	Http      *http.Client
	Useragent string
	Token     *oauth2.Token
	x         ratelimit

	common service // Reuse same struct instead of creating

	Account    *AccountService
	Collection *CollectionService
	Comment    *CommentService
	Flair      *FlairService
	Gold       *GoldService
	Link       *LinkService
	Listing    *ListingService
	Moderation *ModerationService
	Modmail    *ModmailService
	Modpost    *ModpostService
	Report     *ReportService
	Subreddit  *SubredditService
	Users      *UsersService
}

type ratelimit struct {
	used      int
	remaining int
	reset     int
}
type service struct {
	client *Client
}

var NoAuthClient = &Client{
	Http: new(http.Client),
}

//var defaultAuth *internal.Authenticator = internal.DefaultClient

func (u *User) UserClient(token *oauth2.Token) *Client {
	c := &Client{
		Http:      u.Authenticator.Config.Client(oauth2.NoContext, token),
		Useragent: u.Authenticator.Useragent,
		Token:     token,
	}

	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	c.Comment = (*CommentService)(&c.common)
	c.Collection = (*CollectionService)(&c.common)
	c.Flair = (*FlairService)(&c.common)
	c.Gold = (*GoldService)(&c.common)
	c.Link = (*LinkService)(&c.common)
	c.Listing = (*ListingService)(&c.common)
	c.Moderation = (*ModerationService)(&c.common)
	c.Modmail = (*ModmailService)(&c.common)
	c.Modpost = (*ModpostService)(&c.common)
	c.Report = (*ReportService)(&c.common)
	c.Subreddit = (*SubredditService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	return c
}

func (u *User) Authenticate() (*oauth2.Token, error) {

	fmt.Println("Authentication starts from here:  ")
	fmt.Printf("Visit the url given below and paste the code given in url : \n %s", AuthUrl(true, u.Authenticator))

	fmt.Println("\n Enter the code here : ")

	var code string
	fmt.Scan(&code)

	token, err := GetToken(code, u.Authenticator)

	if err != nil {
		log.Fatal("Error in getting token")
		return nil, err
	}

	u.SaveToken("token.json", token)

	return token, nil

}

func (u *User) SaveToken(path string, token *oauth2.Token) {
	SaveToken(path, token)
}

func (u *User) UpdateToken(token *oauth2.Token) {
	UpdateToken(token, u.Authenticator)
}

// func TokenFromFile(filepath string) (*oauth2.Token, error) {
// 	return TokenFromFile(filepath)
// }

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

	// if c.x.remaining < 10 {
	// 	log.Fatal("---YOUR LIMIT HAS EXTENDED Wait for ", c.x.reset)
	// }

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

type PostData map[string]string

var NoPostdata = PostData{}

func (c *Client) Post(endpoint string, postdata PostData) (*http.Response, error) {
	data := url.Values{}

	fullurl := BaseAuthURL + endpoint
	data.Set("api_type", "json")

	for k, v := range postdata {
		data.Set(k, v)
	}
	body := bytes.NewBufferString(data.Encode())
	req, err := http.NewRequest(http.MethodPost, fullurl, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.Useragent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return c.Http.Do(req)

}

func (c *Client) Delete(endpoint string, opts Option) (res *http.Response, err error) {

	temp := BaseAuthURL + endpoint
	u, _ := url.Parse(temp)
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("raw_json", "1")

	for k, v := range opts {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	path := u.String()
	req, err := http.NewRequest("DELETE", path, nil)

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

func (c *Client) Put(endpoint string, data string) (*http.Response, error) {
	//data := url.Values{}

	fullurl := BaseAuthURL + endpoint
	////data.Set("api_type", "json")

	//for k, v := range postdata {
	//	data.Set(k, v)
	//}
	//body := bytes.NewBufferString(data.Encode())

	body := strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPut, fullurl, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.Useragent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	return c.Http.Do(req)

}

func (c *Client) savelimit(resp *http.Response) {
	used, errUser := strconv.Atoi(resp.Header.Get("X-Ratelimit-Used"))
	if errUser != nil {
		log.Fatal("Error in converting ratelimit used")
	}

	remaining, errRemaining := strconv.ParseFloat(resp.Header.Get("X-Ratelimit-Remaining"), 32)
	if errRemaining != nil {
		log.Fatal("Error in converting ratelimit remaining")
	}

	reset, errReset := strconv.Atoi(resp.Header.Get("X-Ratelimit-Reset"))
	if errReset != nil {
		log.Fatal("Error in converting ratelimit reset")
	}

	c.x = ratelimit{
		used:      used,
		remaining: int(remaining),
		reset:     reset,
	}
}
