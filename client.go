package slashred

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
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

type Response struct {
	*http.Response

	// later more field will be added
}
const (
	BaseAuthURL = "https://oauth.reddit.com/"
	BaseURL     = "http://reddit.com/"
)

const (
	CommentPrefix   = "t1_"
	AccountPrefix   = "t2_"
	LinkPrefix      = "t3_"
	MessagePrefix   = "t4_"
	SubredditPrefix = "t5_"
	AwardPrefix     = "t6_"
)

// TODO(hmble): Add authenticated user info to client so that we can use it
// along different methods.
type Client struct {
	client      *http.Client
	UserAgent string
	// Token     *oauth2.Token
	
	x         ratelimit

	common service // Reuse same struct instead of creating

	Print      bool
	Account    *AccountService
	Collection *CollectionService
	Comment    *CommentService
	Flair      *FlairService
	Gold       *GoldService
	Link       *LinkService
	Listing    *ListingService
	Message    *MessageService
	Moderation *ModerationService
	Modmail    *ModmailService
	Modpost    *ModpostService
	Multis     *MultisService
	Report     *ReportService
	Subreddit  *SubredditService
	Users      *UsersService
	Wiki       *WikiService
}

type ratelimit struct {
	used      int
	remaining int
	reset     int
}
type service struct {
	client *Client
}

//var defaultAuth *internal.Authenticator = internal.DefaultClient

func NewClient(httpClient *http.Client, userAgent string) *Client {
	if httpClient == nil {
		// Default NoAuthClient
		httpClient = &http.Client{}
		
	}
	c := &Client{
		client: httpClient,
		UserAgent: userAgent,
	}
	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	// here we can't use service struct because we included `path` member in
	// CommentService
	c.Comment = &CommentService{client: c}
	c.Collection = (*CollectionService)(&c.common)
	c.Flair = (*FlairService)(&c.common)
	c.Gold = (*GoldService)(&c.common)
	c.Link = (*LinkService)(&c.common)
	c.Listing = (*ListingService)(&c.common)
	c.Message = (*MessageService)(&c.common)
	c.Moderation = (*ModerationService)(&c.common)
	c.Modmail = (*ModmailService)(&c.common)
	c.Modpost = (*ModpostService)(&c.common)
	c.Multis = (*MultisService)(&c.common)
	c.Report = (*ReportService)(&c.common)
	c.Subreddit = (*SubredditService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Wiki = (*WikiService)(&c.common)

	return c
}

type Option map[string]string

var NoOptions Option = Option{}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	baseURL,_:= url.Parse(BaseAuthURL)

	u, err := baseURL.Parse(urlStr)

	if err != nil {
		return nil, err
	}
	var buf io.Reader
	if body != nil {
		v, err := query.Values(body)
		
		if err != nil {
			return nil, err
		}
		v.Add("raw_json", "1")

		buf = strings.NewReader(v.Encode())
	}

	fmt.Println("The URL string is ", u.String())
	if buf == nil {
		fmt.Println("Its a get method")
	}
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	token,err := c.Token()
	if err != nil {
		return nil, errors.New("slashred: token not found")
	}
	str := fmt.Sprintf("bearer %s", token.AccessToken)
	req.Header.Add("Authorization", str)


		return req, nil
	}

	func (c *Client) BareDo(req *http.Request) (*Response, error) {
		// Later we will do more things here 
		resp, err := c.client.Do(req)

		if err != nil {
			log.Println("From BareDO ")
			defer resp.Body.Close()
			return nil, err
		}

		response := newResponse(resp)
		return response, nil
		
	}
	func (c *Client) Do(req *http.Request, v interface{}) (*Response, error){
		resp, err := c.BareDo(req)

		if err != nil {
			log.Println("From Do")
			return resp, err
		}
		defer resp.Body.Close()

		switch v := v.(type) {
		case nil:
		case io.Writer:
						_, err = io.Copy(v, resp.Body)
		default:
						decErr := json.NewDecoder(resp.Body).Decode(v)
						if decErr == io.EOF {
							decErr = nil // ignore EOF errors caused by empty response body
						}
						if decErr != nil {
							err = decErr
						}
	}
		return resp, err
	}

	func newResponse(r *http.Response) *Response {
		response := &Response{Response: r}

		return response
	}
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

	// req.Header.Add("User-Agent", c.Useragent)
	token,err := c.Token()
	if err != nil {
		return nil, errors.New("slashred: token not found")
	}
	str := fmt.Sprintf("bearer %s", token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.client.Do(req)

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

	// req.Header.Add("User-Agent", c.Useragent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return c.client.Do(req)

}

// Reference
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html
func (c *Client) PostImageUpload(endpoint string, postdata PostData, filename string) (*http.Response, error) {

	fullurl := BaseAuthURL + endpoint

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for k, v := range postdata {
		bodyWriter.WriteField(k, v)
	}

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)

	if err != nil {
		log.Fatal("error in writing to buffer from fileWriter:  ", err)

		return nil, err

	}

	fh, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error in openeing file ", err)
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	body := bytes.NewBufferString(bodyBuf.String())
	req, err := http.NewRequest(http.MethodPost, fullurl, body)

	if err != nil {
		return nil, err
	}

	// req.Header.Add("User-Agent", c.Useragent)
	req.Header.Add("Content-Type", contentType)
	token,err := c.Token()
	if err != nil {
		return nil, errors.New("slashred: token not found")
	}
	str := fmt.Sprintf("bearer %s", token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.client.Do(req)

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

	// req.Header.Add("User-Agent", c.Useragent)

	token,err := c.Token()
	if err != nil {
		return nil, errors.New("slashred: token not found")
	}
	str := fmt.Sprintf("bearer %s", token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.client.Do(req)

}

func (c *Client) Put(endpoint string, postdata PostData) (*http.Response, error) {
	fullurl := BaseAuthURL + endpoint

	data := url.Values{}

	data.Set("api_type", "json")

	for k, v := range postdata {
		data.Set(k, v)
	}
	body := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest(http.MethodPut, fullurl, body)

	if err != nil {
		return nil, err
	}

	// req.Header.Add("User-Agent", c.Useragent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	return c.client.Do(req)

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

func (c *Client) SetPrint() {
	if !c.Print {
		c.Print = true
	}
}
