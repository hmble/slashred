package slashred

import (
	"bytes"
	"log"
	"strings"
)

type ListingService service

func (c *Client) getlisting(endpoint, logmsg string, opts Option) {
	resp, err := c.Get(endpoint, opts)

	if err != nil {
		log.Fatal(logmsg)
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}
func (l *ListingService) Listings() {
	resp, err := l.client.Get(API_PATH["trending_subreddits"], NoOptions)

	if err != nil {
		log.Fatal("Error in getting trending subreddits response")

	}

	defer resp.Body.Close()

	//SaveResponse(resp.Body, "test_data/trending_subreddits.json")
}

func (l *ListingService) ListingByID(names []string) {
	var buf bytes.Buffer
	buf.WriteString(API_PATH["by_id"])
	buf.WriteString(strings.Join(names, ","))

	url := buf.String()

	resp, err := l.client.Get(url, NoOptions)

	if err != nil {
		log.Fatal("Error in getting listings by name")
	}
	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/by_id.json")

}

// Should provide PostData with following keys
/*
kind(string):    one of (link, self, image, video, videogif)
nsfw:   true or false
sr:     name of subreddit
spoiler: true or false
text: text body (string)
title: string()
send_replies: true or false

// If kind is "link" then
link : a valid url
video_poster_url : url
// If wanted flair then
flair_id : a string no longer than 36 characters
flair_text : a string no longer than 64 characters

// Used for redirects
extensions: used for redirects
*/

func (l *ListingService) LinkSubmit(postdata PostData) {

	resp, er := l.client.Post(API_PATH["submit"], postdata)

	if er != nil {
		log.Fatal("Error in getting response of post body")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// fullname of link with LinkPrefix

func (c *Client) edit(endpoint, thingId, text string) {
	postdata := PostData{

		"thing_id": thingId,
		"text":     text,
	}

	resp, err := c.Post(endpoint, postdata)

	if err != nil {
		log.Fatal("Error in posting comment")
	}

	defer resp.Body.Close()
	PrintHeader(resp)

}

func (l *ListingService) EditUserText(thingId, text string) {
	l.client.edit(API_PATH["editusertext"], thingId, text)

}

// Test remained
func (l *ListingService) EventPostTime(start, end, timezone, fullname string) {
	postdata := PostData{
		"event_start": start,
		"event_end":   end,
		"event_tz":    timezone,
		"id":          fullname,
	}

	resp, err := l.client.Post(API_PATH["event_post_time"], postdata)

	if err != nil {
		log.Fatal("Error in editing event post time")
	}

	defer resp.Body.Close()

}

// Test remained

// follow : string(bool) [true/false]
func (l *ListingService) FollowPost(follow, fullname string) {
	postdata := PostData{
		"follow":   follow,
		"fullname": fullname,
	}
	resp, err := l.client.Post(API_PATH["follow_post"], postdata)

	if err != nil {
		log.Fatal("Error in following post")
	}

	defer resp.Body.Close()

}
