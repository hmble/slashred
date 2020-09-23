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

// Retrieve the current permission settings for page
func (w *WikiService) GetPageSettings(subreddit, page string) {
	path := fmt.Sprintf("/r/%s/wiki/settings/%s", subreddit, page)

	opts := Option{
		"page": page,
	}
	resp, err := w.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

//Update the permissions and visibility of wiki page
//
// listed	boolean value
// page	the name of an existing wiki page
// permlevel	an integer

func (w *WikiService) UpdatePageSettings(subreddit, postdata PostData) {
	path := fmt.Sprintf("/r/%s/api/wiki/hide", subreddit)

	resp, err := w.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Allow/deny username to edit this wiki page
func (w *WikiService) AllowEditor(subreddit, act, username, page string) {
	path := fmt.Sprintf("/r/%s/api/wiki/alloweditor/%s", subreddit, act)

	postdata := PostData{
		"act":      act,
		"page":     page,
		"username": username,
	}
	resp, err := w.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Edit a wiki page
//
// content  page content
// page	    the name of an existing page or a new page to create

// previous	the starting point revision for this edit
// reason	a string up to 256 characters long, consisting of printable characters.
func (w *WikiService) WikiEdit(subreddit, postdata PostData) {
	path := fmt.Sprintf("/r/%s/api/wiki/edit", subreddit)

	resp, err := w.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

// Toggle the public visibility of a wiki page revision
//
// page	    the name of an existing wiki page
// revision	a wiki revision ID
func (w *WikiService) WikiHide(subreddit, postdata PostData) {
	path := fmt.Sprintf("/r/%s/api/wiki/hide", subreddit)

	resp, err := w.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}

//Revert a wiki page to revision
//
// page	    the name of an existing wiki page
// revision	a wiki revision ID
func (w *WikiService) WikiRevert(subreddit, postdata PostData) {
	path := fmt.Sprintf("/r/%s/api/wiki/hide", subreddit)

	resp, err := w.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body)
}
