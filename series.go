package tvdb

import (
	"encoding/xml"
	"net/url"
	"fmt"
)

type Series struct {
	XML []byte `xml:",innerxml"`
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
	XML []byte `xml:",innerxml"`
	XMLName xml.Name `xml:"Data"`
	Series []Series
}

func (t *TVDB) GetSeries(seriesname, language string) ([]Series, error) {
	args := &url.Values{}
	if seriesname == "" {
		return nil, fmt.Errorf("GetSeriesURL: Series name must not be empty")
	}

	args.Add("seriesname", seriesname)

	if language != "" {
		args.Add("language", language)
	}

	var data GetSeriesData
	err := t.QueryAndUnmarshal("GetSeries.php", args, &data)

	if err != nil {
		return nil, err
	}

	return data.Series, nil
}
