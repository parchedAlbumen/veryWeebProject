package apiFolder

//file dedicated for models

type Manga struct {
	Title    string   `json:"title"`
	Synopsis string   `json:"synopsis"`
	Chapters int      `json:"chapters"`
	Id       int      `json:"mal_id"`
	Images   JpgPhoto `json:"images"`
}

type MangaData struct {
	Data []Manga `json:"data"` //this is a slice of object: Manga
}

type MangaScore struct {
	Score      int `json:"score"`
	Votes      int `json:"votes"`
	Percentage int `json:"percentage"`
}

type ScoreData struct {
	Completed int          `json:"completed"`
	Total     int          `json:total"`
	Scores    []MangaScore `json:scores"` //this is a slice of object: MangaScore
}

type MangaScoreData struct {
	Data ScoreData `json:"data"`
}

type EntryRecommended struct {
	Id     int      `json:"mal_id"`
	Url    string   `json:"url"`
	Images JpgPhoto `json:"images"`
	Title  string   `json:"title"`
}

type RecommendData struct {
	Entry EntryRecommended `json:"entry"`
}

type MangaRecommendData struct {
	Data []RecommendData `json:"data"`
}

// images
type Jpgs struct {
	Image_URL   string `json:"image_url"`
	Small_Image string `json:"small_image_url"`
	Large_Image string `json:"large_image_url"`
}

// "jpg: { stuff }"
type JpgPhoto struct {
	Jpg Jpgs `json:"jpg"`
}

// lines for ways to respond to use
// used for sending back information in terms of this json format, so I can send multiple informations
type ResponseData struct {
	Response string `json:"response"`
	ImageUrl string `json:"imageurl"`
}
