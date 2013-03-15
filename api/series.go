package api

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"strconv"
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

func GetSeriesURL(seriesname, language string) (result string, err error) {
	var args url.Values
	if seriesname == "" {
		err = fmt.Errorf("GetSeriesURL: Series name must not be empty")
		return
	}

	args.Add("seriesname", seriesname)

	if language != "" {
		args.Add("language", language)
	}

	return GetURL("GetSeries.php", &args)
}



func GetSeries(seriesname, language string) ([]*Series, error) {
	query, err := GetSeriesURL(seriesname, language)
	if err != nil {
		return nil, err
	}

	body, err := HttpGet(query)

	if err != nil {
		return nil, err
	}

	data, err := ParseGetSeries(body)

	if err != nil {
		return nil, err
	}

	result := make([]*Series, 0, len(data.Series))
	for _, v := range data.Series {
		result = append(result, &v)
	}

	return result, nil
}

func GetEpisodesURL(apikey string, seriesid int, language string) (string, error) {
	query, err := url.Parse(TVDB_API)
	if err != nil {
		return "", err
	}

	if query, err = query.Parse(apikey + "/series/" + strconv.Itoa(seriesid) + "/all/" + language + ".xml"); err != nil {
		return "", err
	}

	return query.String(), nil
}

func GetEpisodes(apikey string, seriesid int, language string) (error, error) {
	query, err := GetEpisodesURL(apikey, seriesid, language)
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	if resp, err = http.Get(query); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	return nil, nil
}

func GetEpisodeByDefault(apikey string, seriesid int, language string)
