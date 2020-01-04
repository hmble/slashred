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
	Author  string  `json:"author"`
	Score   int     `json:"score"`
	Body    string  `json:"body"`
	Depth   int     `json:"depth"`
	Url     string  `json:"permalink"`
	Replies Replies `json:"replies"`
	//	Replies Replies `json:"replies"`
}

type Replies struct {
	Data ReplyData `json:"data"`
}

type ReplyData struct {
	Children []RepliesArray
}

type RepliesArray struct {
	Kind string
	Data Comment
}

// func (r *RepliesArray) UnmarshalJSON(b []byte) error {

// 	var tmp map[string]json.RawMessage

// 	if err := json.Unmarshal(b, &tmp); err != nil {
// 		return err
// 	}

// 	var data interface{}
// 	if r.Kind == "t1" {
// 		data = Comment{}
// 	} else {
// 		data = More{}
// 	}

// 	if err := json.Unmarshal(b, &data); err != nil {
// 		return err
// 	}

// 	r.Data = data
// 	return nil

// }

type More struct {
	Count    int
	Name     string
	ParentID string
	ID       string
	Depth    int
	Children []string
}

func (r *Replies) UnmarshalJSON(b []byte) error {
	if string(b) == "\"\"" {
		return nil
	}

	var tmp map[string]json.RawMessage

	if err := json.Unmarshal(b, &tmp); err != nil {
		return errors.New("Error in unmarshaling raw message")
	}

	//for k, _ := range tmp {
	//	//	fmt.Printf("%s : %s", k, string(v))
	//	fmt.Println(k)
	//}

	var data ReplyData

	if err := json.Unmarshal(tmp["data"], &data); err != nil {
		return errors.New("Data point does not exists")

	}

	r.Data = data

	return nil

}

//func (r *Replies2) UnmarshalJSON(b []byte) error {
//	if string(b) == "\"\"" {
//		return nil
//	}

//	r.Replies = &Replies{}
//	//	var temp Replies
//	if err := json.Unmarshal(b, r.Replies); err != nil {
//		return errors.New("error in unmarshaling replies")
//	}
//	return nil
//}

type CommentListing struct {
	Kind string
	Data Comment `json:"data"`
}
type Data struct {
	Children []CommentListing
}
