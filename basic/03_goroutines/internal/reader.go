package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type EntriesReader interface {
	Read(word string) (Entries, error)
}

type reader struct {
	client http.Client
}

func (r reader) Read(word string) (Entries, error) {
	res, err := r.client.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response error %d", res.StatusCode)
	}

	entries := Entries{}
	err = json.Unmarshal(b, &entries)
	return entries, err
}

// NewEntriesReader create new instance of entries reader
func NewEntriesReader() EntriesReader {
	return &reader{
		http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
