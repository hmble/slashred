package slashred

import (
	"fmt"
)

type GoldService service

// TODO(hmble): This endpoint is not tested as it requires premium subscription
//
// Key              Value
// ==============================
// fullname	    fullname of a thing

func (g *GoldService) Gild(fullname string) {
	path := fmt.Sprintf("/api/v1/gold/gild/%s", fullname)

	postdata := PostData{
		"fullname": fullname,
	}

	resp, err := g.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

}

// TODO(hmble): This endpoint is not tested as it requires premium subscription

// Key                 Value
// =============================================
// months      an integer between 1 and 36
// username    A valid, existing reddit username

func (g *GoldService) Give(username, months string) {
	path := fmt.Sprintf("/api/v1/gold/gild/%s", username)

	postdata := PostData{
		"username": username,
		"months":   months,
	}

	resp, err := g.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

}
