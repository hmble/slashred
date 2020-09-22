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
	printBytes(resp.Body)
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
	printBytes(resp.Body)
}
