package slashred

import "fmt"

type WikiService service

// Retrieve a list of wiki pages in this subreddit
func (w *WikiService) GetWikiPages(subreddit string) {
	path := fmt.Sprintf("/r/%s/wiki/pages", subreddit)

	resp, err := w.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Retrieve a list of discussions about this wiki page
// TODO(hmble): add Option parameter ?
func (w *WikiService) GetPageDiscussion(subreddit, page string) {
	path := fmt.Sprintf("/r/%s/wiki/discussions/%s", subreddit, page)

	resp, err := w.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Retrieve a list of wiki pages in this subreddit
func (w *WikiService) GetRevision(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/wiki/revisions", subreddit)

	resp, err := w.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

//Retrieve a list of revisions of this wiki page
// TODO(hmble): add Option parameter ?
func (w *WikiService) GetPageRevision(subreddit, page string) {
	path := fmt.Sprintf("/r/%s/wiki/discussions/%s", subreddit, page)

	resp, err := w.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Return the content of a wiki page

// If v is given, show the wiki page as it was at that version If both v and v2 are given, show a diff of the two
func (w *WikiService) GetWikiContent(subreddit, page, v, v2 string) {
	path := fmt.Sprintf("/r/%s/wiki/%s", subreddit, page)

	opts := Option{
		"v":    v,
		"v2":   v2,
		"page": page,
	}
	resp, err := w.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}
