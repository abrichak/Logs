package requests

type SaveLogRequest struct {
	Timestamp string `json:"timestamp"`
	IP  	  string `json:"ip"`
	Url 	  string `json:"url"`
}
