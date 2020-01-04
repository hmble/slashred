package slashred

import (
	"encoding/json"
	"errors"
)

type Submission struct {
	SelfText    string `json:"selftext"`
	Score       int    `json:"score"`
	Author      string `json:"author"`
	NumComments int    `json:"num_comments"`
	Url         string `json:"url"`
}

type Comment struct {
	Author string `json:"author"`
	// Score   int     `json:"score"`
	// Body    string  `json:"body"`
	// Depth   int     `json:"depth"`
	// Url     string  `json:"permalink"`
	Replies ReplyData `json:"replies"`
	//	Replies Replies `json:"replies"`
}

type RepliesListing struct {
	Kind string
	Data Reply `json:"data"`
}
type _Comment Comment
type Reply struct {
	Comment
}

type ReplyData struct {
	Dist string
	//	Children string
	Children []RepliesListing
}

// type Replies struct {
// 	*ReplyData
// }

func (r *ReplyData) UnmarshalJSON(b []byte) error {
	if string(b) == "\"\"" {
		return nil
	}

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return errors.New("Error in unmarshaling raw message")
	}

	var children []RepliesListing

	if err := json.Unmarshal(tmp["children"], &children); err != nil {
		return err
	}

	r.Children = children
	return nil

}

//func (r *Replies) UnmarshalJSON(b []byte) error {
//	if string(b) == "\"\"" {
//		return nil
//	}

//	r.ReplyData = &ReplyData{}
//	//	var temp Replies
//	if err := json.Unmarshal(b, r.ReplyData); err != nil {
//		return errors.New("error in unmarshaling replies")
//	}
//	fmt.Println(len(r.Children))
//	return nil
//}

type CommentListing struct {
	Kind string
	Data Comment `json:"data"`
}
type Data struct {
	Children []CommentListing
}
