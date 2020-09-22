package slashred

import (
	"fmt"
)

type MultisService service

// Get subreddit subscriber multis of subreddit
func (m *MultisService) Mine() {

	path := fmt.Sprintf("/api/multi/mine")

	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	// TODO(hmble): Make struct field of multis to parse
	printBytes(resp.Body)

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
