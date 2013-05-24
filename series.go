package tvdb

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
)

type Series struct {
	XMLName    xml.Name `xml:"Series"`
	SeriesId   int      `xml:"seriesid"`
	Language   string   `xml:"language"`
	SeriesName string
	Overview   string
	Banner     string `xml:"banner"`
	FirstAired string
	ImdbId     string `xml:"IMDB_ID"`
	Zap2idId   string `xml:"zap2it_id"`
	Id         int    `xml:"id"`
}

type GetSeriesData struct {
	XMLName xml.Name `xml:"Data"`
	Series  []Series
}

type DetailSeries struct {
	XMLNAME        xml.Name `xml:"Series"`
	Id             int      `xml:"id"`
	Actors         string
	Airs_DayOfWeek string
	Airs_Time      string
	ContentRating  string
	FirstAired     string
	Genre          string
	ImdbId         string `xml:"IMDB_ID"`
	Language       string
	Network        string
	NetworkId      string `xml:"NetworkID"`
	Overview       string
	Rating         float32
	RatingCount    int
	Runtime        int
	SeriesId       string `xml:"SeriesID"`
	SeriesName     string
	Status         string
	Banner         string `xml:"banner"`
	Fanart         string `xml:"fanart"`
	Lastupdated    int    `xml:"lastupdated"`
	Poster         string `xml:"poster"`
	Zap2itId       string `xml:"zap2it_id"`
}

type GetDetailSeriesData struct {
	XMLName xml.Name `xml:"Data"`
	Series  []DetailSeries
}

func (t *TVDB) GetSeries(seriesname, language string) ([]byte, error) {
	args := &url.Values{}
	if seriesname == "" {
		return nil, fmt.Errorf("GetSeriesURL: Series name must not be empty")
	}

	args.Add("seriesname", seriesname)

	if language != "" {
		args.Add("language", language)
	}

	return t.QueryURL("GetSeries.php", args)
}

func ParseGetSeries(src []byte) (*GetSeriesData, error) {
	var result GetSeriesData
	if err := xml.Unmarshal(src, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TVDB) GetDetailSeriesById(seriesId int, language string) ([]byte, error) {
	return t.QueryURL(t.ApiKey+"/series/"+strconv.Itoa(seriesId)+"/"+
		language+".xml", nil)
}

func ParseDetailSeriesData(src []byte) (*GetDetailSeriesData, error) {
	var result GetDetailSeriesData
	if err := xml.Unmarshal(src, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
