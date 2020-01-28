package slashred

import "log"

// This is of scope modposts
// don't know what that is.
//
// TODO: Need to implement this later
//
// Scope: modposts
// API https://www.reddit.com/dev/api/#POST_api_v1_collections_add_post_to_collection

type CollectionService service

func (c *CollectionService) Collection() {
	resp, err := c.client.Get(API_PATH["collection"], NoOptions)

	if err != nil {
		log.Fatal("Error in getting collection response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/collection.json")
}
