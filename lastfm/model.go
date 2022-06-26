package lastfm

type Info struct {
	Track Track `json:"track"`
}

type Track struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Artist   Artist `json:"artist"`
	Album    Album  `json:"album"`
}

type Artist struct {
	Name string `json:"name"`
	MbID string `json:"mbid"`
	Url  string `json:"url"`
}

type Album struct {
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Url    string  `json:"url"`
	Image  []Image `json:"image"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}
