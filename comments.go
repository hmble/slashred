package slashred

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	maxDepth int = 20
)

type CommentService struct {
	client *Client
	// To reuse path varaible accross vaious methods like continueThread and
	// GetComments
	path string
}
type CommentListing struct {
	Kind    string
	Data    json.RawMessage `json:"data"`
	Comment *Comment
	More    *More
}
type CommentData struct {
	Children []CommentListing
}

func (c *CommentListing) UnmarshalJSON(b []byte) error {

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	kind := string(tmp["kind"])

	c.Kind = kind[1 : len(kind)-1]

	if string(tmp["kind"]) == "\"t1\"" {
		var comment Comment
		if err := json.Unmarshal(tmp["data"], &comment); err != nil {
			return err
		}
		c.Comment = &comment
		c.More = nil

	} else {
		var more More
		if err := json.Unmarshal(tmp["data"], &more); err != nil {
			return err
		}
		c.More = &more
		c.Comment = nil
	}

	return nil
}

type Comment struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	Body   string `json:"body"`
	Depth  int    `json:"depth"`
	Url    string `json:"permalink"`
	LinkID string `json:"link_id"`
	Id     string `json:"id"`
	Parent string `json:"parent_id"`

	Replies Replies `json:"replies"`
}

type Replies struct {
	Data ReplyData `json:"data"`
}

type ReplyData struct {
	Children []RepliesArray
}

type RepliesArray struct {
	Kind    string
	Data    json.RawMessage `json:"data"`
	Comment *Comment
	More    *More
}

type More struct {
	Count    int `json:"count"`
	Name     string
	ParentID string `json:"parent_id"`
	ID       string `json:"id"`
	Depth    int
	Children []string
}

func (r *RepliesArray) UnmarshalJSON(b []byte) error {

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	kind := string(tmp["kind"])

	r.Kind = kind[1 : len(kind)-1]

	if string(tmp["kind"]) == "\"t1\"" {
		var comment Comment
		if err := json.Unmarshal(tmp["data"], &comment); err != nil {
			return err
		}
		r.Comment = &comment
		r.More = nil

	} else {
		var more More
		if err := json.Unmarshal(tmp["data"], &more); err != nil {
			return err
		}
		r.More = &more
		r.Comment = nil
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

// Get comments for a subreddit
func (c *CommentService) GetComments(path, commentId string) []*Comment {
	// [/r/subreddit]/comments/article
	c.path = path
	split := strings.Split(path[23:], "/")
	endpoint := fmt.Sprintf("/%s", strings.Join(split[0:4], "/"))

	options := Option{
		"depth": "8",
		"limit": "100",
		"sort":  "best",
		// Reusing get comments to get continue thread
		"comment": commentId,
	}

	resp, err := c.client.Get(endpoint, options)

	if err != nil {
		respError(endpoint)
	}

	defer resp.Body.Close()

	PrintHeader(resp)

	type listComment struct {
		Kind string      `json:"kind"`
		Data CommentData `json:"data"`
	}
	// https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	result := make([]listComment, 0)

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Error in decoding json response")
	}

	comments := make([]*Comment, 0)
	commentListing := result[1].Data.Children
	for _, item := range commentListing {

		if item.Comment != nil {
			comments = append(comments, item.Comment)

			comments = append(comments, c.getReplies(item.Comment, 8)...)

		} else {

			comments = append(comments, c.getMore(item.More, split[3])...)

		}

	}

	return comments

}

func (c *CommentService) getMore(more *More, linkId string) []*Comment {

	comments := make([]*Comment, 0)
	endpoint := "/api/morechildren"

	postdata := PostData{
		"children":       strings.Join(more.Children, ","),
		"depth":          "100",
		"limit_children": "false",
		"link_id":        linkId,
	}

	resp, err := c.client.Post(endpoint, postdata)

	if err != nil {
		respError(endpoint)
	}
	defer resp.Body.Close()

	var moreReplies struct {
		Json struct {
			Data struct {
				Things []CommentListing `json:"things"`
			} `json:"data"`
		} `json:"json"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&moreReplies); err != nil {
		log.Fatal("Error in getting more response")
	}

	commentsArr := moreReplies.Json.Data.Things

	for _, comment := range commentsArr {
		if comment.Comment != nil {
			comments = append(comments, comment.Comment)
		} else {
			if comment.More.Count != 0 {
				comments = append(comments, c.getMore(comment.More, linkId)...)

			} else {
				// To get continue thread we need to get comment by id.
				comments = append(comments, c.continueThread(comment.More.ID)...)
			}
		}

	}

	return comments
}

// To use with continue thread.
func (c *CommentService) continueThread(comment string) []*Comment {

	return c.GetComments(c.path, comment)

}

// here comment is Parent whose replies we are getting
func (c *CommentService) getReplies(comment *Comment, depth int) []*Comment {

	comments := make([]*Comment, 0)
	replies := comment.Replies.Data.Children

	if depth >= 0 {
		for _, reply := range replies {

			if reply.Comment != nil {
				comments = append(comments, reply.Comment)

				comments = append(comments, c.getReplies(reply.Comment, depth-1)...)

			} else {
				// getmore comments

				comments = append(comments, c.getMore(reply.More, comment.LinkID)...)
				//c.getMore(reply.More, comment.LinkID)
				//fmt.Println("from nested reply ", reply.More.Count)
			}
		}
	}

	return comments

}
