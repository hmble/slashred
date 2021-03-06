package slashred

import "log"

// Test remained
// Its may be not used by many user so if a user wants to report then
// he should follow api link for data fields
// https://www.reddit.com/dev/api/#POST_api_report
type ReportService service

func (r *ReportService) Report(postdata PostData) {
	resp, err := r.client.Post(API_PATH["report"], postdata)

	if err != nil {
		log.Fatal("Error in reporting")
	}

	defer resp.Body.Close()
}

func (r *ReportService) ReportAward(awardId, reason string) {
	postdata := PostData{

		"award_id": awardId,
		"reason":   reason,
	}

	resp, err := r.client.Post(API_PATH["report_award"], postdata)

	if err != nil {
		log.Fatal("Error in reporting")
	}

	defer resp.Body.Close()
}
