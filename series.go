package tvdb

type Series struct {
	SeriesId int `xml:"seriesid"`
	Language string `xml:"language"`
	SeriesName string
}
