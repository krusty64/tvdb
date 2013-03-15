package api

type Episode struct {
	Id int `xml:"id"`
	CombinedEpisodeNumber int `xml:"Combined_episodenumber"`
	CombinedSeason int `xml:"Combined_season"`
	EpisodeName string
	EpisodeNumber int
	SeasonNumber int
}

func GetAllEpisodes()
