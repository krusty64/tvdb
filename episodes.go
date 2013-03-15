package tvdb

import (
	"strconv"
	"encoding/xml"
)

type Episode struct {
	Id uint64 `xml:"id"`
	SeasonId uint64 `xml:"seasonid"`
	CombinedEpisodeNumber string `xml:"Combined_episodenumber"`
	CombinedSeason string `xml:"Combined_season"`
	EpisodeName string
	EpisodeNumber int
	SeasonNumber int
	FirstAired string
	Director string
	Writer string
	GuestStars string
	Overview string
	ProductionCode string
	LastUpdated uint64 `xml:"lastupdated"`
	Flagged int `xml:"flagged"`
	DvdDiscId string `xml:"DVD_discid"`
	DvdSeason string `xml:"DVD_season"`
	DvdEpisodeNumber string `xml:"DVD_episodenumber"`
	DvdChapter string `xml:"DVD_chapter"`
	AbsoluteNumber string `xml:"absolute_number"`
	Filename string `xml:"filename"`
	SeriesId uint64 `xml:"seriesid"`
	MirrorUpdate string `xml:"mirrorupdate"`
	ImdbId string `xml:"IMDB_ID"`
	EpImgFlag string
	Rating string
	Language string
}

type singleEpisodeData struct {
	XMLName xml.Name `xml:"Data"`
	Episode *Episode
}

type FullSeriesData struct {
	XMLName xml.Name `xml:"Data"`
	Series Series
	Episode []Episode
}

func (t *TVDB) GetEpisodeBySeasonEp(seriesId, season, episode int, language string) (result *Episode, err error) {
	var data singleEpisodeData
	err = t.QueryAndUnmarshal(t.ApiKey + "/series/" + strconv.Itoa(seriesId) + "/default/" + strconv.Itoa(season) +
				"/" + strconv.Itoa(episode) + "/" + language + ".xml", nil, &data)
	if err != nil {
		return
	}
	result = data.Episode
	return
}

func (t *TVDB) GetFullSeriesData(seriesId int, language string) (result FullSeriesData, err error) {
	err = t.QueryAndUnmarshal(t.ApiKey + "/series/" + strconv.Itoa(seriesId) + "/all/" + language + ".xml", nil, &result)
	return
}
