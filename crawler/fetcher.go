package crawler

import (
	"errors"
	"net/http"
)

// Get html response from a url
// return only responses with a valid html body
func fetchHtml(url string) (response *http.Response, err error)  {
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("unable to fetch data")
	}
	return resp, nil
}
