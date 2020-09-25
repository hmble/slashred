package slashred

import (
	"fmt"
	"log"
)

type ModerationService service

// Moderation.Log
// Get a list of recent moderation actions.
//
// Moderator actions taken within a subreddit are logged. This listing is a view of
// that log with various filters to aid in analyzing the information.
//
// The optional mod parameter can be a comma-delimited list of moderator names to
// restrict the results to, or the string a to restrict the results to admin
// actions taken within the subreddit.
//
// The type parameter is optional and if sent limits the log entries returned to
// only those of the type specified.
//
// Key             Value
// ==============================
// after        a ModAction ID
// before       a ModAction ID
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 500)
// mod          (optional) a moderator filter
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
// type         (for list of types visit Reference below)
//
// Reference : https://www.reddit.com/dev/api/#GET_about_log

func (m *ModerationService) Log(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/log", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Return a listing of things that have been edited recently
//
// Key             Value
// =================================================
// after        fullname of a thing
// before       fullname of a thing
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 100)
// location     (dont know what should we pass in this parameter)
// only         one of (links, comments)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
//
// Reference: https://www.reddit.com/dev/api/#GET_about_{location}
func (m ModerationService) Edited(subreddit string, opts Option) {

	path := fmt.Sprintf("/r/%s/about/edited", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Return a listing of things that have been reported
//
// Key             Value
// =================================================
// after        fullname of a thing
// before       fullname of a thing
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 100)
// location     (dont know what should we pass in this parameter)
// only         one of (links, comments)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
//
// Reference: https://www.reddit.com/dev/api/#GET_about_{location}

func (m ModerationService) Reports(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/reports", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Return a listing of things that have been marked as spam or otherwise
// removed.
//
// Key             Value
// =================================================
// after        fullname of a thing
// before       fullname of a thing
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 100)
// location     (dont know what should we pass in this parameter)
// only         one of (links, comments)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
//
// Reference: https://www.reddit.com/dev/api/#GET_about_{location}

func (m ModerationService) Spam(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/spam", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Return a listing of things requiring moderator review, such as reported
// things and items caught by the spam filter.
//
// Key             Value
// =================================================
// after        fullname of a thing
// before       fullname of a thing
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 100)
// location     (dont know what should we pass in this parameter)
// only         one of (links, comments)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
//
// Reference: https://www.reddit.com/dev/api/#GET_about_{location}

func (m ModerationService) Modequeue(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/modqueue", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Return a listing of things that have yet to be approved/removed by a mod.
//
// Key             Value
// =================================================
// after        fullname of a thing
// before       fullname of a thing
// count        a positive integer (default: 0)
// limit        the maximum number of items desired (default: 25, maximum: 100)
// location     (dont know what should we pass in this parameter)
// only         one of (links, comments)
// show         (optional) the string all
// sr_detail    (optional) expand subreddits
//
// Reference: https://www.reddit.com/dev/api/#GET_about_{location}

func (m ModerationService) Unmoderated(subreddit string, opts Option) {
	path := fmt.Sprintf("/r/%s/about/unmoderated", subreddit)

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Accept an invite to moderate the specified subreddit.
// The authenticated user must have been invited to moderate the subreddit by
// one of its current moderators.
//
// Reference : https://www.reddit.com/dev/api/#POST_api_accept_moderator_invite
func (m *ModerationService) AcceptInvite(subreddit string) {
	path := fmt.Sprintf("/r/%s/api/accept_moderator_invite", subreddit)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Approve a link or comment.
//
// If the thing was removed, it will be re-inserted into appropriate listings. Any
// reports on the approved thing will be discarded.
//
// Key     Value
// =================
// id      fullname of a thing
//
// Reference: https://www.reddit.com/dev/api/#POST_api_approve

func (m *ModerationService) Approve(id string) {

	path := "/api/approve"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Distinguish a thing's author with a sigil.

// This can be useful to draw attention to and confirm the identity of the user
// in the context of a link or comment of theirs. The options for distinguish
// are as follows:
//
//    yes     - add a moderator distinguish ([M]). only if the user is a moderator
//              of the subreddit the thing is in.
//    no      - remove any distinguishes.
//    admin   - add an admin distinguish ([A]). admin accounts only.
//    special - add a user-specific distinguish. depends on user.

// The first time a top-level comment is moderator distinguished, the author of
// the link the comment is in reply to will get a notification in their inbox.
//
// sticky is a boolean flag for comments, which will stick the distingushed
// comment to the top of all comments threads. If a comment is marked sticky, it
// will override any other stickied comment for that link (as only one comment
// may be stickied at a time.) Only top-level comments may be stickied.
//
// Key         Value
// =============================================
// how      one of (yes, no, admin, special)
// id       fullname of a thing
// sticky   boolean value
//
// Reference: https://www.reddit.com/dev/api/#POST_api_distinguish

func (m *ModerationService) Distinguish(how, fullname, sticky string) {
	path := "/api/distinguish"
	postdata := PostData{
		"how":    how,
		"sticky": sticky,
		"id":     fullname,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Prevent future reports on a thing from causing notifications.

// Any reports made about a thing after this flag is set on it will not cause
// notifications or make the thing show up in the various moderation listings.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) IgnoreReports(id string) {
	path := "/api/ignore_reports"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Allow future reports on a thing to cause notifications.
//
// Key      Value
// ===================
// id       fullname of a thing

func (m *ModerationService) UnignoreReports(id string) {

	path := "/api/unignore_reports"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Abdicate approved user status in a subreddit.
//
// Key      Value
// ===================
// id       fullname of a thing

func (m *ModerationService) LeaveContributor(id string) {
	path := "/api/leavecontributor"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Abdicate moderator status in a subreddit.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) LeaveModerator(id string) {
	path := "/api/leavemoderator"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// For muting user via modmail.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) MuteMessageAuthor(id string) {
	path := "/api/mute_message_author"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// For unmuting user via modmail.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) UnmuteMessageAuthor(id string) {
	path := "/api/unmute_message_author"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Remove a link, comment, or modmail message.
//
// If the thing is a link, it will be removed from all subreddit listings. If the
// thing is a comment, it will be redacted and removed from all subreddit comment
// listings.
//
// Key      Value
// ===================
// id       fullname of a thing
// spam     boolean value
func (m *ModerationService) Remove(fullname, spam string) {
	path := "/api/remove"
	postdata := PostData{
		"id":   fullname,
		"spam": spam,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Mark a comment that it should not be collapsed because of crowd control.
// The comment could still be collapsed for other reasons.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) ShowComment(id string) {
	path := "/api/unmute_message_author"
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}

// Redirect to the subreddit's stylesheet if one exists.
//
// Key      Value
// ===================
// id       fullname of a thing
func (m *ModerationService) Stylesheet(subreddit string) {
	path := fmt.Sprintf("/r/%s/stylesheet", subreddit)
	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		log.Fatal("Error in getting subreddit stylesheet")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// Change the post's crowd control level.
//
// Key      Value
// ===================
// id       fullname of a thing
// level    an integer between 0 and 3oolean value
func (m *ModerationService) UpdateCrowdControlLevel(id, level string) {

	path := "/api/update_crowd_control_level"
	postdata := PostData{
		"id":    id,
		"level": level,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, m.client)
}
