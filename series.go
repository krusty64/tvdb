package tvdb

import (
	"encoding/xml"
	"net/url"
	"fmt"
)

type Series struct {
	XMLName xml.Name `xml:"Series"`
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
