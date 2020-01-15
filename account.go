package slashred

import (
	"encoding/json"
	"fmt"
	"log"
)

type AccountService service
type Account struct {
	CommentKarma int     `json:"comment_karma"`
	Created      float32 `json:"created"`
	CreatedUtc   float32 `json:"created_utc"`
	Features     struct {
		ActivityServiceRead    bool `json:"activity_service_read"`
		ActivityServiceWrite   bool `json:"activity_service_write"`
		AdblockTest            bool `json:"adblock_test"`
		AdsAuction             bool `json:"ads_auction"`
		AdsAutoExtend          bool `json:"ads_auto_extend"`
		AdsAutoRefund          bool `json:"ads_auto_refund"`
		AdserverReporting      bool `json:"adserver_reporting"`
		AdzerkDoNotTrack       bool `json:"adzerk_do_not_track"`
		AdzerkReporting2       bool `json:"adzerk_reporting_2"`
		DoNotTrack             bool `json:"do_not_track"`
		EuCookiePolicy         bool `json:"eu_cookie_policy"`
		ExpandoEvents          bool `json:"expando_events"`
		ForceHTTPS             bool `json:"force_https"`
		GiveHstsGrants         bool `json:"give_hsts_grants"`
		HTTPSRedirect          bool `json:"https_redirect"`
		ImageUploads           bool `json:"image_uploads"`
		ImgurGifConversion     bool `json:"imgur_gif_conversion"`
		LegacySearchPref       bool `json:"legacy_search_pref"`
		LiveHappeningNow       bool `json:"live_happening_now"`
		MoatTracking           bool `json:"moat_tracking"`
		MobileNativeBanner     bool `json:"mobile_native_banner"`
		MobileSettings         bool `json:"mobile_settings"`
		MobileWebTargeting     bool `json:"mobile_web_targeting"`
		NewLoggedinCachePolicy bool `json:"new_loggedin_cache_policy"`
		NewReportDialog        bool `json:"new_report_dialog"`
		OrangeredsAsEmails     bool `json:"orangereds_as_emails"`
		OutboundClicktracking  bool `json:"outbound_clicktracking"`
		PauseAds               bool `json:"pause_ads"`
		PostEmbed              bool `json:"post_embed"`
		ScreenviewEvents       bool `json:"screenview_events"`
		ScrollEvents           bool `json:"scroll_events"`
		ShowNewIcons           bool `json:"show_new_icons"`
		StickyComments         bool `json:"sticky_comments"`
		SubredditRules         bool `json:"subreddit_rules"`
		Timeouts               bool `json:"timeouts"`
		UpgradeCookies         bool `json:"upgrade_cookies"`
		YoutubeScraper         bool `json:"youtube_scraper"`
	} `json:"features"`
	GoldCreddits            int    `json:"gold_creddits"`
	GoldExpiration          int    `json:"gold_expiration"`
	HasVerifiedEmail        bool   `json:"has_verified_email"`
	HideFromRobots          bool   `json:"hide_from_robots"`
	ID                      string `json:"id"`
	InBeta                  bool   `json:"in_beta"`
	InboxCount              int    `json:"inbox_count"`
	IsEmployee              bool   `json:"is_employee"`
	IsGold                  bool   `json:"is_gold"`
	IsMod                   bool   `json:"is_mod"`
	IsSuspended             bool   `json:"is_suspended"`
	LinkKarma               int    `json:"link_karma"`
	Name                    string `json:"name"`
	Over18                  bool   `json:"over_18"`
	SuspensionExpirationUtc int    `json:"suspension_expiration_utc"`
}

// GetMe retrieves the user account for the currently authenticated user. Requires the 'identity' OAuth scope.
func (a *AccountService) GetMe() (*Account, error) {

	//url := fmt.Sprintf("%s/api/v1/me", BaseAuthURL)

	resp, err := a.client.Get(API_PATH["me"], NoOptions)
	if err != nil {
		log.Fatal(err, "Error in response Do")
		return nil, err
	}
	defer resp.Body.Close()

	PrintHeader(resp)
	var account Account

	//SaveResponse(resp.Body, "account.json")
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		log.Fatal("Error while decoding Account response")
		return nil, err
	}

	return &account, nil
}

type Karma struct {
	CommentKarma int    `json:"comment_karma"`
	LinkKarma    int    `json:"link_karma"`
	Subreddit    string `json:"sr"`
}

func (a *AccountService) GetKarma() ([]Karma, error) {
	resp, err := a.client.Get(API_PATH["karma"], NoOptions)

	if err != nil {
		log.Fatal("Error in getting karma response from /api/v1/me/karma")
	}

	defer resp.Body.Close()
	PrintHeader(resp)

	var karmaListing struct {
		Kind      string  `json:"kind"`
		KarmaList []Karma `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&karmaListing)
	if err != nil {
		log.Fatal("Error while decoding Karma response")
		return nil, err
	}

	fmt.Println(karmaListing.Kind)

	return karmaListing.KarmaList, nil
}

