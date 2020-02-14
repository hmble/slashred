package slashred

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

const (
	maxDepth int = 20
)

type CommentService service
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
	//	ReplyArray []*Comment
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

	// replyArray := make([]*Comment, 0)
	// for _, reply := range data.Children {

	// 	replyArray = append(replyArray, reply.Comment)

	// }

	//r.ReplyArray = replyArray
	return nil

}

// Methods
func (c *CommentService) GetComments(path, sort string) []CommentListing {
	u, pathErr := url.Parse(path)

	if pathErr != nil {
		panic(pathErr)
	}

	pathArray := strings.Split(u.Path, "/")
	subreddit := pathArray[2]
	article := pathArray[4]

	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	opt := Option{
		"limit":   "100",
		"context": "100",
		"sort":    sort,
		"depth":   "8",
	}
	resp, err := c.client.Get(endpoint, opt)

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	PrintHeader(resp)

	c.client.savelimit(resp)
	type listComment struct {
		Kind string      `json:"kind"`
		Data CommentData `json:"data"`
	}

	// https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	result := make([]listComment, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	if er != nil {
		panic(er)
	}

	commentListing := result[1].Data.Children

	return commentListing
}

// Comment.GetCommentsId
func (c *CommentService) GetCommentsId(path, comment, sort string, depth int) []*Comment {
	u, pathErr := url.Parse(path)

	if pathErr != nil {
		panic(pathErr)
	}

	pathArray := strings.Split(u.Path, "/")

	subreddit := pathArray[2]
	article := pathArray[4]

	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	options := Option{
		"comment": comment,
		"depth":   "8",
		"limit":   "100",
		"sort":    sort,
		"context": "100",
	}
	resp, err := c.client.Get(endpoint, options)

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	//SaveResponse(resp.Body, "test_data/commeny_by_id.json")
	//PrintHeader(resp)

	c.client.savelimit(resp)
	type ListSub struct {
		Kind string      `json:"kind"`
		Data CommentData `json:"data"`
	}

	result := make([]ListSub, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	if er != nil {
		panic(er)
	}

	commentArray := result[1].Data.Children

	comments := make([]*Comment, 0)
	for _, comment := range commentArray {

		if comment.Comment != nil {
			comments = append(comments, comment.Comment)

			replies := c.getAllReplies(depth, comment.Comment, sort, path)

			comments = append(comments, replies...)
		}
	}

	return comments

}
func (c *CommentService) PostComment(thingId, text string) {
	c.client.edit(API_PATH["comment"], thingId, text)
}

// Test Remained
// state is bool
// Comment.SendReplies
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

// TODO(hmble): Need to add this More method to Comments and More Response

// Comment.ReplaceMore
func (c *CommentService) ReplaceMore(more *More,
	linkId, sort, path string) []*Comment {
	tempdata := PostData{}

	tempdata["children"] = strings.Join(more.Children, ",")
	tempdata["link_id"] = linkId
	// false option gives more object in response after ReplaceMore method
	// true option limits that more object on parent node
	tempdata["limit_children"] = "false"
	tempdata["depth"] = "8"
	tempdata["sort"] = sort

	resp, err := c.client.Post(API_PATH["morechildren"], tempdata)

	c.client.savelimit(resp)
	if err != nil {
		log.Fatal("Error in getting more response")

	}
	defer resp.Body.Close()

	//SaveResponse(resp.Body, "test_data/MemesResponse.json")
	type moreReplies struct {
		Json struct {
			Data struct {
				Things []CommentListing `json:"things"`
			} `json:"data"`
		} `json:"json"`
	}

	var response moreReplies
	moreErr := json.NewDecoder(resp.Body).Decode(&response)

	if moreErr != nil {
		log.Fatal("Error in getting more replies response")
	}

	comments := response.Json.Data.Things

	commentsArray := make([]*Comment, 0)
	for _, comment := range comments {

		if comment.Comment != nil {
			commentsArray = append(commentsArray, comment.Comment)

		} else {

			if comment.More.Count != 0 {
				moreArray := c.ReplaceMore(comment.More, linkId, sort, path)
				commentsArray = append(commentsArray, moreArray...)
			} else {
				// count 0 case will go here
				continueParent := strings.Split(comment.More.ParentID, "_")

				contComments := c.GetCommentsId(path, continueParent[1], sort, 8)

				commentsArray = append(commentsArray, contComments...)
			}

		}
	}

	return commentsArray

}

// Comment.List
func (c *CommentService) List(list []CommentListing, depth int, sort, path string, fetchMore bool) []*Comment {
	comments := make([]*Comment, 0)

	if depth > maxDepth {
		log.Fatal("Depth should be less than 8")
	}

	var linkId string
	for _, item := range list {

		if item.Comment != nil {
			linkId = item.Comment.LinkID
			comments = append(comments, item.Comment)

			temp := c.getAllReplies(depth-1, item.Comment, sort, path)

			comments = append(comments, temp...)
		}
		if fetchMore {
			if item.More != nil {

				moreComment := c.ReplaceMore(item.More, linkId, sort, path)

				comments = append(comments, moreComment...)

			}
		}
	}

	return comments

}

func (c *CommentService) getAllReplies(depth int, comment *Comment, sort, path string) []*Comment {
	moreReplies := make([]*Comment, 0)

	replies := comment.Replies.Data.Children
	if depth >= 0 {
		for _, reply := range replies {

			if reply.Comment != nil {
				moreReplies = append(moreReplies, reply.Comment)
				tempReply := c.getAllReplies(depth-1, reply.Comment, sort, path)

				moreReplies = append(moreReplies, tempReply...)
			} else {
				fetchedMore := c.ReplaceMore(reply.More, comment.LinkID, sort, path)

				for _, item := range fetchedMore {
					moreReplies = append(moreReplies, item)

					itemReplies := c.getAllReplies(depth-1, item, sort, path)

					moreReplies = append(moreReplies, itemReplies...)
				}
			}
		}

	}

	return moreReplies
}

func (c *CommentService) Replies(depth int, comment *Comment, sort, path string) []*Comment {

	return c.getAllReplies(depth, comment, sort, path)
}
