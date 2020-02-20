
## Account

[Api Refernce](https://www.reddit.com/dev/api/#section_account)


- [x] GET   /api/v1/me          ( any )
- [x] GET   /api/v1/me/karma    ( mysubreddits )
- [x] GET   /api/v1/me/prefs    ( identity )
- [ ] PATCH /api/v1/me/prefs    ( account )
- [x] GET   /api/v1/me/trophies ( identity )
- [x] GET   /prefs/where        ( read ) ( rss support )
  - [x] /prefs/friends
  - [x] /prefs/blocked
  - [x] /prefs/messaging
  - [x] /prefs/trusted
  - [x] /api/v1/me/friends
  - [x] /api/v1/me/blocked


## Captcha

[Api Refernce]()
- [ ] GET  /api/needs_captcha ( any ) 

## Collections
- [ ] POST /api/v1/collections/add_post_to_collection           ( modposts )
- [ ] GET  /api/v1/collections/collection                       ( read )
- [ ] POST /api/v1/collections/create_collection                ( modposts )
- [ ] POST /api/v1/collections/delete_collection                ( modposts )
- [ ] POST /api/v1/collections/follow_collection                ( subscribe )
- [ ] POST /api/v1/collections/reorder_collection               ( modposts )
- [ ] GET  /api/v1/collections/subreddit_collections            ( read )
- [ ] POST /api/v1/collections/update_collection_description    ( modposts )
- [ ] POST /api/v1/collections/update_collection_display_layout ( modposts )
- [ ] POST /api/v1/collections/update_collection_title          ( modposts )

## Emoji
- [ ] POST   /api/v1/subreddit/emoji.json                       ( structuredstyles )
- [ ] DELETE /api/v1/subreddit/emoji/emoji_name                 ( structuredstyles )
- [ ] POST   /api/v1/subreddit/emoji_asset_upload_s3.json       ( structuredstyles )
- [ ] POST   /api/v1/subreddit/emoji_custom_size                ( structuredstyles )
- [ ] GET    /api/v1/subreddit/emojis/all                       ( read )

## Flair
- [x] POST  [/r/subreddit]/api/clearflairtemplates  ( modflair )
- [x] POST  [/r/subreddit]/api/deleteflair          ( modflair )
- [x] POST  [/r/subreddit]/api/deleteflairtemplate  ( modflair )
- [x] POST  [/r/subreddit]/api/flair                ( modflair )
- [ ] PATCH [/r/subreddit]/api/flair_template_order ( modflair )
- [x] POST  [/r/subreddit]/api/flairconfig          ( modflair )
- [x] POST  [/r/subreddit]/api/flaircsv             ( modflair )
- [x] GET   [/r/subreddit]/api/flairlist            ( modflair )
- [x] POST  [/r/subreddit]/api/flairselector        ( flair )
- [x] POST  [/r/subreddit]/api/flairtemplate        ( modflair )
- [x] POST  [/r/subreddit]/api/flairtemplate_v2     ( modflair )
- [x] GET   [/r/subreddit]/api/link_flair           ( flair )
- [x] GET   [/r/subreddit]/api/link_flair_v2        ( flair )
- [x] POST  [/r/subreddit]/api/selectflair          ( flair )
- [x] POST  [/r/subreddit]/api/setflairenabled      ( flair )
- [x] GET   [/r/subreddit]/api/user_flair           ( flair )
- [x] GET   [/r/subreddit ]/api/user_flair_v2       ( flair )

## Reddit Gold
- [x] POST  /api/v1/gold/gild/fullname ( creddits ) 
- [x] POST  /api/v1/gold/give/username ( creddits ) 


## Links & Comments
- [x] POST /api/comment              ( any )
- [x] POST /api/del                  ( edit )
- [x] POST /api/editusertext         ( edit )
- [x] POST /api/event_post_time      ( modposts )
- [x] POST /api/follow_post          ( subscribe )
- [x] POST /api/hide                 ( report )
- [x] GET  [/r/subreddit]/api/info   ( read )
- [x] POST /api/lock                 ( modposts )
- [x] POST /api/marknsfw             ( modposts )
- [x] GET  /api/morechildren         ( read )
- [x] POST /api/report               ( report )
- [x] POST /api/report_award         ( report )
- [x] POST /api/save                 ( save )
- [x] GET  /api/saved_categories     ( save )
- [x] POST /api/sendreplies          ( edit )
- [x] POST /api/set_contest_mode     ( modposts )
- [x] POST /api/set_subreddit_sticky ( modposts )
- [x] POST /api/set_suggested_sort   ( modposts )
- [x] POST /api/spoiler              ( modposts )
- [x] POST /api/store_visits         ( save )
- [x] POST /api/submit               ( submit )
- [x] POST /api/unhide               ( report )
- [x] POST /api/unlock               ( modposts )
- [x] POST /api/unmarknsfw           ( modposts )
- [x] POST /api/unsave               ( save )
- [x] POST /api/unspoiler            ( modposts )
- [x] POST /api/vote                 ( vote )


## Listings
- [x] GET  /api/trending_subreddits 
- [x] GET  /best                           ( read ) ( rss support )
- [x] GET  /by_id/names                    ( read )
- [x] GET  [/r/subreddit]/comments/article ( read ) ( rss support )
- [x] GET  /duplicates/article             ( read ) ( rss support )
- [x] GET  [/r/subreddit]/hot              ( read ) ( rss support )
- [x] GET  [/r/subreddit]/new              ( read ) ( rss support )
- [x] GET  [/r/subreddit]/random           ( read )
- [x] GET  [/r/subreddit]/rising           ( read ) ( rss support )
- [x] GET  [/r/subreddit]/sort             ( read ) ( rss support )
  - [x]  [/r/subreddit]/top
  - [x]  [/r/subreddit]/controversial 


## Live Threads
- [ ] GET  /api/live/by_id/names                        ( read )
- [ ] POST /api/live/create                             ( submit )
- [ ] GET  /api/live/happening_now                      ( read )
- [ ] POST /api/live/thread/accept_contributor_invite   ( livemanage )
- [ ] POST /api/live/thread/close_thread                ( livemanage )
- [ ] POST /api/live/thread/delete_update               ( edit )
- [ ] POST /api/live/thread/edit                        ( livemanage )
- [ ] POST /api/live/thread/hide_discussion             ( livemanage )
- [ ] POST /api/live/thread/invite_contributor          ( livemanage )
- [ ] POST /api/live/thread/leave_contributor           ( livemanage )
- [ ] POST /api/live/thread/report                      ( report )
- [ ] POST /api/live/thread/rm_contributor              ( livemanage )
- [ ] POST /api/live/thread/rm_contributor_invite       ( livemanage )
- [ ] POST /api/live/thread/set_contributor_permissions ( livemanage )
- [ ] POST /api/live/thread/strike_update               ( edit )
- [ ] POST /api/live/thread/unhide_discussion           ( livemanage )
- [ ] POST /api/live/thread/update                      ( submit )
- [ ] GET  /live/thread                                 ( read ) ( rss support  )
- [ ] GET  /live/thread/about                           ( read )
- [ ] GET  /live/thread/contributors                    ( read )
- [ ] GET  /live/thread/discussions read                ( rss support  )
- [ ] GET  /live/thread/updates/update_id               ( read )


## Private Messages
- [ ] POST /api/block              ( privatemessages )
- [ ] POST /api/collapse_message   ( privatemessages )
- [ ] POST /api/compose            ( privatemessages )
- [ ] POST /api/del_msg            ( privatemessages )
- [ ] POST /api/read_all_messages  ( privatemessages )
- [ ] POST /api/read_message       ( privatemessages )
- [ ] POST /api/unblock_subreddit  ( privatemessages )
- [ ] POST /api/uncollapse_message ( privatemessages )
- [ ] POST /api/unread_message     ( privatemessages )
- [ ] GET  /message/where          ( privatemessages ) ( rss support )
  - [ ] /message/inbox
  - [ ] /message/unread
  - [ ] /message/sent


## Misc
- [ ] GET [/r/subreddit ]/api/saved_media_text ( submit )
- [ ] GET /api/v1/scopes                       ( any )


## Moderation
- [x] GET [/r/subreddit]/about/log modlog              ( rss support )
- [x] GET [/r/subreddit]/about/location                ( read )
  - [x] [/r/subreddit]/about/reports
  - [x] [/r/subreddit]/about/spam
  - [x] [/r/subreddit]/about/modqueue
  - [x] [/r/subreddit]/about/unmoderated
  - [x] [/r/subreddit]/about/edited
- [x] POST [/r/subreddit ]/api/accept_moderator_invite ( modself )
- [x] POST /api/approve                                ( modposts )
- [x] POST /api/distinguish                            ( modposts )
- [x] POST /api/ignore_reports                         ( modposts )
- [x] POST /api/leavecontributor                       ( modself )
- [x] POST /api/leavemoderator                         ( modself )
- [x] POST /api/mute_message_author                    ( modcontributors )
- [x] POST /api/remove                                 ( modposts )
- [x] POST /api/show_comment                           ( modposts )
- [x] POST /api/unignore_reports                       ( modposts )
- [x] POST /api/unmute_message_author                  ( modcontributors )
- [x] GET  [/r/subreddit ]/stylesheet                  ( modconfig )


## New Modmail
- [ ] POST   /api/mod/bulk_read                                ( modmail )
- [ ] GET    /api/mod/conversations                            ( modmail )
- [ ] POST   /api/mod/conversations                            ( modmail )
- [ ] GET    /api/mod/conversations/:conversation_id           ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id           ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id/archive   ( modmail )
- [ ] DELETE /api/mod/conversations/:conversation_id/highlight ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id/highlight ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id/mute      ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id/unarchive ( modmail )
- [ ] POST   /api/mod/conversations/:conversation_id/unmute    ( modmail )
- [ ] GET    /api/mod/conversations/:conversation_id/user      ( modmail )
- [ ] POST   /api/mod/conversations/read                       ( modmail )
- [ ] GET    /api/mod/conversations/subreddits                 ( modmail )
- [ ] POST   /api/mod/conversations/unread                     ( modmail )
- [ ] GET    /api/mod/conversations/unread/count               ( modmail )


## Multis
- [ ] POST   /api/multi/copy                  ( subscribe )
- [ ] GET    /api/multi/mine                  ( read )
- [ ] GET    /api/multi/user/username         ( read )
- [ ] DELETE /api/multi/multipath             ( subscribe )
  - [ ] /api/filter/filterpath
- [ ] GET    /api/multi/multipath             ( read )
  - [ ] /api/filter/filterpath
- [ ] POST   /api/multi/multipath             ( subscribe )
  - [ ] /api/filter/filterpath
- [ ] PUT    /api/multi/multipath             ( subscribe )
  - [ ] /api/filter/filterpath
- [ ] GET    /api/multi/multipath/description ( read )
- [ ] PUT    /api/multi/multipath/description ( read )
- [ ] DELETE /api/multi/multipath/r/srname    ( subscribe )
  - [ ] /api/filter/filterpath
- [ ] GET    /api/multi/multipath/r/srname    ( read )
  - [ ] /api/filter/filterpath
- [ ] PUT    /api/multi/multipath/r/srname    ( subscribe )
  - [ ] /api/filter/filterpath


## Search
- [ ] GET [/r/subreddit]/search             ( read ) ( rss support ) 


## Subreddits
- [x] GET [/r/subreddit]/about/where ( read ) ( rss support ) 
  - [x] [/r/subreddit]/about/banned
  - [x] [/r/subreddit]/about/muted
  - [x] [/r/subreddit]/about/wikibanned
  - [x] [/r/subreddit]/about/contributors
  - [x] [/r/subreddit]/about/wikicontributors
  - [x] [/r/subreddit]/about/moderators  
- [x] POST [/r/subreddit]/api/delete_sr_banner ( modconfig )
- [x] POST [/r/subreddit]/api/delete_sr_header ( modconfig )
- [x] POST [/r/subreddit]/api/delete_sr_icon   ( modconfig )
- [x] POST [/r/subreddit]/api/delete_sr_img    ( modconfig )
- [x] GET  /api/recommend/sr/srnames           ( read )
- [x] GET  /api/search_reddit_names            ( read )
- [x] POST /api/search_reddit_names            ( read )
- [x] POST /api/search_subreddits              ( read )
- [x] POST /api/site_admin                         ( modconfig )
- [x] GET  [/r/subreddit]/api/submit_text          ( submit )
- [x] GET  /api/subreddit_autocomplete             ( read )
- [x] GET  /api/subreddit_autocomplete_v2          ( read )
- [x] POST [/r/subreddit]/api/subreddit_stylesheet ( modconfig )
- [x] POST /api/subscribe                          ( subscribe )
- [x] POST [/r/subreddit]/api/upload_sr_img        ( modconfig )
- [x] GET  /api/v1/subreddit/post_requirements     ( submit )
- [ ] GET  /r/subreddit/about                      ( read )
- [ ] GET  /r/subreddit/about/edit                 ( modconfig )
- [ ] GET  /r/subreddit/about/rules                ( read )
- [ ] GET  /r/subreddit/about/traffic              ( modconfig )
- [ ] GET  [/r/subreddit]/si
- [ ] GET  [/r/subreddit]/sticky                   ( read )
- [ ] GET  /subreddits/mine/where                  ( mysubreddits ) ( rss support )
  - [ ] /subreddits/mine/subscriber
  - [ ] /subreddits/mine/contributor
  - [ ] /subreddits/mine/moderator
  - [ ] /subreddits/mine/streams
- [ ] GET /subreddits/search                       ( read ) ( rss support )
- [ ] GET /subreddits/where                        ( read ) ( rss support )
  - [ ] /subreddits/popular
  - [ ] /subreddits/new
  - [ ] /subreddits/gold
  - [ ] /subreddits/default
- [ ] GET /users/search                            ( read ) ( rss support ) 
- [ ] GET /users/where                             ( read ) ( rss support ) 
  - [ ] /users/popular
  - [ ] /users/new


## Users
- [ ] POST   /api/block_user                   ( account )
- [ ] POST   [/r/subreddit]/api/friend         ( any )
- [ ] POST   /api/report_user                  ( report )
- [ ] POST   [/r/subreddit]/api/setpermissions ( modothers )
- [ ] POST   [/r/subreddit]/api/unfriend       ( any )
- [ ] GET    /api/user_data_by_account_ids     ( privatemessages )
- [ ] GET    /api/username_available           ( any )
- [ ] DELETE /api/v1/me/friends/username       ( subscribe )
- [ ] GET    /api/v1/me/friends/username       ( mysubreddits )
- [ ] PUT    /api/v1/me/friends/username       ( subscribe )
- [ ] GET    /api/v1/user/username /trophies   ( read )
- [ ] GET    /user/username/about              ( read )
- [ ] GET    /user/username/where history      ( rss support )
  - [ ] /user/username/overview
  - [ ] /user/username/submitted
  - [ ] /user/username/comments
  - [ ] /user/username/upvoted
  - [ ] /user/username/downvoted
  - [ ] /user/username/hidden
  - [ ] /user/username/saved
  - [ ] /user/username/gilded


## Widgets
- [ ] POST   /api/widget                 ( structuredstyles )
- [ ] DELETE /api/widget/widget_id       ( structuredstyles )
- [ ] PUT    /api/widget/widget_id       ( structuredstyles )
- [ ] POST   /api/widget_image_upload_s3 ( structuredstyles )
- [ ] PATCH  /api/widget_order/section   ( structuredstyles )
- [ ] GET    /api/widgets                ( structuredstyles )


## Wiki
- [ ] POST [/r/subreddit]/api/wiki/alloweditor/act ( modwiki )
  - [ ] [/r/subreddit]/api/wiki/alloweditor/       ( del )
  - [ ] [/r/subreddit]/api/wiki/alloweditor/       ( add )
- [ ] POST [/r/subreddit]/api/wiki/edit            ( wikiedit )
- [ ] POST [/r/subreddit]/api/wiki/hide            ( modwiki )
- [ ] POST [/r/subreddit]/api/wiki/revert          ( modwiki )
- [ ] GET  [/r/subreddit]/wiki/discussions/page    ( wikiread )
- [ ] GET  [/r/subreddit]/wiki/pages               ( wikiread )
- [ ] GET  [/r/subreddit]/wiki/revisions           ( wikiread )
- [ ] GET  [/r/subreddit]/wiki/revisions/page      ( wikiread )
- [ ] GET  [/r/subreddit]/wiki/settings/page       ( modwiki )
- [ ] POST [/r/subreddit]/wiki/settings/page       ( modwiki )
-
