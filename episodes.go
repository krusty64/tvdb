package tvdb

import (
	"encoding/xml"
	"strconv"
)

type Episode struct {
	Id                    uint64 `xml:"id"`
	SeasonId              uint64 `xml:"seasonid"`
	CombinedEpisodeNumber string `xml:"Combined_episodenumber"`
	CombinedSeason        string `xml:"Combined_season"`
	EpisodeName           string
	EpisodeNumber         int
	SeasonNumber          int
	FirstAired            string
	Director              string
	Writer                string
	GuestStars            string
	Overview              string
	ProductionCode        string
	LastUpdated           uint64 `xml:"lastupdated"`
	Flagged               int    `xml:"flagged"`
	DvdDiscId             string `xml:"DVD_discid"`
	DvdSeason             string `xml:"DVD_season"`
	DvdEpisodeNumber      string `xml:"DVD_episodenumber"`
	DvdChapter            string `xml:"DVD_chapter"`
	AbsoluteNumber        string `xml:"absolute_number"`
	Filename              string `xml:"filename"`
	SeriesId              uint64 `xml:"seriesid"`
	MirrorUpdate          string `xml:"mirrorupdate"`
	ImdbId                string `xml:"IMDB_ID"`
	EpImgFlag             string
	Rating                string
	Language              string
}

type SingleEpisodeData struct {
	XMLName xml.Name `xml:"Data"`
	Episode *Episode
}

type FullSeriesData struct {
	XMLName xml.Name `xml:"Data"`
	Series  Series
	Episode []Episode
}

func (t *TVDB) GetEpisodeBySeasonEp(seriesId, season, episode int, language string) ([]byte, error) {
	return t.QueryURL(t.ApiKey+"/series/"+strconv.Itoa(seriesId)+"/default/"+strconv.Itoa(season)+
		"/"+strconv.Itoa(episode)+"/"+language+".xml", nil)
}

func ParseSingleEpisode(src []byte) (*SingleEpisodeData, error) {
	var r SingleEpisodeData
	if err := xml.Unmarshal(src, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func ParseEpisode(src []byte) (*Episode, error) {
	var r Episode
	if err := xml.Unmarshal(src, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (t *TVDB) GetFullSeriesData(seriesId int, language string) ([]byte, error) {
	return t.QueryURL(t.ApiKey+"/series/"+strconv.Itoa(seriesId)+"/all/"+language+".xml", nil)
}

func ParseFullSeriesData(src []byte) (*FullSeriesData, error) {
	var f FullSeriesData
	if err := xml.Unmarshal(src, &f); err != nil {
		return nil, err
	}
	return &f, nil
}
