package slashred

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
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

	//r.Kind = string(tmp["kind"])

	// var comment Comment
	// var more More

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
	ID     string `json:"id"`
	Parent string `json:"parent_id"`

	Replies Replies `json:"replies"`
	//	Replies Replies `json:"replies"`
}

type Replies struct {
	Data       ReplyData `json:"data"`
	ReplyArray []*Comment
}

type ReplyData struct {
	Children []RepliesArray
}
type More struct {
	Count    int `json:"count"`
	Name     string
	ParentID string `json:"parent_id"`
	ID       string `json:"id"`
	Depth    int
	Children []string
}
type RepliesArray struct {
	Kind    string
	Data    json.RawMessage `json:"data"`
	Comment *Comment
	More    *More
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

	replyArray := make([]*Comment, 0)
	for _, reply := range data.Children {

		replyArray = append(replyArray, reply.Comment)

	}

	r.ReplyArray = replyArray
	return nil

}

// Methods
func (c *CommentService) GetComments(subreddit, article, sort string) []CommentListing {
	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	opt := Option{
		"limit":   "100",
		"context": "100",
		"sort":    sort,
		"depth":   "8",
	}
	//	resp, err := c.client.Get(endpoint, NoOptions)
	resp, err := c.client.Get(endpoint, opt)

	if err != nil {
		log.Fatal("Error in getting comments response")
	}
	defer resp.Body.Close()

	//fmt.Println("-------Got Reponse of comments------------")
	//SaveResponse(resp.Body, "test_data/askreddit.json")
	PrintHeader(resp)

	type listComment struct {
		Kind string      `json:"kind"`
		Data CommentData `json:"data"`
	}

	// https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	result := make([]listComment, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	//	fmt.Println("----got here after decode")
	if er != nil {
		panic(er)
		//log.Fatal("Error in decoding comments")
	}

	commentListing := result[1].Data.Children

	//comments := make([]Comment, 0)
	//more := make([]More, 0)

	//for _, comment := range commentListing {
	//	comments = append(comments, comment.Comment)
	//	more = append(more, comment.More)
	//}

	////commentReply := make([]Comment, 0)

	//return comments, more

	return commentListing
}

func (c *CommentService) GetCommentsID(subreddit, article, comment, sort string) {
	endpoint := fmt.Sprintf("/r/%s/comments/%s", subreddit, article)

	// q.Add("comment", comment)
	// q.Add("depth", "7")

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
	PrintHeader(resp)

	type ListSub struct {
		Kind string      `json:"kind"`
		Data CommentData `json:"data"`
	}

	result := make([]ListSub, 0)
	er := json.NewDecoder(resp.Body).Decode(&result)

	if er != nil {
		panic(er)
		//log.Fatal("Error in decoding comments")
	}

	comments := result[1].Data.Children

	for _, comment := range comments {
		fmt.Printf("Author : %s ID[%s]\n", comment.Comment.Author, comment.Comment.ID)

		replies := comment.Comment.Replies.Data.Children

		for _, reply := range replies {

			fmt.Printf("\t Author : %s Parent[%s]\n", reply.Comment.Author, reply.Comment.Parent)
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

// TODO(hmble): Need to add this More method to Comments and More Response

func (c *CommentService) ReplaceMore(more *More,
	linkId, sort, parent string) []*Comment {
	tempdata := PostData{}

	//for k, v := range postdata {
	//	tempdata[k] = v
	//}
	////tempdata["children"] = str
	tempdata["children"] = strings.Join(more.Children, ",")
	tempdata["link_id"] = linkId
	tempdata["limit_children"] = "false"
	tempdata["depth"] = "8"
	tempdata["sort"] = sort

	resp, err := c.client.Post(API_PATH["morechildren"], tempdata)

	if err != nil {
		log.Fatal("Error in getting more response")

	}
	defer resp.Body.Close()

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

	parent = CommentPrefix + parent
	commentsArray := make([]*Comment, 0)
	for _, comment := range comments {
		if comment.Comment != nil {
			if comment.Comment.Parent == parent {
				commentsArray = append(commentsArray, comment.Comment)
			}
		}
	}

	return commentsArray

}

func (c *CommentService) List(list []CommentListing, sort string) []*Comment {
	comments := make([]*Comment, 0)

	for _, item := range list {
		if item.Comment != nil {
			comments = append(comments, item.Comment)

			temp := c.getAllReplies(8, item.Comment, sort)

			comments = append(comments, temp...)
		}
	}

	return comments
}

func (c *CommentService) getAllReplies(depth int, comment *Comment, sort string) []*Comment {
	moreReplies := make([]*Comment, 0)

	replies := comment.Replies.Data.Children
	if depth >= 0 {
		for _, reply := range replies {

			if reply.Comment != nil {
				moreReplies = append(moreReplies, reply.Comment)
				tempReply := c.getAllReplies(depth-1, reply.Comment, sort)

				moreReplies = append(moreReplies, tempReply...)
			} else {
				fetchedMore := c.ReplaceMore(reply.More, comment.LinkID, sort, comment.ID)

				for _, item := range fetchedMore {
					moreReplies = append(moreReplies, item)

					itemReplies := c.getAllReplies(depth-1, item, sort)

					moreReplies = append(moreReplies, itemReplies...)
				}
			}
		}

	}

	return moreReplies
}
