package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {

	tests := []struct {
		name  string
		verb  string
		route string
		want  int
	}{
		{"POST not allowed", "POST", "/word", 405},
		{"PUT not allowed", "PUT", "/word", 405},
		{"DELETE not allowed", "DELETE", "/word", 405},
		{"PATCH not allowed", "PATCH", "/word", 405},
		{"POST not allowed", "HEAD", "/word", 405},
		{"OPTIONS not allowed", "OPTIONS", "/word", 405},
		{"Test no route", "GET", "/", 404},
	}

	r := SetupRouter()

	for _, tt := range tests {

		request, err := http.NewRequest(tt.verb, tt.route, nil)
		if err != nil {
			t.Errorf("%s request to %s failed", tt.verb, tt.route)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)
		got := w.Code

		t.Run(tt.name, func(t *testing.T) {
			if err != nil {
				t.Errorf("Request with verb %s failed", tt.verb)
			}
			if got != tt.want {
				t.Errorf("Request verb (%s) got %v, want %v", tt.verb, got, tt.want)
			}
		})
	}
}

func TestHelloServer(t *testing.T) {

	tests := []struct {
		name     string
		testcase string
		want     int
	}{
		{"empty string", "/", 404},
		{"string not allowed", "/string!", 400},
		{"valid string", "/word", 200},
	}

	r := SetupRouter()

	for _, tt := range tests {

		request, err := http.NewRequest("GET", tt.testcase, nil)
		if err != nil {
			t.Errorf("%s request to failed", tt.testcase)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)
		got := w.Code

		t.Run(tt.name, func(t *testing.T) {
			if got != tt.want {
				t.Errorf("Request got %v, want %v", got, tt.want)
			}
		})
	}
}
