package api

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
)

type Series struct {
	SeriesId int `xml:"seriesid"`
	Language string `xml:"language"`
	SeriesName string
	Overview string
	Banner string `xml:"banner"`
	FirstAired string
	ImdbId string `xml:"IMDB_ID"`
	Zap2idId string `xml:"zap2it_id"`
	Id int `xml:"id"`
}

type GetSeriesData struct {
	XMLName xml.Name `xml:"Data"`
	Series []Series
}

func ParseGetSeries(body []byte) (result *GetSeriesData, err error) {
	result = &GetSeriesData{}
	err = xml.Unmarshal(body, result)
	return
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

func GetSeries(seriesname, language string) ([]*Series, error) {
	query, err := GetSeriesURL(seriesname, language)
	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	var resp *http.Response
	if resp, err = http.Get(query); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	data, err := ParseGetSeries(body)
	fmt.Println(err)
	fmt.Println(data)

	fmt.Println(string(body))

	return nil, nil
}
