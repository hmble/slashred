package slashred

import (
	"encoding/json"
	"fmt"
	"log"
)

type FlairService service

type Flair struct {
	CssClass     string `json:"css_class"`
	ID           string `json:"id"`
	Text         string `json:"text"`
	TextEditable bool   `json:"text_editable"`
	Type         string `json:"type"`
	//TODO(hmble) : RichText field is not implemented?
}

type Flairv2 struct {
	Flair
	AllowableContent string `json:"allowable_content"`
	BackgroundColor  string `json:"background_color"`
	MaxEmojis        int    `json:"max_emojis"`
	TextColor        string `json:"text_color"`
}

// Get flair list
// Flair.List
func (f *FlairService) List(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/api/flairlist", subreddit)

	resp, err := f.client.Get(path, opts)

	if err != nil {
		log.Fatal("Error in getting flairlist response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/flairlist.json")
}

// Flair.LinkFlair
func (f *FlairService) LinkFlair(subreddit string) []*Flair {
	path := fmt.Sprintf("/r/%s/api/link_flair", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting flairlist response")
	}

	defer resp.Body.Close()

	flairs := make([]*Flair, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs
}

// Flair.LinkFlairV2
func (f *FlairService) LinkFlairV2(subreddit string) []*Flairv2 {
	path := fmt.Sprintf("/r/%s/api/link_flair_v2", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting flairlist response")
	}

	defer resp.Body.Close()

	flairs := make([]*Flairv2, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs
}

func (f *FlairService) UserFlair(subreddit string) []*Flair {
	path := fmt.Sprintf("/r/%s/api/user_flair", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting flairlist response")
	}

	defer resp.Body.Close()
	flairs := make([]*Flair, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs

}

// Flair.UserFlairV2
func (f *FlairService) UserFlairV2(subreddit string) []*Flairv2 {
	path := fmt.Sprintf("/r/%s/api/user_flair_v2", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting flairlist response")
	}

	defer resp.Body.Close()
	flairs := make([]*Flairv2, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs

}

