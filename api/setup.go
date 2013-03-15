package api

import (
	"net/url"
	"net/http"
	"io/ioutil"
)

const (
	TVDB_API = "http://thetvdb.com/api/"
)

func GetURL(path string, args *url.Values) (string, error) {
	query, err := url.Parse(TVDB_API)
	if err != nil {
		return "", err
	}

	if query, err = query.Parse(path); err != nil {
		return "", err
	}

	if args != nil {
		query.RawQuery = args.Encode()
	}
	return query.String(), nil
}

func HttpGet(query string) ([]byte, error) {
	resp, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
