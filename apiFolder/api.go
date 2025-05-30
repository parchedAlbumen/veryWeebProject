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

func GetRec(manga []Manga) string {
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
	if id := getMangaId(data, name); id > 0 {
		httpreq := "https://api.jikan.moe/v4/manga/" + strconv.Itoa(id) + "/statistics"
		var scoreData MangaScoreData
		dataInfo := getScoreStatistic(httpreq, &scoreData)
		if len(dataInfo.Scores) > 0 {
			infoFormat := ""
			fmt.Println(infoFormat)
			infoFormat += ("Completed: " + strconv.Itoa(dataInfo.Completed) + "\n")
			infoFormat += ("Total: " + strconv.Itoa(dataInfo.Total) + "\n")
			mangaScores := dataInfo.Scores[0]
			infoFormat += ("Votes: " + strconv.Itoa(mangaScores.Votes) + "\n")
			infoFormat += ("Percentages: " + strconv.Itoa(mangaScores.Percentage) + "%\n")
			return infoFormat
		} else {
			return "no info what the!"
		}
	} else {
		return "bad at getMangaScore"
	}
}

func getScoreStatistic(theReq string, data *MangaScoreData) ScoreData {
	fmt.Println(theReq + "lol")
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
