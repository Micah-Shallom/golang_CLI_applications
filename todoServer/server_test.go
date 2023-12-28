package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Micah-Shallom/golang_CLI_applications/todo"
)

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()

	tempTodoFile, err := os.CreateTemp("", "todotest")
	if err != nil {
		t.Fatal(err)
	}

	//this is exactly why a newMux function was created seperately....so that it can be reused during testing
	ts := httptest.NewServer(newMux(tempTodoFile.Name()))

	// Adding a couple of items for testing
	for i := 1; i < 3; i++ {
		var body bytes.Buffer
		taskName := fmt.Sprintf("Task number %d.", i)
		item := struct {
			Task string `json:"task"`
		}{
			Task: taskName,
		}
		if err := json.NewEncoder(&body).Encode(item); err != nil {
			t.Fatal(err)
		}
		r, err := http.Post(ts.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatal(err)
		}
		if r.StatusCode != http.StatusCreated {
			t.Fatalf("Failed to add initial items: Status: %d", r.StatusCode)
		}
	}
	return ts.URL, func() {
		ts.Close()
		os.Remove(tempTodoFile.Name())
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
		{name: "GetAll", path: "/todo",
			expCode:    http.StatusOK,
			expItems:   2,
			expContent: "Task number 1.",
		},
		{name: "GetOne", path: "/todo/1",
			expCode:    http.StatusOK,
			expItems:   1,
			expContent: "Task number 1.",
		},
		{name: "NotFound", path: "/todo/500",
			expCode: http.StatusNotFound,
		},
	}

	url, cleanup := setupAPI(t)
	defer cleanup()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				resp struct {
					Results      todo.List `json:"results"`
					Date         int64     `json:"date"`
					TotalResults int       `json:"total_results"`
				}
				body []byte
				err  error
			)
			if resp.TotalResults != tc.expItems {
				t.Errorf("Expected %d items, got %d.", tc.expItems, resp.TotalResults)
			}
			if resp.Results[0].Task != tc.expContent {
				t.Errorf("Expected %q, got %q.", tc.expContent, resp.Results[0].Task)
			}
			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Error(err)
			}
			defer r.Body.Close()
			if r.StatusCode != tc.expCode {
				t.Fatalf("Expected %q, got %q.", http.StatusText(tc.expCode), http.StatusText(r.StatusCode))
			}

			switch {
			case r.Header.Get("Content-Type") == "application/json":
				if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
					t.Error(err)
				}

			case strings.Contains(r.Header.Get("Content-Type"), "text/plain"):
				if body, err = io.ReadAll(r.Body); err != nil {
					t.Error(err)
				}
				if !strings.Contains(string(body), tc.expContent) {
					t.Errorf("Expected %q, got %q.", tc.expContent, string(body))
				}
			default:
				t.Fatalf("Unsupported Content-Type: %q", r.Header.Get("Content-Type"))
			}
		})
	}
}

func TestAdd(t *testing.T) {
	url, cleanup := setupAPI(t)
	defer cleanup()

	taskName := "Task number 3."
	t.Run("Add", func(t *testing.T) {
		var body bytes.Buffer
		item := struct {
			Task string `json:"task"`
		}{
			Task: taskName,
		}
		if err := json.NewEncoder(&body).Encode(item); err != nil {
			t.Fatal(err)
		}
		r, err := http.Post(url+"/todo", "application/json", &body)
		if err != nil {
			t.Fatal(err)
		}
		if r.StatusCode != http.StatusCreated {
			t.Errorf("Expected %q, got %q.", http.StatusText(http.StatusCreated), http.StatusText(r.StatusCode))
		}
	})

	t.Run("CheckAdd", func(t *testing.T) {
		r, err := http.Get(url + "/todo/3")
		if err != nil {
			t.Error(err)
		}
		if r.StatusCode != http.StatusOK {
			t.Fatalf("Expected %q, got %q.", http.StatusText(http.StatusOK), http.StatusText(r.StatusCode))
		}

		var resp todoResponse
		if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
			t.Fatal(err)
		}
		r.Body.Close()

		if resp.Results[0].Task != taskName {
			t.Errorf("Exxpected %q, got %q.", taskName, resp.Results[0].Task)
		}
	})
}

func TestDelete(t *testing.T) {
	url, cleanup := setupAPI(t)
	defer cleanup()

	t.Run("Delete", func(t *testing.T) {
		u := fmt.Sprintf("%s/todo/1", url)
		req, err := http.NewRequest(http.MethodDelete, u, nil)
		if err != nil {
			t.Fatal(err)
		}
		r, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if r.StatusCode != http.StatusNoContent {
			t.Fatalf("Expected %q, got %q.", http.StatusText(http.StatusNoContent), http.StatusText(r.StatusCode))
		}
	})

	t.Run("CheckDelete", func(t *testing.T) {
		r, err := http.Get(url + "/todo")
		if err != nil {
			t.Error(err)
		}
		if r.StatusCode != http.StatusOK {
			t.Fatalf("Expected %q, got %q.", http.StatusText(http.StatusOK), http.StatusText(r.StatusCode))
		}

		var resp todoResponse
		if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
			t.Fatal(err)
		}
		r.Body.Close()
		if len(resp.Results) != 1 {
			t.Errorf("Expected 1 item, got %d.", len(resp.Results))
		}
		expTask := "Task number 2."
		if resp.Results[0].Task != expTask {
			t.Errorf("Expected %q, got %q.", expTask, resp.Results[0].Task)
		}
	})
}

func TestComplete(t *testing.T) {
	url, cleanup := setupAPI(t)
	defer cleanup()
	t.Run("Complete", func(t *testing.T) {
		u := fmt.Sprintf("%s/todo/1?complete", url)
		req, err := http.NewRequest(http.MethodPatch, u, nil)
		if err != nil {
			t.Fatal(err)
		}
		r, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if r.StatusCode != http.StatusNoContent {
			t.Fatalf("Expected %q, got %q.",
				http.StatusText(http.StatusNoContent), http.StatusText(r.StatusCode))
		}
		t.Run("CheckComplete", func(t *testing.T) {
			r, err := http.Get(url + "/todo")
			if err != nil {
				t.Error(err)
			}
			if r.StatusCode != http.StatusOK {
				t.Fatalf("Expected %q, got %q.",
					http.StatusText(http.StatusOK), http.StatusText(r.StatusCode))
			}
			var resp todoResponse
			if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
				t.Fatal(err)
			}
			r.Body.Close()
			if len(resp.Results) != 2 {
				t.Errorf("Expected 2 items, got %d.", len(resp.Results))
			}
			if !resp.Results[0].Done {
				t.Error("Expected Item 1 to be completed")
			}
			if resp.Results[1].Done {
				t.Error("Expected Item 2 not to be completed")
			}
		})
	})
}
