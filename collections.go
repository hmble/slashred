package slashred

import (
	"fmt"
	"log"
)

// This is of scope modposts
// don't know what that is.
//
// TODO: Need to implement this later
//
// Scope: modposts
// API https://www.reddit.com/dev/api/#POST_api_v1_collections_add_post_to_collection

type CollectionService service

func (c *CollectionService) Collection() {
	path := "/api/v1/collections/collection"
	resp, err := c.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting collection response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/collection.json")
}

func (c *CollectionService) SubredditCollection(subreddit string) {
	path := "/api/v1/collections/subreddit_collections"

	fmt.Println(SubredditPrefix + subreddit)
	resp, err := c.client.Get(path, Option{"sr_fullname": SubredditPrefix + subreddit})

	if err != nil {
		log.Fatal("Error in getting collections response")
	}

	defer resp.Body.Close()
	SaveResponse(resp.Body, "test_data/collection.json")
}
