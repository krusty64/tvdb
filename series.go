package tvdb

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
)

type Series struct {
	SeriesId int `xml:"seriesid"`
	Language string `xml:"language"`
	SeriesName string
}

func GetSeriesURL(seriesname, language string) (string, error) {
	query, err := url.Parse(TVDB_API)
	if err != nil {
		return "", err
	}

	if query, err = query.Parse("GetSeries.php"); err != nil {
		return "", err
	}

	args := url.Values{}
	args.Add("seriesname", seriesname)
	args.Add("language", language)
	query.RawQuery = args.Encode()

	return query.String(), nil
}

func GetSeries(seriesname, language string) (result []*Series, err error) {
	var query string
	if query, err = GetSeriesURL(seriesname, language); err != nil {
		return
	}
	var resp http.Response
	if resp, err = http.Get(query); err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

	return nil, nil
}
