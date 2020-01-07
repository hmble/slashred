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

func (c *Client) GetComments(article string) {
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

	for _, comment := range comments {

		fmt.Printf("Comment: %s  --> %s\n", comment.Data.Author, comment.Kind)

		replies := comment.Data.Replies.Data.Children

		for _, reply := range replies {

			if reply.Kind == "t1" {
				fmt.Printf("\tReply: %s  --> %s\n", reply.Comment.Author, reply.Kind)
			} else {
				fmt.Printf("\t More: %s --> %s\n", reply.More.Name, reply.Kind)
			}
		}
	}

}

