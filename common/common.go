package common

import (
	"io/ioutil"
	"net/http"
)

// FetchDatas is a basic http body fetch.
func FetchDatas(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		bod, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return bod, nil
	}
}
