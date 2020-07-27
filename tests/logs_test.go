package tests

import (
	"github.com/stretchr/testify/assert"
	"metrics-service/server/requests"
	"metrics-service/tests/helpers"
	"net/http"
	"testing"
)

const firstUserIp = "83.150.59.250"
const secondUserIp = "83.150.59.251"

func TestWalkParseLogsApi(t *testing.T) {
	cases := []helpers.TestCase{
		{
			"Save log message success",
			helpers.Request{
				Method: http.MethodPost,
				Url:    "/logs",
			},
			requests.SaveLogRequest{
				Timestamp: "2020-06-24T15:27:00.123456Z",
				IP:  	   firstUserIp,
				Url: 	   "https://example.com/homepage",
			},
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "Log message has been saved",
			},
			1,
		},
		{
			"Check IPs count increase for different IPs",
			helpers.Request{
				Method: http.MethodPost,
				Url:    "/logs",
			},
			requests.SaveLogRequest{
				Timestamp: "2020-06-24T15:27:00.123456Z",
				IP:  	   secondUserIp,
				Url: 	   "https://example.com/homepage",
			},
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "Log message has been saved",
			},
			2,
		},
		{
			"Check unique IPs count not changed for repeated IPs",
			helpers.Request{
				Method: http.MethodPost,
				Url:    "/logs",
			},
			requests.SaveLogRequest{
				Timestamp: "2020-06-24T15:27:00.123456Z",
				IP:  	   secondUserIp,
				Url: 	   "https://example.com/homepage",
			},
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "Log message has been saved",
			},
			2,
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			recorder := helpers.PrepareServerFromTestCase(s, test)

			assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			if assert.Equal(t, test.Expected.StatusCode, recorder.Code) {
				if recorder.Code == http.StatusOK {
					helpers.AssertIPsCount(t, s, test.ExpectedIPsCount)
				}
			}
		})
	}
}
