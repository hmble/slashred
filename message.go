package slashred

import "strings"

type MessageService service

// Post Methods

func (m *MessageService) id(path, id string) {
	postdata := PostData{
		"id": id,
	}

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Message.Block
func (m *MessageService) Block(id string) {
	m.id("/api/block", id)
}

// Message.Collapse
func (m *MessageService) Collapse(id string) {
	m.id("/api/collapse_message", id)
}

// Message.Compose
func (m *MessageService) Compose(postdata PostData) {
	path := "/api/compose"

	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Message.Delete
func (m *MessageService) Delete(id string) {
	m.id("/api/del_msg", id)
}

// Message.ReadAll

func (m *MessageService) ReadAll(filterTypes []string) {
	path := "/api/read_all_messages"

	postdata := PostData{
		"filter_types": strings.Join(filterTypes, ","),
	}
	resp, err := m.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Message.ReadMessage
func (m *MessageService) ReadMessage(id string) {
	m.id("/api/read_message", id)
}

// Message.UnblockSubreddit
func (m *MessageService) UnblockSubreddit(id string) {
	m.id("/api/unblock_subreddit", id)
}

// Message.Uncollapse
func (m *MessageService) Uncollapse(id string) {
	m.id("/api/uncollapse_message", id)
}

// Message.Unread
func (m *MessageService) Unread(id string) {
	m.id("/api/unread_message", id)
}

// Get Methods

func (m *MessageService) getwhere(path string, opts Option) {
	resp, err := m.client.Get(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()
}

// Message.GetInbox
func (m *MessageService) GetInbox(opts Option) {
	m.getwhere("/message/inbox", opts)
}

// Message.GetUnread
func (m *MessageService) GetUnread(opts Option) {
	m.getwhere("/message/unread", opts)
}

// Message.GetSent
func (m *MessageService) GetSent(opts Option) {
	m.getwhere("/message/sent", opts)
}
