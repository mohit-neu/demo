package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		output string
	}{
		{
			name:   "User with a name",
			path:   "/joe",
			output: "Hello, joe!",
		},
		{
			name:   "User with no name",
			path:   "/",
			output: "Hello, stranger!",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			th := http.HandlerFunc(handler)

			th.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("unexpected status code: %v", rr.Code)
			}
			if got, want := rr.Body.String(), tc.output; !strings.Contains(got, want) {
				t.Errorf("unexpected body in response.  got: %v, want: %v", got, want)
			}
		})
	}
}
