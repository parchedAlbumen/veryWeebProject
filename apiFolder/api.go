package apiFolder

//main file basically

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

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
	if theManga := getManga(data, name); len(theManga) > 0 {
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
		dataInfo := getScoreStatistic(httpreq, &scoreData) //no need to check if exist or not,, because we can only get something if there's an id
		infoFormat := ""
		fmt.Println(infoFormat)
		infoFormat += ("Completed: " + strconv.Itoa(dataInfo.Completed) + "\n")
		infoFormat += ("Total Users: " + strconv.Itoa(dataInfo.Total) + "\n")
		averageRating := calculateAverageRating(dataInfo)
		infoFormat += ("Current Rating: " + strconv.FormatFloat(averageRating, 'f', 2, 64))
		infoFormat += ("Percentage of people that dropped it: ")
		return infoFormat
	} else {
		return "bad at getMangaScore"
	}
}

func getScoreStatistic(theReq string, data *MangaScoreData) ScoreData {
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
	if theManga := getManga(data, name); len(theManga) > 0 {
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
