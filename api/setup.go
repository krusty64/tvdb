package api

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

const (
	TVDB_API = "http://thetvdb.com/api/"
	TVDB_DEFAULT_KEY = "CACE3A94B49F1566"
)

type TVDB struct {
	Location string
	ApiKey string
}

func Open() *TVDB {
	return &TVDB{
		Location: TVDB_API,
		ApiKey: TVDB_DEFAULT_KEY,
	}
}

func (t *TVDB) GetURL(path string, args *url.Values) (string, error) {
	query, err := url.Parse(t.Location)
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

	fmt.Println(string(body))

	return body, nil
}

func (t *TVDB) QueryURL(path string, args *url.Values) ([]byte, error) {
	query, err := t.GetURL(path, args)
	if err != nil {
		return nil, err
	}
	return HttpGet(query)
}

func (t *TVDB) QueryAndUnmarshal(path string, args *url.Values, result interface{}) error {
	body, err := t.QueryURL(path, args)
	if err != nil {
		return err
	}
    return xml.Unmarshal(body, result)
}
