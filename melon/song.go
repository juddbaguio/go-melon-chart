package melon

type ListType string

var (
	TOP_100 ListType = "TOP 100"
	DAILY   ListType = "DAILY"
	MONTHLY ListType = "MONTHLY"
	AGE     ListType = "AGE"
)

type TopSongsSummary struct {
	Summary ListType     `json:"summary"`
	Date    string       `json:"date"`
	Time    string       `json:"time"`
	Records SongDataList `json:"records"`
}

type SongData struct {
	Id       int    `json:"id"`
	Rank     int    `json:"rank"`
	AlbumImg string `json:"albumImg"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
}

type SongDataList []SongData
