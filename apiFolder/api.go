package apiFolder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Manga struct {
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	Chapters int    `json:"chapters"`
	Id       int    `json:"mal_id"`
}

type MangaData struct {
	Data []Manga `json:"data"`
}

type MangaScore struct {
	Score      int `json:"score"`
	Votes      int `json:"votes"`
	Percentage int `json:"percentage"`
}

type ScoreData struct {
	Completed int          `json:"completed"`
	Total     int          `json:total"`
	Scores    []MangaScore `json:scores"`
}

type MangaScoreData struct {
	Data []ScoreData `json:"data"`
}

func GetRec(manga []Manga) string {
	// getManga(data, name)
	if len(manga) > 0 {
		first := manga[0]
		fmt.Println("giving user informations")
		return first.Title + " " + first.Synopsis + "\n"
	} else {
		return "No manga exist with this name bruh\n"
	}
}

func GetSynopsis(data *MangaData, name string) string {
	theManga := getManga(data, name)
	if len(theManga) > 0 {
		first := theManga[0]
		return first.Title + "\n" + first.Synopsis + "\n"
	} else {
		return "lol xd\n"
	}
}

func GetMangaScore(data *MangaData, name string) string {
	//get manga id first
	//bro i love this initalization and if statement in one line
	if id := getMangaId(data, name); id > 0 {
		//do this
		httpStuff := "https://api.jikan.moe/v4/manga/" + strconv.Itoa(id) + "/statistics"
		var scoreData MangaScoreData
		yomama := getScoreStatistic(httpStuff, &scoreData) //yo mama is []ScoreData
		if len(yomama) > 0 {
			//do magic wor here
		} else {
			return "bad at yomama"
		}
		return "hi"
	} else {
		return "bad at getMangaScore"
	}
}

// why do I thik []scoredata is the right return type, figure this out
func getScoreStatistic(theReq string, data *MangaScoreData) []ScoreData {
	resp, err := http.Get(theReq)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &data)
	return data.Data
}

func getMangaId(data *MangaData, name string) int {
	theManga := getManga(data, name)
	if len(theManga) > 0 {
		first := theManga[0]
		return first.Id
	} else {
		return -1
	}
}

func getManga(data *MangaData, name string) []Manga {
	theName := "https://api.jikan.moe/v4/manga?q=" + name
	resp, err := http.Get(theName)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &data)

	return data.Data
}
