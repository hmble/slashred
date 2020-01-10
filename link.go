package slashred

import (
	"encoding/json"
	"fmt"
	"log"
)

// Need to look at this api schema afterwords for confirmations
//
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

func (c *Client) Best(opts Option) []SubmissionData {

	endpoint := API_PATH["best"]
	resp, err := c.Get(endpoint, opts)

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

func (c *Client) Hot(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/hot", subreddit)

	resp, err := c.Get(endpoint, opts)

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
func (c *Client) New(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/new", subreddit)

	resp, err := c.Get(endpoint, opts)

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

func (c *Client) Rising(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/rising", subreddit)

	resp, err := c.Get(endpoint, opts)

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

func (c *Client) Top(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/top", subreddit)

	resp, err := c.Get(endpoint, opts)

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

func (c *Client) Controversial(subreddit string, opts Option) []SubmissionData {

	endpoint := fmt.Sprintf("/r/%s/controversial", subreddit)

	resp, err := c.Get(endpoint, opts)

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
