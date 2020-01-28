package slashred

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type CommentService service
type CommentListing struct {
	Kind string
	Data Comment `json:"data"`
}
type Data struct {
	Children []CommentListing
}

type Comment struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	Body   string `json:"body"`
	Depth  int    `json:"depth"`
	Url    string `json:"permalink"`
	LinkID string `json:"link_id"`
	ID     string `json:"id"`

	Replies Replies `json:"replies"`
	//	Replies Replies `json:"replies"`
}

type Replies struct {
	Data ReplyData `json:"data"`
}

type ReplyData struct {
	Children []RepliesArray
}
type More struct {
	Count    int
	Name     string
	ParentID string
	ID       string
	Depth    int
	Children []string
}
type RepliesArray struct {
	Kind    string
	Data    json.RawMessage `json:"data"`
	Comment Comment
	More    More
}

func (r *RepliesArray) UnmarshalJSON(b []byte) error {

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	kind := string(tmp["kind"])

	r.Kind = kind[1 : len(kind)-1]

	//r.Kind = string(tmp["kind"])

	// var comment Comment
	// var more More

	if string(tmp["kind"]) == "\"t1\"" {
		var comment Comment
		if err := json.Unmarshal(tmp["data"], &comment); err != nil {
			return err
		}
		r.Comment = comment

	} else {
		var more More
		if err := json.Unmarshal(tmp["data"], &more); err != nil {
			return err
		}
		r.More = more
	}

	return nil
}

func (r *Replies) UnmarshalJSON(b []byte) error {
	if string(b) == "\"\"" {
		return nil
	}

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return errors.New("Error in unmarshaling raw message")
	}

	var data ReplyData

	if err := json.Unmarshal(tmp["data"], &data); err != nil {
		return errors.New("Data point does not exists")

	}

	r.Data = data

	return nil

}
func (c *CommentService) GetComments(subreddit, article string) []CommentListing {
	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	resp, err := c.client.Get(endpoint, NoOptions)

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

func (c *CommentService) GetCommentsID(article, comment string) {
	endpoint := fmt.Sprintf("/r/LifeProTips/comments/%s", article)

	// q.Add("comment", comment)
	// q.Add("depth", "7")

	options := Option{
		"comment": comment,
		"depth":   "7",
	}
	resp, err := c.client.Get(endpoint, options)

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
func (c *CommentService) PostComment(thingId, text string) {
	c.client.edit(API_PATH["comment"], thingId, text)
}

// Test Remained
// state is bool
func (c *CommentService) SendReplies(fullname, state string) {
	postdata := PostData{
		"id":    fullname,
		"state": state,
	}
	resp, err := c.client.Post(API_PATH["sendreplies"], postdata)

	if err != nil {
		log.Fatal("Error in sendreplies")
	}

	defer resp.Body.Close()
}
