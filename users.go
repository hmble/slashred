package slashred

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type UsersService service

func (u *UsersService) UsersSearch(opts Option) {
	logmsg := "Error in searching users"

	u.client.getlisting(API_PATH["users_search"], logmsg, opts)
}
func (u *UsersService) UsersPopular(opts Option) {
	logmsg := "Error in getting popular users"

	u.client.getlisting(API_PATH["users_popular"], logmsg, opts)
}
func (u *UsersService) UsersNew(opts Option) {
	logmsg := "Error in getting new users"

	u.client.getlisting(API_PATH["users_new"], logmsg, opts)
}

// accountID fullname
func (u *UsersService) Block(accountID, name string) {
	postdata := PostData{
		"name":       name,
		"account_id": accountID,
	}

	resp, err := u.client.Post(API_PATH["block_user"], postdata)

	if err != nil {
		log.Fatal("Error in blocking users")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (u *UsersService) Friend(banContext, message, reason, duration, name, note, reltnType string) {
	postdata := PostData{
		"ban_context": banContext,
		"ban_message": message,
		"ban_reason":  reason,
		"duration":    duration,
		"name":        name,
		"note":        note,
		"type":        reltnType,
	}

	resp, err := u.client.Post(API_PATH["friend"], postdata)

	if err != nil {
		log.Fatal("Error in getting friend")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

// func (u *UsersService) ReportUser(details, reason, user string) {
// 	postdata := PostData{
// 		"details": details,
// 		"reason":  reason,
// 		"user":    user,
// 	}

// 	resp, err := u.client.Post(API_PATH["report_user"], postdata)

// 	if err != nil {
// 		log.Fatal("Error in reporting user")
// 	}

// 	defer resp.Body.Close()
// 	PrintHeader(resp)
// }

func (u *UsersService) Unfriend(subreddit, id, name, reltnType string) {
	postdata := PostData{
		"id":   id,
		"name": name,
		"type": reltnType,
	}

	endpoint := fmt.Sprintf("/r/%s/%s", subreddit, API_PATH["unfriend"])
	resp, err := u.client.Post(endpoint, postdata)

	if err != nil {
		log.Fatal("Error in doing unfriend to user")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (u *UsersService) UserDataByAccountIds(ids []string) {
	idstring := strings.Join(ids, ",")
	opts := Option{
		"ids": idstring,
	}
	resp, err := u.client.Get(API_PATH["user_data_by_account_ids"], opts)

	if err != nil {
		log.Fatal("Error in getting users data for given ids")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	SaveResponse(resp.Body, "test_data/userdata_byid.json")
}

func (u *UsersService) UsernameAvailable(username string) {
	opts := Option{
		"user": username,
	}
	resp, err := u.client.Get(API_PATH["username_available"], opts)

	if err != nil {
		log.Fatal("Error in checking username available or not")
	}

	defer resp.Body.Close()

	PrintHeader(resp)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func (u *UsersService) DeleteFriend(ownusername, id string) {
	opts := Option{
		"id": id,
	}

	endpoint := fmt.Sprintf("%s/%s", API_PATH["friend_v1"], ownusername)

	resp, err := u.client.Delete(endpoint, opts)

	if err != nil {
		log.Fatal("Error in deleting username")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (u *UsersService) FriendInfo(ownusername, id string) {
	opts := Option{
		"id": id,
	}

	endpoint := fmt.Sprintf("%s/%s", API_PATH["friend_v1"], ownusername)

	resp, err := u.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal("Error in deleting username")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
}

func (u *UsersService) AddFriend(ownusername, name string) {

	//data := fmt.Sprintf(`{"name": "%s", "note": "%s"}`, name, note)
	data := fmt.Sprintf(`{"name": "%s"}`, name)

	fmt.Println(data)
	endpoint := fmt.Sprintf("%s/%s", API_PATH["friend_v1"], ownusername)
	fmt.Println("------url: ", endpoint)
	resp, err := u.client.Put(endpoint, data)

	if err != nil {
		log.Fatal("Error in making friends")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}

func (u *UsersService) GetUserTrophies(ownusername, id string) {
	opts := Option{
		"id": id,
	}

	endpoint := fmt.Sprintf("%s/%s/trophies", API_PATH["get_user_trophies"], ownusername)
	resp, err := u.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal("Error in getting users trophies")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func (u *UsersService) About(ownusername string) {
	opts := Option{
		"username": ownusername,
	}

	endpoint := fmt.Sprintf("/user/%s/about", ownusername)

	resp, err := u.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal("Error in getting own about info")
	}

	defer resp.Body.Close()

	PrintHeader(resp)
	SaveResponse(resp.Body, "test_data/workpcinfo.json")
}

func (u *UsersService) getuserwhere(where, username, logmsg string, opts Option) {
	endpoint := fmt.Sprintf("/user/%s/%s", username, where)

	resp, err := u.client.Get(endpoint, opts)

	if err != nil {
		log.Fatal(logmsg)
	}

	defer resp.Body.Close()
	PrintHeader(resp)

}

func (u *UsersService) Overview(ownusername string, opts Option) {
	logmsg := "Error in getting overview response"
	u.getuserwhere("overview", ownusername, logmsg, opts)
}

func (u *UsersService) Submitted(ownusername string, opts Option) {
	logmsg := "Error in getting submittedresponse"
	u.getuserwhere("submitted", ownusername, logmsg, opts)
}

func (u *UsersService) Comments(ownusername string, opts Option) {
	logmsg := "Error in getting comments response"
	u.getuserwhere("comments", ownusername, logmsg, opts)
}

func (u *UsersService) Upvoted(ownusername string, opts Option) {
	logmsg := "Error in getting upvoted response"
	u.getuserwhere("upvoted", ownusername, logmsg, opts)
}

func (u *UsersService) Downvoted(ownusername string, opts Option) {
	logmsg := "Error in getting downvoted response"
	u.getuserwhere("downvoted", ownusername, logmsg, opts)
}

func (u *UsersService) Hidden(ownusername string, opts Option) {
	logmsg := "Error in getting hidden response"
	u.getuserwhere("hidden", ownusername, logmsg, opts)
}

func (u *UsersService) Saved(ownusername string, opts Option) {
	logmsg := "Error in getting saved response"
	u.getuserwhere("saved", ownusername, logmsg, opts)
}
func (u *UsersService) Gilded(ownusername string, opts Option) {
	logmsg := "Error in getting gilded response"
	u.getuserwhere("gilded", ownusername, logmsg, opts)
}

