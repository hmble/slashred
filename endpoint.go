package slashred

var API_PATH map[string]string = map[string]string{
	"me":                  "/api/v1/me",
	"karma":               "/api/v1/me/karma",
	"trophies":            "/api/v1/me/trophies",
	"preferences":         "/api/v1/me/prefs",
	"prefs_friends":       "/prefs/friends",
	"prefs_blocked":       "/prefs/blocked",
	"prefs_messaging":     "/prefs/messaging",
	"prefs_trusted":       "/prefs/trusted",
	"me_friends":          "/api/v1/me/friends",
	"me_blocked":          "/api/v1/me/blocked",
	"collection":          "/api/v1/collections/collection",
	"trending_subreddits": "/api/trending_subreddits",
	"by_id":               "/by_id/",
	"best":                "/best",
}
