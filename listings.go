package slashred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func (c *Client) Listings() {
	resp, err := c.Get(API_PATH["trending_subreddits"])

	if err != nil {
		log.Fatal("Error in getting trending subreddits response")

	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/trending_subreddits.json")
}

func (c *Client) ListingByID(names []string) {
	var buf bytes.Buffer
	buf.WriteString(API_PATH["by_id"])
	buf.WriteString(strings.Join(names, ","))

	url := buf.String()

	resp, err := c.Get(url)

	if err != nil {
		log.Fatal("Error in getting listings by name")
	}
	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/by_id.json")

}

func (c *Client) GetComments(article string) []CommentListing {
	endpoint := fmt.Sprintf("/r/LifeProTips/comments/%s", article)
	u, _ := url.Parse(endpoint)
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("raw_json", "1")

	u.RawQuery = q.Encode()

	resp, err := c.Get(u.String())

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	fmt.Println("-------Got Reponse of comments------------")
	//	SaveResponse(resp.Body, "test_data/comments2.json")
	PrintHeader(resp)

	type ListSub struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}

	// https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	result := make([]ListSub, 0)
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
	u, _ := url.Parse(endpoint)
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("raw_json", "1")
	q.Add("comment", comment)
	q.Add("depth", "7")

	u.RawQuery = q.Encode()

	resp, err := c.Get(u.String())

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	fmt.Println("-------Got Reponse of comments------------")
	//SaveResponse(resp.Body, "test_data/commeny_by_id.json")
	PrintHeader(resp)

	fmt.Println("Respnse saved")
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

