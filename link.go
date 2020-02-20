package slashred

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Need to look at this api schema afterwords for confirmations
//

type LinkService service
type Submission struct {
	ApprovedBy          string  `json:"approved_by"`
	Archived            bool    `json:"archived"`
	Author              string  `json:"author"`
	AuthorFlairCSSClass string  `json:"author_flair_css_class"`
	AuthorFlairText     string  `json:"author_flair_text"`
	BannedBy            string  `json:"banned_by"`
	Clicked             bool    `json:"clicked"`
	ContestMode         bool    `json:"contest_mode"`
	Created             float64 `json:"created"`
	CreatedUtc          float64 `json:"created_utc"`
	Distinguished       string  `json:"distinguished"`
	Domain              string  `json:"domain"`
	Downs               int     `json:"downs"`
	// when not edited api returns bool of false
	// when it is edited api returns float64 epoch timestamp
	Edited interface{} `json:"edited"`

	Gilded            int    `json:"gilded"`
	Hidden            bool   `json:"hidden"`
	HideScore         bool   `json:"hide_score"`
	ID                string `json:"id"`
	IsSelf            bool   `json:"is_self"`
	Likes             bool   `json:"likes"`
	LinkFlairCSSClass string `json:"link_flair_css_class"`
	LinkFlairText     string `json:"link_flair_text"`
	Locked            bool   `json:"locked"`
	// Media               Media         `json:"media"`
	MediaEmbed       interface{}   `json:"media_embed"`
	ModReports       []interface{} `json:"mod_reports"`
	Name             string        `json:"name"`
	NumComments      int           `json:"num_comments"`
	NumReports       int           `json:"num_reports"`
	Over18           bool          `json:"over_18"`
	Permalink        string        `json:"permalink"`
	Quarantine       bool          `json:"quarantine"`
	RemovalReason    interface{}   `json:"removal_reason"`
	ReportReasons    []interface{} `json:"report_reasons"`
	Saved            bool          `json:"saved"`
	Score            int           `json:"score"`
	SecureMedia      interface{}   `json:"secure_media"`
	SecureMediaEmbed interface{}   `json:"secure_media_embed"`
	SelftextHTML     string        `json:"selftext_html"`
	Selftext         string        `json:"selftext"`
	Stickied         bool          `json:"stickied"`
	Subreddit        string        `json:"subreddit"`
	SubredditID      string        `json:"subreddit_id"`
	SuggestedSort    string        `json:"suggested_sort"`
	Thumbnail        string        `json:"thumbnail"`
	Title            string        `json:"title"`
	URL              string        `json:"url"`
	Ups              int           `json:"ups"`
	UserReports      []interface{} `json:"user_reports"`
	Visited          bool          `json:"visited"`
}

type SubmissionData struct {
	Kind string
	Data Submission
}

var listSub struct {
	Kind string
	Data struct {
		Children []SubmissionData
		Before   string
		After    string
	}
}

func (l *ListingService) Best(opts Option) []SubmissionData {

	endpoint := API_PATH["best"]
	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal("Error in getting best response")
	}

	defer resp.Body.Close()
	//	SaveResponse(resp.Body, "test_data/best.json")

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children
}

func (l *ListingService) Hot(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/hot", subreddit)

	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children
}

// Test remained
func (l *ListingService) New(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/new", subreddit)

	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children
}

func (l *ListingService) Rising(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/rising", subreddit)

	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children

}

func (l *ListingService) Top(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/top", subreddit)

	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children

}

func (l *ListingService) Controversial(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/controversial", subreddit)

	resp, err := l.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	er := json.NewDecoder(resp.Body).Decode(&listSub)

	if er != nil {
		log.Fatal(er)
	}

	return listSub.Data.Children

}

// TODO : Implement MoreChildren method

func (l *LinkService) Unsave(fullname string) {
	l.client.unlink(API_PATH["unsave"], fullname)
}

func (l *LinkService) Save(category, fullname string) {
	postdata := PostData{
		"category": category,
		"id":       fullname,
	}
	resp, err := l.client.Post(API_PATH["vote"], postdata)

	if err != nil {
		log.Fatal("Errro in casting vote")
	}
	defer resp.Body.Close()
}

// fullname of link or comment
func (l *LinkService) DeleteLink(fullname string) {
	l.client.unlink(API_PATH["delete"], fullname)
}
func (c *Client) vote(direction, fullname string) {
	postdata := PostData{

		"dir": direction,
		"id":  fullname,
	}

	resp, err := c.Post(API_PATH["vote"], postdata)

	if err != nil {
		log.Fatal("Errro in casting vote")
	}
	defer resp.Body.Close()

	PrintHeader(resp)
}

func (l *LinkService) Upvote(fullname string) {
	l.client.vote("1", fullname)
}

func (l *LinkService) ClearVote(fullname string) {
	l.client.vote("0", fullname)
}

func (l *LinkService) Downvote(fullname string) {
	l.client.vote("-1", fullname)
}

// Not tested
// Link.ApiInfo

func (l *LinkService) ApiInfo(subreddit, id, url string) {
	path := fmt.Sprintf("/r/%s/api/info", subreddit)

	opts := Option{
		"id":  id,
		"url": url,
	}

	resp, err := l.client.Get(path, opts)

	if err != nil {
		log.Fatalf("Error in getting response for path : %s\n", path)
	}

	defer resp.Body.Close()
}

// Link.SavedCategories

func (l *LinkService) SavedCategories() {
	path := "/api/saved_categories"

	resp, err := l.client.Get(path, NoOptions)

	if err != nil {
		log.Fatalf("Error in getting response for path : %s\n", path)
	}

	defer resp.Body.Close()
}

// Requires Premium
// Link.StoreVisits

func (l *LinkService) StoreVisits(links []string) {
	path := "/api/store_visits"

	postdata := PostData{
		"links": strings.Join(links, ","),
	}

	resp, err := l.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for path : %s\n", path)
	}

	defer resp.Body.Close()

}
