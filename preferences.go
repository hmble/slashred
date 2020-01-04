package slashred

import (
	"encoding/json"
	"log"
)

type Preferences struct {
	AcceptPms                             string      `json:"accept_pms"`
	ActivityRelevantAds                   bool        `json:"activity_relevant_ads"`
	AllowClicktracking                    bool        `json:"allow_clicktracking"`
	Beta                                  bool        `json:"beta"`
	Clickgadget                           bool        `json:"clickgadget"`
	CollapseLeftBar                       bool        `json:"collapse_left_bar"`
	CollapseReadMessages                  bool        `json:"collapse_read_messages"`
	Compress                              bool        `json:"compress"`
	DefaultCommentSort                    string      `json:"default_comment_sort"`
	DefaultThemeSr                        interface{} `json:"default_theme_sr"`
	DesignBeta                            bool        `json:"design_beta"`
	DomainDetails                         bool        `json:"domain_details"`
	EmailDigests                          bool        `json:"email_digests"`
	EmailMessages                         bool        `json:"email_messages"`
	EmailUnsubscribeAll                   bool        `json:"email_unsubscribe_all"`
	EnableDefaultThemes                   bool        `json:"enable_default_themes"`
	Geopopular                            string      `json:"geopopular"`
	HideAds                               bool        `json:"hide_ads"`
	HideDowns                             bool        `json:"hide_downs"`
	HideFromRobots                        bool        `json:"hide_from_robots"`
	HideUps                               bool        `json:"hide_ups"`
	HighlightControversial                bool        `json:"highlight_controversial"`
	HighlightNewComments                  bool        `json:"highlight_new_comments"`
	IgnoreSuggestedSort                   bool        `json:"ignore_suggested_sort"`
	LabelNsfw                             bool        `json:"label_nsfw"`
	Lang                                  string      `json:"lang"`
	Layout                                int         `json:"layout"`
	LegacySearch                          bool        `json:"legacy_search"`
	LiveOrangereds                        bool        `json:"live_orangereds"`
	MarkMessagesRead                      bool        `json:"mark_messages_read"`
	Media                                 string      `json:"media"`
	MediaPreview                          string      `json:"media_preview"`
	MinCommentScore                       int         `json:"min_comment_score"`
	MinLinkScore                          int         `json:"min_link_score"`
	MonitorMentions                       bool        `json:"monitor_mentions"`
	Newwindow                             bool        `json:"newwindow"`
	Nightmode                             bool        `json:"nightmode"`
	NoProfanity                           bool        `json:"no_profanity"`
	NumComments                           int         `json:"num_comments"`
	Numsites                              int         `json:"numsites"`
	Over18                                bool        `json:"over_18"`
	PrivateFeeds                          bool        `json:"private_feeds"`
	ProfileOptOut                         bool        `json:"profile_opt_out"`
	PublicServerSeconds                   bool        `json:"public_server_seconds"`
	PublicVotes                           bool        `json:"public_votes"`
	Research                              bool        `json:"research"`
	SearchIncludeOver18                   bool        `json:"search_include_over_18"`
	SendWelcomeMessages                   bool        `json:"send_welcome_messages"`
	ShowFlair                             bool        `json:"show_flair"`
	ShowGoldExpiration                    bool        `json:"show_gold_expiration"`
	ShowLinkFlair                         bool        `json:"show_link_flair"`
	ShowSnoovatar                         bool        `json:"show_snoovatar"`
	ShowStylesheets                       bool        `json:"show_stylesheets"`
	ShowTrending                          bool        `json:"show_trending"`
	ShowTwitter                           bool        `json:"show_twitter"`
	StoreVisits                           bool        `json:"store_visits"`
	ThirdPartyDataPersonalizedAds         bool        `json:"third_party_data_personalized_ads"`
	ThirdPartySiteDataPersonalizedAds     bool        `json:"third_party_site_data_personalized_ads"`
	ThirdPartySiteDataPersonalizedContent bool        `json:"third_party_site_data_personalized_content"`
	ThreadedMessages                      bool        `json:"threaded_messages"`
	ThreadedModmail                       bool        `json:"threaded_modmail"`
	TopKarmaSubreddits                    bool        `json:"top_karma_subreddits"`
	UseGlobalDefaults                     bool        `json:"use_global_defaults"`
	VideoAutoplay                         bool        `json:"video_autoplay"`
}

// API https://www.reddit.com/dev/api/#GET_api_v1_me_prefs

// TODO
// 1. Patch method for preferences
// 2. Other prefs endpoint response is not complete need to
//    get that response and implement scheme and method accoring to
//    response. For now this methods only save those response to
//    test_data folder locally. Not commited to github.
//
// PRECAUSTION : DO NOT USE THIS METHODS FOR NOW
//

func (c *Client) GetMyPreferences() (*Preferences, error) {
	resp, err := c.Get(API_PATH["preferences"])

	if err != nil {

		log.Fatal("Error in getting preferences response")
	}

	defer resp.Body.Close()

	var preferences Preferences

	prefErr := json.NewDecoder(resp.Body).Decode(&preferences)

	if prefErr != nil {
		log.Fatal("Error in decoding preferences response")
		return nil, err
	}

	return &preferences, nil
}

func (c *Client) PrefsFriends() {

	resp, err := c.Get(API_PATH["prefs_friends"])

	if err != nil {
		log.Fatal("error in getting prefs friends response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/prefs_friends.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}

func (c *Client) PrefsBlocked() {

	resp, err := c.Get(API_PATH["prefs_blocked"])

	if err != nil {
		log.Fatal("error in getting prefs/blocked response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/prefs_blocked.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}
func (c *Client) PrefsMessaging() {

	resp, err := c.Get(API_PATH["prefs_messaging"])

	if err != nil {
		log.Fatal("error in getting prefs/messaging response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/prefs_messaging.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}
func (c *Client) PrefsTrusted() {

	resp, err := c.Get(API_PATH["prefs_trusted"])

	if err != nil {
		log.Fatal("error in getting prefs/trusted response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/prefs_trusted.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}
func (c *Client) MeFriends() {

	resp, err := c.Get(API_PATH["me_friends"])

	if err != nil {
		log.Fatal("error in getting me_friends response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/me_friends.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}
func (c *Client) MeBlocked() {

	resp, err := c.Get(API_PATH["me_blocked"])

	if err != nil {
		log.Fatal("error in getting me_blocked response")
	}

	defer resp.Body.Close()

	SaveResponse(resp.Body, "test_data/me_blocked.json")

	// out, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(out))

}

