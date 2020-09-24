package slashred

import (
	"encoding/json"
	"fmt"
	"log"
)

type MultisService service

type Multi struct {
	CanEdit        bool      `json:"can_edit"`
	Name           string    `json:"name"`
	NumSubscribers int       `json:"num_subscribers"`
	Subreddits     []Multisr `json:"subreddits"`
	Path           string    `json:"path"`
	Owner          string    `json:"owner"`
	DescriptionMd  string    `json:"description_md"`
	IsFavorited    bool      `json:"is_favorited"`
}

type Multisr struct {
	Name string
}

// Get subreddit subscriber multis of subreddit
func (m *MultisService) Mine() []Multi {

	path := fmt.Sprintf("/api/multi/mine")

	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	type labeledMulti struct {
		Kind string
		Data Multi
	}

	result := make([]labeledMulti, 0)

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Error in decoding json response in MultisService")
	}

	var multis []Multi
	for _, item := range result {
		multis = append(multis, item.Data)
	}

	return multis
}

// Delete subreddit multis
// multipath: path variable of multis got from /api/multi/mine Endpoint
// subreddit: subreddit name

func (m *MultisService) DeleteSr(multipath string, subreddit string) {

	path := fmt.Sprintf("/api/multi/%s/r/%s", multipath, subreddit)

	opts := Option{
		"multipath": multipath,
		"srname":    subreddit,
	}

	resp, err := m.client.Delete(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (m *MultisService) AddSr(multipath string, subreddit string) {

	path := fmt.Sprintf("/api/multi/%s/r/%s", multipath, subreddit)

	model := fmt.Sprintf(`{"name": "%s"}`, subreddit)
	resp, err := m.client.Put(path, PostData{
		"multipath": multipath,
		"srname":    subreddit,
		"model":     model,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	printBytes(resp.Body, m.client)
}

func (m *MultisService) UpdateDescription(multipath, description string) {
	path := fmt.Sprintf("/api/multi/%s/description", multipath)

	model := fmt.Sprintf(`{"body_md": "%s"}`, description)
	resp, err := m.client.Put(path, PostData{
		"multipath": multipath,
		"model":     model,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	printBytes(resp.Body, m.client)
}

// Get multi description
func (m *MultisService) GetDescription(multipath string) {
	path := fmt.Sprintf("/api/multi/%s/description", multipath)

	resp, err := m.client.Get(path, Option{
		"multipath": multipath,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)

}

// Create multi.
// TODO(hmble): Add more parameters to create data string ?
// TODO(hmble): Add doc for json data string.
func (m *MultisService) CreateMulti(user, name string, data string) {
	// api/multi/user/{user}/m/{multi}/
	path := fmt.Sprintf("/api/multi/user/%s/m/%s", user, name)

	postdata := PostData{
		"model":     data,
		"multipath": fmt.Sprintf("/user/%s/m/%s", user, name),
	}
	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)

}

// Update a multi

func (m *MultisService) UpdateMulti(user, name string, data string) {
	// api/multi/user/{user}/m/{multi}/
	path := fmt.Sprintf("/api/multi/user/%s/m/%s", user, name)

	postdata := PostData{
		"model":     data,
		"multipath": fmt.Sprintf("/user/%s/m/%s", user, name),
	}
	resp, err := m.client.Put(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)

}

// Fetch a list of public multis belonging to a username

func (m *MultisService) FetchPublicMultis(username string) {
	path := fmt.Sprintf("/api/multi/user/%s", username)

	resp, err := m.client.Get(path, Option{
		"username": username,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Delete a multi

// TODO(hmble): Accept only multi name ?
// Update this todo once we add global user info to client. Also do the same for
// other methods in this file.
// multipath is /user/{user}/m/{multi}/
func (m *MultisService) DeleteMulti(multipath string) {
	path := fmt.Sprintf("/api/multi/%s", multipath)

	resp, err := m.client.Delete(path, Option{
		"multipath": multipath,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Fetch a multis data and subreddit's name

func (m *MultisService) GetSingleMultiData(multipath string) {

	path := fmt.Sprintf("/api/multi/%s", multipath)

	resp, err := m.client.Get(path, Option{
		"multipath": multipath,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Get data about subreddit in a multi

func (m *MultisService) GetSrData(multipath, srname string) {

	path := fmt.Sprintf("/api/multi/%s/r/srname", multipath)

	resp, err := m.client.Get(path, Option{
		"multipath": multipath,
		"srname":    srname,
	})

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Copy a multi
func (m *MultisService) CopyMulti(from, to, description, displayName string) {
	path := "/api/multi/copy"

	postdata := PostData{
		"from":           from,
		"to":             to,
		"description_md": description,
		"display_name":   displayName,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}
