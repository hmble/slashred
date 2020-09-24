package slashred

import (
	"encoding/json"
	"fmt"

	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type SubredditService service

func (s *SubredditService) aboutWhere(subreddit, path string, opts Option) {

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
//
// Key              Value
// ================================================
// after        fullname of thing
// before       fullname of thing
// count        a positive integer
// limit        a maximum number of items desired (default: 25, max: 100)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
// username     A valid, existing username

func (s *SubredditService) GetBanned(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/banned", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
// user        A valid, existing reddit username

func (s *SubredditService) GetMuted(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/muted", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
// user        A valid, existing reddit username
func (s *SubredditService) GetWikiBanned(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/wikibanned", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
// user        A valid, existing reddit username
func (s *SubredditService) GetContributor(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/contributors", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
// user        A valid, existing reddit username
func (s *SubredditService) GetWikiContributor(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/wikicontributors", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

// This endpoint is a listing
// This returns a list of banned persons in a subreddit
//
// Option fields
// Note: Option is map[string]string
//       All fields key values should be string
// Example:
//      opts := Option{
//          "count": "2",
//          "limit": "25",
//          "show": "alll",
//      }
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
// user        A valid, existing reddit username
func (s *SubredditService) GetModerators(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/moderators", subreddit)
	s.aboutWhere(subreddit, path, opts)
}

func (s *SubredditService) deleteSrWhere(path string) {

	resp, err := s.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Remove the subreddit's custom mobile banner.
//
// See also method UploadSrImg
//
// Reference: https://www.reddit.com/dev/api/#POST_api_delete_sr_banner
func (s *SubredditService) DeleteSrBanner(subreddit string) {
	path := fmt.Sprintf("/r/%s/api/delete_sr_banner", subreddit)
	s.deleteSrWhere(path)
}

// Remove the subreddit's custom header image.
//
// The sitewide-default header image will be shown again after this call.
//
// Reference: https://www.reddit.com/dev/api/#POST_api_delete_sr_header
func (s *SubredditService) DeleteSrHeader(subreddit string) {
	path := fmt.Sprintf("/r/%s/api/delete_sr_header", subreddit)
	s.deleteSrWhere(path)
}

// Remove the subreddit's custom mobile icon.
//
// See also method UploadSrImg
//
// Reference: https://www.reddit.com/dev/api/#POST_api_delete_sr_icon
func (s *SubredditService) DeleteSrIcon(subreddit string) {
	path := fmt.Sprintf("/r/%s/api/delete_sr_icon", subreddit)
	s.deleteSrWhere(path)
}

// Remove an image from the subreddit's custom image set.

// The image will no longer count against the subreddit's image limit. However,
// the actual image data may still be accessible for an unspecified amount of
// time. If the image is currently referenced by the subreddit's stylesheet,
// that stylesheet will no longer validate and won't be editable until the image
// reference is removed.
//
// Reference: https://www.reddit.com/dev/api/#POST_api_delete_sr_img
func (s *SubredditService) DeleteSrImg(subreddit, imgName string) {
	path := fmt.Sprintf("/r/%s/api/delete_sr_img", subreddit)
	postdata := PostData{
		"img_name": imgName,
	}

	resp, err := s.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Some option parameter have been excluded

// List subreddit names that begin with a query string.

// Subreddits whose names begin with query will be returned. If include_over_18
// is false, subreddits with over-18 content restrictions will be filtered from
// the results.
//
// Reference: https://www.reddit.com/dev/api/#GET_api_search_reddit_names
//
//  Key                 Value
// ===========================================
// exact                boolean value
// include_over_18      boolean value
// query                a string up to 50 characters long,
//                      consisting of printable characters.

func (s *SubredditService) SearchSrNames(exact, includeOver18, query string) {

	path := "/api/search_reddit_names"

	if len(query) > 50 {
		log.Fatal("Query length should be less than 50 characters")
		return
	}
	opts := Option{
		"exact":           exact,
		"include_over_18": includeOver18,
		"query":           query,
		"search_query_id": uuid.New().String(),
	}

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
	printBytes(resp.Body, s.client)

}

// NOTE : This endpoint expects all options on every request
//
// An easier way is to pass postdata for site_admin by user because
// this method is used only once in while as it
// creates a subreddit and edit its options.
//
// For list of options see reference
//
// Reference: https://www.reddit.com/dev/api/#POST_api_site_admin
//

func (s *SubredditService) SiteAdmin(postdata PostData) {
	path := "/api/site_admin"
	resp, err := s.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Get the submission text for the subreddit.

// This text is set by the subreddit moderators and intended to be displayed on
// the submission form.
//
// Reference: https://www.reddit.com/dev/api/#GET_api_submit_text
func (s *SubredditService) SubmitText(subreddit string) {

	path := fmt.Sprintf("/r/%s/api/submit_text", subreddit)
	resp, err := s.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Return a list of subreddits and data for subreddits whose names start with 'query'.
// Key                   Value
// ========================================
// include_over_18      boolean value
// include_profiles     boolean value
// query                a string up to 25 characters long,
//                      consisting of printable characters.

func (s *SubredditService) GetAutocomplete(includeOver18, includeProfiles, query string) {

	path := "/api/subreddit_autocomplete"
	if len(query) > 50 {
		log.Fatal("Query length should be less than 50")
		return
	}
	opts := Option{
		"include_over_18":  includeOver18,
		"include_profiles": includeProfiles,
		"query":            query,
	}

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Typeahead provides exact matches, typo correction, fuzzy matching and boosts
// subreddits to the top that the user is subscribed to.
//
// Key                     Value
// =============================================
// include_over_18      boolean value
// include_profiles     boolean value
// limit                an integer between 1 and 10 (default: 5)
// query                a string up to 25 characters long,
//                      consisting of printable characters.
// search_query_id      a uuid
// typeahead_active     boolean value or None
//
// Reference: https://www.reddit.com/dev/api/#GET_api_subreddit_autocomplete_v2

func (s *SubredditService) GetAutocompleteV2(includeCategories, over18, includeProfiles, query string, limit int) {

	path := "/api/subreddit_autocomplete_v2"
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

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Update a subreddit's stylesheet.
// op should be save to update the contents of the stylesheet.
//
// Key                         Value
// ==================================================
// op                       one of (save, preview)
// reason                   a string up to 256 characters long,
//                          consisting of printable characters.
// stylesheet_contents      the new stylesheet content
//
// Reference : https://www.reddit.com/dev/api/#POST_api_subreddit_stylesheet

func (s *SubredditService) SubmitStylesheet(subreddit, op, reason, content string) {
	path := fmt.Sprintf("/r/%s/api/subreddit_stylesheet", subreddit)
	postdata := PostData{
		"op":                 op,
		"reason":             reason,
		"stylesheet_content": content,
	}

	resp, err := s.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Subscribe to or unsubscribe from a subreddit.

// To subscribe, action should be sub. To unsubscribe, action should be unsub. The
// user must have access to the subreddit to be able to subscribe to it.

// The skip_initial_defaults param can be set to True to prevent automatically
// subscribing the user to the current set of defaults when they take their first
// subscription action. Attempting to set it for an unsubscribe action will result
// in an error.
//
// Key                         Value
// ==================================================
// action                   one of (sub, unsub)
// skip_initial_defaults    boolean value
// sr_name                  A comma-separated list of subreddit names.
//
// Reference : https://www.reddit.com/dev/api/#POST_api_subscribe

func (s *SubredditService) Subscribe(action, skipDefaults string, sr_names []string) {

	path := "/api/subscribe"
	srname := strings.Join(sr_names, ",")
	postdata := PostData{
		"action":               action,
		"skip_inital_defaults": skipDefaults,
		"sr_name":              srname,
	}

	resp, err := s.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

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


Image should of size 256 x 256
*/
func (s *SubredditService) UploadSrImg(subreddit, path,
	uploadType string) {
	endpoint := fmt.Sprintf("/r/%s/api/upload_sr_img", subreddit)

	resp, err := s.client.PostImageUpload(endpoint, PostData{
		"upload_type": uploadType,
		"img_type":    "png",
	}, path)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

	//PrintHeader(resp)

	var image_response struct {
		Errors       []string `json:"errors"`
		Img_src      string   `json:"img_src"`
		Error_values string   `json:"error_values"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&image_response); err != nil {
		log.Fatal("Error in reading response body response ", err)
	}

	fmt.Println(image_response)

}

func (s *SubredditService) about(path, subreddit string) {
	resp, err := s.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Return information about the subreddit.
// Data includes the subscriber count, description, and header image.
func (s *SubredditService) About(subreddit string) {
	path := fmt.Sprintf("/r/%s/about", subreddit)
	s.about(path, subreddit)
}

// Get the rules for the current subreddit
func (s *SubredditService) AboutRules(subreddit string) {
	path := fmt.Sprintf("/r/%s/about/rules", subreddit)
	s.about(path, subreddit)
}
func (s *SubredditService) AboutTraffic(subreddit string) {
	path := fmt.Sprintf("/r/%s/about/traffic", subreddit)
	s.about(path, subreddit)
}

// Get the sidebar for the current subreddit
func (s *SubredditService) AboutSidebar(subreddit string) {
	path := fmt.Sprintf("/r/%s/about/sidebar", subreddit)
	s.about(path, subreddit)
}

// num : default int between 1 and 2

// Redirect to one of the posts stickied in the current subreddit

// The "num" argument can be used to select a specific sticky, and will default to
// 1 (the top sticky) if not specified. Will 404 if there is not currently a sticky
// post in this subreddit.
// Reference : https://www.reddit.com/dev/api/#GET_sticky

// TODO(hmble): Check how this endpoint works. Also check what happens if we
// pass num value more than 2
func (s *SubredditService) AboutSticky(subreddit, num string) {
	// TODO: check for num ?

	path := fmt.Sprintf("/r/%s/about/sticky", subreddit)
	opts := Option{
		"num": num,
	}
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Get subreddits the user is subscribed to
//
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
func (s *SubredditService) MineSubscriber(opts Option) {
	path := "/subreddits/mind/subscriber"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Get subreddits the user is approved user in
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
func (s *SubredditService) MineContributor(opts Option) {
	path := "/subreddits/mind/contributor"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Get subreddits the user is moderator of
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
func (s *SubredditService) MineModerator(opts Option) {
	path := "/subreddits/mind/moderator"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Get subscribed to subreddits that contain hosted video links
// Key              Value
// ================================================
// after       fullname of a thing
// before      fullname of a thing
// count       a positive integer (default: 0)
// limit       the maximum number of items desired (default: 25, maximum: 100)
// show        (optional) the string all
// sr_detail   (optional) expand subreddits
func (s *SubredditService) MineStreams(opts Option) {
	path := "/subreddits/mind/streams"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Search subreddits by title and description.
//
// This endpoint is a listing
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// q                    a search query
// search_query_id      a uuid
// show                 (optional) the string all
// show_users           boolean value
// sort                 one of (relevance, activity)
// sr_detail            (optional) expand subreddits
// typeahead_active     boolean value or None
//
// Reference : https://www.reddit.com/dev/api/#GET_subreddits_search

func (s *SubredditService) Search(opts Option) {
	path := "/subreddits/search"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)
}

// Get all subreddits by sort category popular
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits

// Reference: https://www.reddit.com/dev/api/#GET_subreddits_{where}
func (s *SubredditService) Popular(opts Option) {
	path := "/subreddits/popular"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Get all subreddits by sort category gold
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits
// Reference: https://www.reddit.com/dev/api/#GET_subreddits_{where}
func (s *SubredditService) Gold(opts Option) {
	path := "/subreddits/gold"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Get all subreddits by sort category new
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits
// Reference: https://www.reddit.com/dev/api/#GET_subreddits_{where}
func (s *SubredditService) New(opts Option) {
	path := "/subreddits/new"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Get all subreddits by sort category default
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits
// Reference: https://www.reddit.com/dev/api/#GET_subreddits_{where}
func (s *SubredditService) Default(opts Option) {
	path := "/subreddits/default"

	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, s.client)

}

// Return information about the subreddit.
// Data includes the subscriber count, description, and header image.
func (s *SubredditService) AboutSubreddit(subreddit string) {
	path := fmt.Sprintf("/r/%s/about", subreddit)

	resp, err := s.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

}

// Get the current settings of a subreddit.
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

// Search user profiles by title and description.
//
// This endpoint is a listing
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// q                    a search query
// search_query_id      a uuid
// show                 (optional) the string all
// show_users           boolean value
// sort                 one of (relevance, activity)
// sr_detail            (optional) expand subreddits
// typeahead_active     boolean value or None
func (s *SubredditService) UserSearch(opts Option) {
	path := "/users/search"
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Get all users subreddits by category popular
// This endpoint is a listing
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits
func (s *SubredditService) UserPopular(opts Option) {
	path := "/users/popular"
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Get all users subreddits by category new
// This endpoint is a listing
//
// Key                     Value
// ==================================================
// after                fullname of a thing
// before               fullname of a thing
// count                a positive integer (default: 0)
// limit                the maximum number of items desired (default: 25, maximum: 100)
// show                 (optional) the string all
// sr_detail            (optional) expand subreddits
func (s *SubredditService) UserNew(opts Option) {
	path := "/users/new"
	resp, err := s.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}
