package slashred

import (
	"fmt"
	"log"
)

type GoldService service

// Post Methods

// This endpoint is not tested as it requires premium subscription
// Gold.Gild
func (g *GoldService) Gild(fullname string) {
	path := fmt.Sprintf("/api/v1/gold/gild/%s", fullname)

	postdata := PostData{
		"fullname": fullname,
	}

	resp, err := g.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting resonse from %s : %v", path, err)
	}

	defer resp.Body.Close()

}

// This endpoint is not tested as it requires premium subscription
// Gold.Give
func (g *GoldService) Give(username, months string) {
	path := fmt.Sprintf("/api/v1/gold/gild/%s", username)

	postdata := PostData{
		"username": username,
		"months":   months,
	}

	resp, err := g.client.Post(path, postdata)

	if err != nil {
		log.Fatalf("Error in getting resonse from %s : %v", path, err)
	}

	defer resp.Body.Close()

}
