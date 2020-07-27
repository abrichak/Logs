package tests

import (
	"github.com/stretchr/testify/assert"
	"metrics-service/server/requests"
	"metrics-service/tests/helpers"
	"net/http"
	"testing"
)

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
				IP:  	   "83.150.59.250",
				Url: 	   "https://example.com/homepage",
			},
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "Log message has been saved",
			},
		},
	}

	s := helpers.NewServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			recorder := helpers.PrepareServerFromTestCase(s, test)

			assert.Equal(t, test.Expected.StatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
		})
	}
}
