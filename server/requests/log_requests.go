package requests

type SaveLogRequest struct {
	Timestamp string `json:"timestamp" example:"2020-06-24T15:27:00.123456Z"`
	IP  	  string `json:"ip" example:"83.150.59.250"`
	Url 	  string `json:"url" example:"https://example.com/homepage"`
}
