package slashred

import (
	"fmt"
	"strings"
)

type ModmailService service

// Get Methods

// Modmail.GetConversations

func (m *ModmailService) GetConversations(opts Option) {
	path := "/api/mod/conversations"

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

}

// Modmail.ConversationByID
func (m *ModmailService) GetConversationByID(id, markread string) {
	path := fmt.Sprintf("/api/mod/conversations/%s", id)

	opts := Option{
		"markRead": markread,
	}

	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.ConversationWithUser
func (m *ModmailService) ConversationWithUser(id string) {
	path := fmt.Sprintf("/api/mod/conversations/%s/user", id)

	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.SubredditConverstions
func (m *ModmailService) SubredditConverstions() {
	path := "/api/mod/conversations/subreddits"

	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.UnreadCount
func (m *ModmailService) UnreadCount() {
	path := "/api/mod/conversations/unread/count"

	resp, err := m.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Post Methods

// Modmail.BulkRead

func (m *ModmailService) BulkRead(entity, state string) {
	path := "/api/mod/bulk_read"

	postdata := PostData{
		"entity": entity,
		"state":  state,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Conversions

func (m *ModmailService) Conversions(subreddit, body, isAuthorHidden, subject, to string) {
	path := "/api/mod/conversations"

	postdata := PostData{
		"srName":         subreddit,
		"isAuthorHidden": isAuthorHidden,
		"body":           body,
		"subject":        subject,
		"to":             to,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.ConversionsByID

func (m *ModmailService) ConversionsByID(id, body, isAuthorHidden, isInternal string) {
	path := fmt.Sprintf("/api/mod/conversations/%s", id)

	postdata := PostData{
		"isAuthorHidden": isAuthorHidden,
		"body":           body,
		"isInternal":     isInternal,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Archive
func (m *ModmailService) Archive(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/archive", id)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Highlight
func (m *ModmailService) Highlight(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/highlight", id)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.HighlightDelete
func (m *ModmailService) HighlightDelete(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/highlight", id)
	resp, err := m.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Mute
func (m *ModmailService) Mute(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/mute", id)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Unarchive
func (m *ModmailService) Unarchive(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/unarchive", id)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Unmute
func (m *ModmailService) Unmute(id string) {

	path := fmt.Sprintf("/api/mod/conversations/%s/unmute", id)
	resp, err := m.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Read
func (m *ModmailService) Read(ids []string) {
	path := "/api/mod/conversations/read"

	postdata := PostData{
		"conversationIds": strings.Join(ids, ","),
	}
	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Modmail.Unread
func (m *ModmailService) Unread(ids []string) {
	path := "/api/mod/conversations/unread"

	postdata := PostData{
		"conversationIds": strings.Join(ids, ","),
	}
	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}
