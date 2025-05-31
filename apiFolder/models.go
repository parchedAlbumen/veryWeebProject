package apiFolder

//file dedicated for models

type Manga struct {
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	Chapters int    `json:"chapters"`
	Id       int    `json:"mal_id"`
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
