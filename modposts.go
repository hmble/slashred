package slashred

import (
	"log"
	"strings"
)

func (c *Client) SetContestMode(fullname, state string) {

	postdata := PostData{
		"id":    fullname,
		"state": state,
	}
	resp, err := c.Post(API_PATH["set_contest_mode"], postdata)

	if err != nil {
		log.Fatal("Error in setting contest mode")
	}

	defer resp.Body.Close()
}

// Test remained
// https://www.reddit.com/dev/api/#POST_api_set_subreddit_sticky
// state is bool
// toProfile is bool
// num parameter is optional
func (c *Client) SetLinkSticky(fullname, num, state, toProfile string) {

	number := ""
	if num != "" {
		number = num
	}
	postdata := PostData{
		"id":         fullname,
		"num":        number,
		"state":      state,
		"to_profile": toProfile,
	}
	resp, err := c.Post(API_PATH["set_subreddit_sticky"], postdata)

	if err != nil {
		log.Fatal("Error in setting contest mode")
	}

	defer resp.Body.Close()
}

// Test Remained
// sort : one of (confidence, top, new, controversial, old, random,
// qa, live, blank)
func (c *Client) SetSuggestedSort(fullname, sort string) {
	postdata := PostData{
		"id":   fullname,
		"sort": sort,
	}
	resp, err := c.Post(API_PATH["set_suggested_sort"], postdata)

	if err != nil {
		log.Fatal("Error in setting suggested sort")
	}

	defer resp.Body.Close()
}

func (c *Client) Spoiler(fullname string) {
	c.unlink(API_PATH["spoiler"], fullname)
}

// Unmarknsfw
func (c *Client) unlink(endpoint, fullname string) {
	postdata := PostData{
		"id": fullname,
	}

	resp, err := c.Post(endpoint, postdata)

	if err != nil {
		log.Fatal("Error in unmarking nsfw")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// fullname of link with LinkPrefix
func (c *Client) Unmarknsfw(fullname string) {
	c.unlink(API_PATH["unmarknsfw"], fullname)
}

// fullname of link with LinkPrefix
func (c *Client) Marknsfw(fullname string) {
	c.unlink(API_PATH["marknsfw"], fullname)
}

// fullname of link with LinkPrefix
func (c *Client) Unhide(fullname string) {
	c.unlink(API_PATH["unhide"], fullname)
}

// fullname of link with LinkPrefix
func (c *Client) Hide(listOfFullnames []string) {
	fullname := strings.Join(listOfFullnames, ",")
	c.unlink(API_PATH["hide"], fullname)
}

// fullname of link with LinkPrefix
func (c *Client) Unlock(fullname string) {
	c.unlink(API_PATH["unlock"], fullname)
}
func (c *Client) Lock(fullname string) {
	c.unlink(API_PATH["lock"], fullname)
}

// fullname of link with LinkPrefix
func (c *Client) Unspoiler(fullname string) {
	c.unlink(API_PATH["unspoiler"], fullname)
}
