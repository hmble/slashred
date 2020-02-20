package slashred

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type SubredditService service

func (s *SubredditService) aboutWhere(subreddit, endpoint, logmsg string, opts Option) {
	url := fmt.Sprintf("/r/%s%s", subreddit, endpoint)

	resp, err := s.client.Get(url, opts)

	if err != nil {
		log.Fatal(logmsg)
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) GetBanned(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_banned"], "Error in getting banned response", opts)
}

func (s *SubredditService) GetMuted(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_muted"], "error in getting muted response", opts)
}

func (s *SubredditService) GetWikiBanned(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_wikibanned"], "Error in getting wiki banned response", opts)
}
func (s *SubredditService) GetContributor(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_contributors"], "Error in getting cotributors response", opts)
}
func (s *SubredditService) GetWikiContributor(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_wikicontributors"], "Error in getting wikicontributors response", opts)
}
func (s *SubredditService) GetModerators(subreddit string, opts Option) {
	s.aboutWhere(subreddit, API_PATH["about_moderators"], "Error in getting moderators response", opts)
}

func (s *SubredditService) deleteSrWhere(endpoint, logmsg string) {

	postdata := PostData{}

	resp, err := s.client.Post(endpoint, postdata)

	if err != nil {
		log.Fatal(logmsg)
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}
func (s *SubredditService) DeleteSrBanner() {
	s.deleteSrWhere(API_PATH["delete_sr_banner"], "Error in deleting banner")
}
func (s *SubredditService) DeleteSrHeader() {
	s.deleteSrWhere(API_PATH["delete_sr_header"], "Error in deleting header")
}
func (s *SubredditService) DeleteSrIcon() {
	s.deleteSrWhere(API_PATH["delete_sr_icon"], "Error in deleting icon")
}
func (s *SubredditService) DeleteSrImg(imgName string) {
	postdata := PostData{

		"img_name": imgName,
	}

	resp, err := s.client.Post(API_PATH["delete_sr_img"], postdata)

	if err != nil {
		log.Fatal("Error in deleting image")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// query should be 50 characters long
// https://www.reddit.com/dev/api/#GET_api_search_reddit_names
// Some option parameter have been excluded
func (s *SubredditService) SearchSrNames(exact, over18, query string) {

	if len(query) > 50 {
		log.Fatal("Query length should be less than 50 characters")
		return
	}
	opts := Option{
		"exact":           exact,
		"over_18":         over18,
		"query":           query,
		"search_query_id": uuid.New().String(),
	}

	resp, err := s.client.Get(API_PATH["search_reddit_names"], opts)

	if err != nil {
		log.Fatal("Error in getting searched results for given query")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// See https://www.reddit.com/dev/api/#POST_api_site_admin
// For option visit above link and make options as user like
//
// NOTE : This endpoint expects all options on every request
// TODO : An easier way to create postdata for site_admin
// Although this method is used only once in while as it only
// creates a subreddit and edit its options.
//

func (s *SubredditService) SiteAdmin(postdata PostData) {
	resp, err := s.client.Post(API_PATH["site_admin"], postdata)

	if err != nil {
		log.Fatal("Error in creating or editing subreddit")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) SubmitText(subreddit string) {

	endpoint := fmt.Sprintf("/r/%s/%s", subreddit, API_PATH["submit_text"])
	resp, err := s.client.Get(endpoint, NoOptions)

	if err != nil {
		log.Fatal("Error in getting submit text")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) GetAutocomplete(over18, includeProfiles, query string) {

	if len(query) > 50 {
		log.Fatal("Query length should be less than 50")
		return
	}
	opts := Option{
		"include_over_18":  over18,
		"include_profiles": includeProfiles,
		"query":            query,
	}

	resp, err := s.client.Get(API_PATH["subreddit_autocomplete"], opts)

	if err != nil {
		log.Fatal("Error in getting autocomplete results")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// limit should be between 1 to 10
func (s *SubredditService) GetAutocompleteV2(includeCategories, over18, includeProfiles, query string, limit int) {
	if limit >= 1 && limit <= 10 || len(query) < 50 {
		log.Fatal("Limit should be between 1 to 10 and query length less than 50")
	}

	opts := Option{
		"include_over_18":    over18,
		"include_profiles":   includeProfiles,
		"query":              query,
		"include_categories": includeCategories,
		"limit":              strconv.Itoa(limit),
		"search_query_id":    uuid.New().String(),
	}

	resp, err := s.client.Get(API_PATH["subreddit_autocomplete"], opts)

	if err != nil {
		log.Fatal("Error in getting autocomplete results")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) SubmitStylesheet(subreddit, op, reason, content string) {
	postdata := PostData{
		"op":                 op,
		"reason":             reason,
		"stylesheet_content": content,
	}

	endpoint := fmt.Sprintf("/r/%s/%s", subreddit, API_PATH["subreddit_stylesheet"])
	resp, err := s.client.Post(endpoint, postdata)

	if err != nil {
		log.Fatal("Error in submitting stylesheet")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// action : one of (sub, unsub)
// Note : skipped option of sr fullname instead used sr_name
// for simplicity
func (s *SubredditService) Subscribe(action, skipDefaults string, sr_names []string) {

	srname := strings.Join(sr_names, ",")
	postdata := PostData{
		"action":               action,
		"skip_inital_defaults": skipDefaults,
		"sr_name":              srname,
	}

	resp, err := s.client.Post(API_PATH["subscribe"], postdata)

	if err != nil {
		log.Fatal("Error in subscribing/unsubscribing to subreddit")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

/*
Add or replace a subreddit image, custom header logo, custom mobile icon, or
custom mobile banner.

    If the upload_type value is img, an image for use in the subreddit
    stylesheet is uploaded with the name specified in name. If the upload_type
    value is header then the image uploaded will be the subreddit's new logo and
    name will be ignored. If the upload_type value is icon then the image
    uploaded will be the subreddit's new mobile icon and name will be ignored.
    If the upload_type value is banner then the image uploaded will be the
    subreddit's new mobile banner and name will be ignored.

For backwards compatibility, if upload_type is not specified, the header field
will be used instead:

    If the header field has value 0, then upload_type is img. If the header
    field has value 1, then upload_type is header.

The img_type field specifies whether to store the uploaded image as a PNG or
JPEG.

Subreddits have a limited number of images that can be in use at any given time.
If no image with the specified name already exists, one of the slots will be
consumed.

If an image with the specified name already exists, it will be replaced. This
does not affect the stylesheet immediately, but will take effect the next time
the stylesheet is saved.
*/
func (s *SubredditService) UploadSrImg(subreddit, file, header, imgType, name,
	uploadType string) {
	postdata := PostData{
		"file":        file,
		"header":      header,
		"img_type":    imgType,
		"name":        name,
		"upload_type": uploadType,
	}

	endpoint := fmt.Sprintf("/r/%s/%s", subreddit, API_PATH["upload_sr_img"])
	resp, err := s.client.Post(endpoint, postdata)
	if err != nil {
		log.Fatal("Error in uploading image to subreddit")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) about(endpoint, subreddit string) {
	url := fmt.Sprintf("/r/%s/%s", subreddit, endpoint)
	resp, err := s.client.Get(url, NoOptions)

	if err != nil {
		log.Fatal("Error in getting about of subreddit")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	//	SaveResponse(resp.Body, "test_data/astar0n_about.json")

}

func (s *SubredditService) About(subreddit string) {
	s.about(API_PATH["about_subreddit"], subreddit)
}
func (s *SubredditService) AboutRules(subreddit string) {
	s.about(API_PATH["rules"], subreddit)

}
func (s *SubredditService) Traffic(subreddit string) {
	s.about(API_PATH["traffic"], subreddit)

}
func (s *SubredditService) Sidebar(subreddit string) {
	s.about(API_PATH["sidebar"], subreddit)

}

// num : default int between 1 and 2
func (s *SubredditService) Sticky(subreddit, num string) {
	// TODO: check for num
	url := fmt.Sprintf("/r/%s/%s", subreddit, API_PATH["sticky"])
	opts := Option{
		"num": num,
	}
	resp, err := s.client.Get(url, opts)

	if err != nil {
		log.Fatal("Error in getting about of subreddit")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (s *SubredditService) MineSubscriber(opts Option) {
	logmsg := "Error in getting mine subscriber"
	s.client.getlisting(API_PATH["mine_subscriber"], logmsg, opts)
}

func (s *SubredditService) MineContributor(opts Option) {
	logmsg := "Error in getting mine contributor"
	s.client.getlisting(API_PATH["mine_contributor"], logmsg, opts)
}

func (s *SubredditService) MineModerator(opts Option) {
	logmsg := "Error in getting mine moderator"
	s.client.getlisting(API_PATH["mine_moderator"], logmsg, opts)
}
func (s *SubredditService) MineStreams(opts Option) {
	logmsg := "Error in getting mine streams"
	s.client.getlisting(API_PATH["mine_streams"], logmsg, opts)
}

func (s *SubredditService) Search(opts Option) {
	logmsg := "Error in subreddit search"

	s.client.getlisting(API_PATH["subreddit_search"], logmsg, opts)

}

func (s *SubredditService) Popular(opts Option) {
	logmsg := "Error in getting subreddit popular"

	s.client.getlisting(API_PATH["subreddit_popular"], logmsg, opts)

}

func (s *SubredditService) Gold(opts Option) {
	logmsg := "Error in getting subreddit gold"

	s.client.getlisting(API_PATH["subreddit_gold"], logmsg, opts)

}

func (s *SubredditService) New(opts Option) {
	logmsg := "Error in getting subreddit New"

	s.client.getlisting(API_PATH["subreddit_new"], logmsg, opts)

}

func (s *SubredditService) Default(opts Option) {
	logmsg := "Error in getting subreddit default "

	s.client.getlisting(API_PATH["subreddit_default"], logmsg, opts)

}

// Subreddit.About
func (s *SubredditService) AboutSubreddit(subreddit string) {
	path := fmt.Sprintf("/r/%s/about", subreddit)

	resp, err := s.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

}

// Subreddit.EditAbout
func (s *SubredditService) EditAbout(subreddit, created, location string) {
	path := fmt.Sprintf("/r/%s/about/edit", subreddit)

	opts := Option{
		"created":  created,
		"location": location,
	}
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

}

func (s *SubredditService) userwhere(path string, opts Option) {
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

}

// Subreddit.UserSearch

func (s *SubredditService) UserSearch(opts Option) {
	s.userwhere("/users/search", opts)
}

// Subreddit.UserPopular

func (s *SubredditService) UserPopular(opts Option) {
	s.userwhere("/users/popular", opts)
}

// Subreddit.UserNew

func (s *SubredditService) UserNew(opts Option) {
	s.userwhere("/users/new", opts)
}
