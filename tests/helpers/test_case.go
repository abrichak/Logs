package helpers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"metrics-service/server"
	"net/http"
	"net/http/httptest"
	"strings"
)

type TestCase struct{
	TestName    	 string
	Request     	 Request
	RequestBody 	 interface{}
	HandlerFunc 	 func(s *server.Server, c echo.Context) error
	Expected    	 ExpectedResponse
	ExpectedIPsCount int
}

type Request struct {
	Method    string
	Url       string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

func PrepareServerFromTestCase(s *server.Server, test TestCase) (recorder *httptest.ResponseRecorder) {
	request, recorder := prepareTestCase(test)
	s.Echo.ServeHTTP(recorder, request)

	return recorder
}

func prepareTestCase(test TestCase) (request *http.Request, recorder *httptest.ResponseRecorder) {
	requestJson, _ := json.Marshal(test.RequestBody)
	request = httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(string(requestJson)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return request, httptest.NewRecorder()
}