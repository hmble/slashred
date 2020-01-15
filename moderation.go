package slashred

import (
	"fmt"
	"log"
)

type ModerationService service

// Requires the "posts" moderator permission for the subreddit
func (m *ModerationService) aboutLocation(subreddit string, opts Option) {
	endpoint := fmt.Sprintf("/r/%s%s", subreddit, API_PATH["about_edited"])

	fmt.Println(endpoint)
	resp, err := m.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal("Error in getting about/edited repsonse")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

	//SaveResponse(resp.Body, "test_data/about_edited.json")

}

// Test remained
func (m ModerationService) Edited(subreddit string, opts Option) {
	m.aboutLocation(subreddit, opts)
}

// Test remained
func (m ModerationService) Reports(subreddit string, opts Option) {
	m.aboutLocation(subreddit, opts)
}

// Test remained
func (m ModerationService) Spam(subreddit string, opts Option) {
	m.aboutLocation(subreddit, opts)
}

// Test remained
func (m ModerationService) Modequeue(subreddit string, opts Option) {
	m.aboutLocation(subreddit, opts)
}

// Test remained
func (m ModerationService) Unmoderated(subreddit string, opts Option) {
	m.aboutLocation(subreddit, opts)
}

// Post method

// See
// https://www.reddit.com/dev/api/#POST_api_friend
// https://www.reddit.com/dev/api/#GET_subreddits_mine_%7Bwhere%7D
func (m *ModerationService) AcceptInvite() {
	postdata := PostData{}
	resp, err := m.client.Post(API_PATH["moderator_invite"], postdata)

	if err != nil {
		log.Fatal("Error in acception invite")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (m *ModerationService) idOnly(fullname, endpoint, logmsg string) {
	postdata := PostData{
		"id": fullname,
	}

	resp, err := m.client.Post(endpoint, postdata)

	if err != nil {
		log.Fatal(logmsg)

	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

// fullname of link or comment
func (m *ModerationService) Approve(fullname string) {
	m.idOnly(fullname, API_PATH["approve"], "Error in accepting invite")
}

// https://www.reddit.com/dev/api/#POST_api_distinguish
// sticky is bool string
func (m *ModerationService) Distinguish(how, fullname, sticky string) {
	postdata := PostData{
		"how":    how,
		"sticky": sticky,
		"id":     fullname,
	}

	resp, err := m.client.Post(API_PATH["distinguish"], postdata)

	if err != nil {
		log.Fatal("Error in distinguishing post or comment")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// Also see https://www.reddit.com/dev/api/#POST_api_unignore_reports

func (m *ModerationService) IgnoreReports(fullname string) {

	m.idOnly(fullname, API_PATH["ignore_reports"], "Error in getting ignore reports")
}
func (m *ModerationService) UnignoreReports(fullname string) {

	m.idOnly(fullname, API_PATH["unignore_reports"], "Error in getting unignore reports")
}
func (m *ModerationService) LeaveContributor(fullname string) {

	m.idOnly(fullname, API_PATH["leave_contributor"], "Error in Leaving as contributor")
}

func (m *ModerationService) LeaveModerator(fullname string) {

	m.idOnly(fullname, API_PATH["leave_moderator"], "Error in leaving as moderator")
}
func (m *ModerationService) MuteMessageAuthor(fullname string) {

	m.idOnly(fullname, API_PATH["mute_message_author"], "Error in muting message author")
}
func (m *ModerationService) UnmuteMessageAuthor(fullname string) {

	m.idOnly(fullname, API_PATH["unmute_message_author"], "Error in unmuting message author")
}

func (m *ModerationService) Remove(fullname, spam string) {
	postdata := PostData{
		"id":   fullname,
		"spam": spam,
	}

	resp, err := m.client.Post(API_PATH["remove"], postdata)

	if err != nil {
		log.Fatal("Error in removing")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

}

func (m *ModerationService) ShowComment(fullname string) {

	m.idOnly(fullname, API_PATH["show_comment"], "Error in showing comment")
}

// Redirect to subreddit stylesheet
func (m *ModerationService) Stylesheet(subreddit string) {
	endpoint := fmt.Sprintf("r/%s%s", subreddit, API_PATH["stylesheet"])
	resp, err := m.client.Get(endpoint, NoOptions)

	if err != nil {
		log.Fatal("Error in getting subreddit stylesheet")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}
