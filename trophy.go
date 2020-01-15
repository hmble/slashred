package slashred

import (
	"encoding/json"
	"log"
)

type Trophy struct {
	AwardID     string      `json:"award_id"`
	Description interface{} `json:"description"`
	ID          string      `json:"id"`
	Icon40      string      `json:"icon_40"`
	Icon70      string      `json:"icon_70"`
	Name        string      `json:"name"`
	URL         interface{} `json:"url"`
}

// type Listing struct {
// 	Data interface{} `json:"data"`
// 	Kind string      `json:"kind"`
// }

func (a *AccountService) Trophies() ([]Trophy, error) {
	resp, err := a.client.Get(API_PATH["trophies"], NoOptions)

	if err != nil {
		log.Fatal("Error in getting Trophies response")
	}

	defer resp.Body.Close()
	PrintHeader(resp)

	//SaveResponse(resp.Body, "trophies.json")

	var trophyListing struct {
		Data struct {
			Trophies []struct {
				Trophy Trophy `json:"data"`
				Kind   string `json:"kind"`
			} `json:"trophies"`
		} `json:"data"`
		Kind string `json:"kind"`
	}

	trophyErr := json.NewDecoder(resp.Body).Decode(&trophyListing)

	if trophyErr != nil {
		log.Fatal("Error in decoding trophies response")
		return nil, trophyErr
	}

	var trophies []Trophy

	for _, data := range trophyListing.Data.Trophies {
		trophies = append(trophies, data.Trophy)

	}

	return trophies, nil

}
