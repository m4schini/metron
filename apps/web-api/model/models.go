package model

const (
	DefaultPort              = "8080"
	DbInfo                   = "root:mypassword@tcp(192.168.188.110:3306)/analytics"
	AccHistoricDataDefAmount = 800 * 4
	SQLTimestampFormat       = "2006-01-02 15:04:05"
)

type Account struct {
	Name       string `json:"name"`
	LastUpdate string `json:"lastUpdate"`
	AvatarUrl  string `json:"avatarUrl"`
	VideoCount int    `json:"videos"`
	Summary    struct {
		Views    AccountSummary `json:"views"`
		Likes    AccountSummary `json:"likes"`
		Comments AccountSummary `json:"comments"`
		Shares   AccountSummary `json:"shares"`
	} `json:"summary"`
}

type AccountSummary struct {
	Value    int     `json:"value"`
	FiveMin  float32 `json:"fiveMin"`
	OneDay   float32 `json:"oneDay"`
	OneMonth float32 `json:"oneMonth"`
}

type Video struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Views       int    `json:"views"`
	Comments    int    `json:"comments"`
	Shares      int    `json:"shares"`
	Likes       int    `json:"likes"`
	PostedBy    string `json:"postedBy"`
	Mentioned   string `json:"mentioned"`
}

/*
NIVO CHART STRUCTS
*/

//LineData single line data struct
type LineData struct {
	Id   string            `json:"id"`
	Data []LineDataElement `json:"data"`
}

//LineDataElement represents a single datapoint of a line graph.
//x and y coordinate.
type LineDataElement struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

//PieData array of pie data
type PieData []PieDataElement

//PieDataElement one part of a pie
type PieDataElement struct {
	Id    string `json:"id"`
	Value int    `json:"value"`
}

type NetworkNode struct {
	Id     string `json:"id"`
	Size   int    `json:"size"`
	Height int    `json:"height"`
}
type NetworkLink struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	Distance int    `json:"distance"`
}
type NetworkData struct {
	Nodes []NetworkNode `json:"nodes"`
	Links []NetworkLink `json:"links"`
}

type TimeRangeData []TimeRangeElement
type TimeRangeElement struct {
	Day   string `json:"day"`
	Value int    `json:"value"`
}
