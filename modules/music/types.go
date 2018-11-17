package music

type Music struct {
	Title         string  `json:"title"`
	Extractor     string  `json:"extractor"`
	Uploader      string  `json:"uploader"`
	Thumbnail     string  `json:"thumbnail"`
	Type          string  `json:"_type"`
	Entries       []Music `json:"entries"`
	URL           string  `json:"webpage_url"`
	Duration      int     `json:"duration"`
	RequesterID   string
	RequesterName string
}
