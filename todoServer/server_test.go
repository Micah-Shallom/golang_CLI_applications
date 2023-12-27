package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()
	ts := httptest.NewServer(newMux(""))
	return ts.URL, func() {
		ts.Close()
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name       string
		path       string
		expCode    int
		expItems   int
		expContent string
	}{
		{name: "GetRoot", path: "/",
			expCode:    http.StatusOK,
			expContent: "There's an API here",
		},
		{name: "NotFound", path: "/todo/500",
			expCode: http.StatusNotFound,
		},
	}

	url, cleanup := setupAPI(t)
	defer cleanup()

	for _, tc := range testCases {
		t.Run(tc.name, func (t *testing.T)  {
			var (
				body []byte
				err error
			)
			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Error(err)
			}
			defer r.Body.Close()
			if r.StatusCode != tc.expCode{
				t.Fatalf("Expected %q, got %q.", http.StatusText(tc.expCode), http.StatusText(r.StatusCode))
			}

			switch {
				case strings.Contains(r.Header.Get("Content-Type"), "text/plain"):
					if body, err = io.ReadAll(r.Body); err != nil {
						t.Error(err)
					} 
					if !strings.Contains(string(body), tc.expContent){
						t.Errorf("Expected %q, got %q.",tc.expContent, string(body))
					}
				default:
					t.Fatalf("Unsupported Content-Type: %q", r.Header.Get("Content-Type"))
			}
		})
	}
}
