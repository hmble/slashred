package slashred

import (
	"encoding/json"
	"errors"
)

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
