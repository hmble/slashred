package slashred

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type FlairService service

type Flair struct {
	CssClass     string `json:"css_class"`
	ID           string `json:"id"`
	Text         string `json:"text"`
	TextEditable bool   `json:"text_editable"`
	Type         string `json:"type"`
	//TODO(hmble) : RichText field is not implemented?
}

type Flairv2 struct {
	Flair
	AllowableContent string `json:"allowable_content"`
	BackgroundColor  string `json:"background_color"`
	MaxEmojis        int    `json:"max_emojis"`
	TextColor        string `json:"text_color"`
}

// Get flair list
// Flair.List
func (f *FlairService) List(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/api/flairlist", subreddit)

	resp, err := f.client.Get(path, opts)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/flairlist.json")
}

// Flair.LinkFlair
func (f *FlairService) LinkFlair(subreddit string) []*Flair {
	path := fmt.Sprintf("/r/%s/api/link_flair", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

	flairs := make([]*Flair, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs
}

// Flair.LinkFlairV2
func (f *FlairService) LinkFlairV2(subreddit string) []*Flairv2 {
	path := fmt.Sprintf("/r/%s/api/link_flair_v2", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

	flairs := make([]*Flairv2, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs
}

func (f *FlairService) UserFlair(subreddit string) []*Flair {
	path := fmt.Sprintf("/r/%s/api/user_flair", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
	flairs := make([]*Flair, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs

}

// Flair.UserFlairV2
func (f *FlairService) UserFlairV2(subreddit string) []*Flairv2 {
	path := fmt.Sprintf("/r/%s/api/user_flair_v2", subreddit)

	resp, err := f.client.Get(path, NoOptions)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
	flairs := make([]*Flairv2, 0)

	if err := json.NewDecoder(resp.Body).Decode(&flairs); err != nil {
		log.Fatal("Error in decoding json response")
	}

	return flairs

}

// Post method

func (f *FlairService) ClearFlairTemplate(subreddit string, flairtype string) {
	postdata := PostData{

		"flair_type": flairtype,
	}

	path := fmt.Sprintf("/r/%s/api/clearflairtemplates", subreddit)
	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/clearflairtemplates.json")
}

func (f *FlairService) DeleteFlair(subreddit string, name string) {
	postdata := PostData{
		"name": name,
	}
	path := fmt.Sprintf("/r/%s/api/deleteflair", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
}

// Flair.DeleteFlairTemplate
func (f *FlairService) DeleteFlairTemplate(subreddit string, id string) {
	postdata := PostData{
		"flair_template_id": id,
	}
	path := fmt.Sprintf("/r/%s/api/deleteflairtemplate", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
}

// Flair.ApiFlair

func (f *FlairService) ApiFlair(subreddit string, postdata PostData) {

	path := fmt.Sprintf("/r/%s/api/deleteflairtemplate", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
}

// Flair.FlairConfig
func (f *FlairService) FlairConfig(subreddit string, postdata PostData) {

	path := fmt.Sprintf("/r/%s/api/flairconfig", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
}

// Flair.FlairCsv
func (f *FlairService) FlairCsv(subreddit, flaircsv string) {

	path := fmt.Sprintf("/r/%s/api/flaircsv", subreddit)

	postdata := PostData{
		"flair_csv": flaircsv,
	}
	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()
}

// Flair.FlairTemplate
func (f *FlairService) FlairTemplate(subreddit string) {
	postdata := PostData{
		"flair_template_id": uuid.New().String(),
		"flair_type":        "USER_FLAIR",
		"text":              "TestFlairByMod",
		"text_editable":     "false",
	}

	path := fmt.Sprintf("/r/%s/api/flairtemplate", subreddit)
	resp, err := f.client.Post(path, postdata)
	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/flairtemplate.json")

}

func (f *FlairService) FlairTemplateV2(subreddit, postdata PostData) {
	path := fmt.Sprintf("/r/%s/api/flairtemplate_v2", subreddit)
	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}

	defer resp.Body.Close()

}

// Flair.FlairSelector
func (f *FlairService) FlairSelector(subreddit, link, name string) {
	postdata := PostData{
		"link": link,
		"name": name,
	}

	path := fmt.Sprintf("/r/%s/flairselector", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}
	defer resp.Body.Close()

}

// Flair.SelectFlair
func (f *FlairService) SelectFlair(subreddit string, postdata PostData) {

	path := fmt.Sprintf("/r/%s/selectflair", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}
	defer resp.Body.Close()

}

// Flair.FlairEnabled
func (f *FlairService) FlairEnabled(subreddit, flairEnabled string) {

	postdata := PostData{
		"flair_enabled": flairEnabled,
	}

	path := fmt.Sprintf("/r/%s/selectflair", subreddit)

	resp, err := f.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting response for %s : %v", path, err)
	}
	defer resp.Body.Close()

}
