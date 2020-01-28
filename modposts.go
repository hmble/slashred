package slashred

import (
	"log"
	"strings"
)

//

type ModpostService service

func (m *ModpostService) SetContestMode(fullname, state string) {

	postdata := PostData{
		"id":    fullname,
		"state": state,
	}
	resp, err := m.client.Post(API_PATH["set_contest_mode"], postdata)

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
func (m *ModpostService) SetLinkSticky(fullname, num, state, toProfile string) {

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
	resp, err := m.client.Post(API_PATH["set_subreddit_sticky"], postdata)

	if err != nil {
		log.Fatal("Error in setting contest mode")
	}

	defer resp.Body.Close()
}

// Test Remained
// sort : one of (confidence, top, new, controversial, old, random,
// qa, live, blank)
func (m *ModpostService) SetSuggestedSort(fullname, sort string) {
	postdata := PostData{
		"id":   fullname,
		"sort": sort,
	}
	resp, err := m.client.Post(API_PATH["set_suggested_sort"], postdata)

	if err != nil {
		log.Fatal("Error in setting suggested sort")
	}

	defer resp.Body.Close()
}

func (m *ModpostService) Spoiler(fullname string) {
	m.client.unlink(API_PATH["spoiler"], fullname)
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
func (m *ModpostService) Unmarknsfw(fullname string) {
	m.client.unlink(API_PATH["unmarknsfw"], fullname)
}

// fullname of link with LinkPrefix
func (m *ModpostService) Marknsfw(fullname string) {
	m.client.unlink(API_PATH["marknsfw"], fullname)
}

// fullname of link with LinkPrefix
func (m *ModpostService) Unhide(fullname string) {
	m.client.unlink(API_PATH["unhide"], fullname)
}

// fullname of link with LinkPrefix
func (m *ModpostService) Hide(listOfFullnames []string) {
	fullname := strings.Join(listOfFullnames, ",")
	m.client.unlink(API_PATH["hide"], fullname)
}

// fullname of link with LinkPrefix
func (m *ModpostService) Unlock(fullname string) {
	m.client.unlink(API_PATH["unlock"], fullname)
}
func (m *ModpostService) Lock(fullname string) {
	m.client.unlink(API_PATH["lock"], fullname)
}

// fullname of link with LinkPrefix
func (m *ModpostService) Unspoiler(fullname string) {
	m.client.unlink(API_PATH["unspoiler"], fullname)
}
