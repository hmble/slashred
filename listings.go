package slashred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (c *Client) Listings() {
	resp, err := c.Get(API_PATH["trending_subreddits"], NoOptions)

	if err != nil {
		log.Fatal("Error in getting trending subreddits response")

	}

	defer resp.Body.Close()

	//SaveResponse(resp.Body, "test_data/trending_subreddits.json")
}

func (c *Client) ListingByID(names []string) {
	var buf bytes.Buffer
	buf.WriteString(API_PATH["by_id"])
	buf.WriteString(strings.Join(names, ","))

	url := buf.String()

	resp, err := c.Get(url, NoOptions)

	if err != nil {
		log.Fatal("Error in getting listings by name")
	}
	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/by_id.json")

}

func (c *Client) GetComments(subreddit, article string) []CommentListing {
	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	resp, err := c.Get(endpoint, NoOptions)

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	fmt.Println("-------Got Reponse of comments------------")
	//	SaveResponse(resp.Body, "test_data/comments2.json")
	PrintHeader(resp)

	type listSub struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}

	// https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	result := make([]listSub, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("----got here after decode")
	if er != nil {
		panic(er)
		//log.Fatal("Error in decoding comments")
	}

	comments := result[1].Data.Children

	//commentReply := make([]Comment, 0)

	return comments
}

func (c *Client) GetCommentsID(article, comment string) {
	endpoint := fmt.Sprintf("/r/LifeProTips/comments/%s", article)

	// q.Add("comment", comment)
	// q.Add("depth", "7")

	options := Option{
		"comment": comment,
		"depth":   "7",
	}
	resp, err := c.Get(endpoint, options)

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	fmt.Println("-------Got Reponse of comments------------")
	//SaveResponse(resp.Body, "test_data/commeny_by_id.json")
	PrintHeader(resp)

	type ListSub struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}

	result := make([]ListSub, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("----got here after decode")
	if er != nil {
		panic(er)
		//log.Fatal("Error in decoding comments")
	}

	comments := result[1].Data.Children

	for _, comment := range comments {
		fmt.Printf("Author : %s\n", comment.Data.Author)

		replies := comment.Data.Replies.Data.Children

		for _, reply := range replies {

			fmt.Printf("Author : %s\n", reply.Comment.Author)
		}
	}

}

type PostData map[string]string

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

func (c *Client) LinkSubmit(postdata PostData) {

	resp, er := c.Post(API_PATH["submit"], postdata)

	if er != nil {
		log.Fatal("Error in getting response of post body")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (c *Client) UnmarkNsfw(id string) {
	postdata := PostData{
		"id": LinkPrefix + id,
	}

	resp, err := c.Post(API_PATH["unmarknsfw"], postdata)

	if err != nil {
		log.Fatal("Error in unmarking nsfw")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}
