package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req),nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {

	return &http.Client{
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {

	jsonToReturn := `	
	{
		"error":false,
		"message: "some message"
	}
	`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email":    "me@here.com",
		"password": "verysercret",
	}

	body, err := json.Marshal(postBody)

	if err != nil {

		fmt.Println("Error: %w", err)
		t.Errorf("error parsing json")

	}

	fmt.Println("before creating request")
	req, err := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))

	fmt.Println("after creating request")
	if err != nil {
		fmt.Println("Error: %w", err)
		t.Errorf("error creating new request")

	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	fmt.Println("before serve http")
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("expected http Status accepted but got %d", rr.Code)
	}

}
