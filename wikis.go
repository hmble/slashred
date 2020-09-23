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
