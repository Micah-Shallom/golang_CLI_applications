package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const timeFormat = "Jan/02 @15:04"

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type response struct {
	Results      []item `json:"results"`
	Date         int    `json:"date"`
	TotalResults int    `json:"total_results"`
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
		return nil, fmt.Errorf("%w:%s", ErrConnection,err)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		msg, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("cannot read body: %w", err)
		}
		err = ErrInvalidResponse
		if r.StatusCode == http.StatusNotFound {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("%w: %s", err, msg)
	}
	var resp response

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}
	if resp.TotalResults == 0 {
		return nil, fmt.Errorf("%w: No results found", ErrNotFound)
	}
	return resp.Results, nil
}

func getAll(apiRoot string)([]item, error){
	u := fmt.Sprintf("%s/todo", apiRoot)
	return getItems(u)
}

func getOne(apiRoot string, id int) (item,error){
	u := fmt.Sprintf("%s/todo/%d",apiRoot,id)
	items, err := getItems(u)
	if err != nil {
		return item{}, err
	}
	if len(items) != 1 {
		return item{}, fmt.Errorf("%w: Invalid results", ErrInvalid)
	}
	return items[0], nil 
}

