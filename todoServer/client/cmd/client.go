package cmd

import (
	errors "github.com/Micah-Shallom/golang_cli_applications/todoServer/todoClient"
	"fmt"
	"net/http"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type response struct {
	Results      []string `json:"results"`
	Date         int      `json:"date"`
	TotalResults int      `json:"total_results"`
}

func newClient() *http.Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	return c
}

func getItems(url string) ([]item, error) {
	r, err := newClient().Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w:%s", errors.ErrConnection)
	}
}