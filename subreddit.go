package slashred

import (
	"fmt"
	"log"

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
	s.aboutWhere(subreddit, API_PATH["about_banned"], "error in getting muted response", opts)
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
