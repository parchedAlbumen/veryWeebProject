package apiFolder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetRecommendation(data *MangaData, name string) (string, string) {
	id, trashImage := getMangaId(data, name)
	if id > 0 {
		theReq := "https://api.jikan.moe/v4/manga/" + strconv.Itoa(id) + "/recommendations"
		var recommendMangaData MangaRecommendData
		recManga := getRecData(&recommendMangaData, theReq)
		if len(recManga) > 0 {
			first := recManga[0]
			// fmt.Println("Id:", strconv.Itoa(first.Entry.Id))
			fmt.Println("Sending these infos:")
			fmt.Println("image:", first.Entry.Images.Jpg.Image_URL)
			fmt.Println("title:", first.Entry.Title)

			return first.Entry.Title, first.Entry.Images.Jpg.Image_URL
		}
	}
	return "hi", trashImage
}

func getRecData(rec *MangaRecommendData, url string) []RecommendData {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close() //defer makes sure resp (http connection) closes no matter what, in order to avoid leaks

	body, err := io.ReadAll(resp.Body) // body is a byte at this point
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, rec)
	return rec.Data
}

func GetSynopsis(data *MangaData, name string) (string, string) {
	if theManga := getManga(data, name); len(theManga) > 0 {
		first := theManga[0]
		return first.Title + "\n" + first.Synopsis + "\n", first.Images.Jpg.Image_URL
	} else {
		return "lol xd\n", ""
	}
}

func GetMangaScore(data *MangaData, name string) (string, string) {
	if id, image_url := getMangaId(data, name); id > 0 {
		httpreq := "https://api.jikan.moe/v4/manga/" + strconv.Itoa(id) + "/statistics"
		var scoreData MangaScoreData
		dataInfo := getScoreStatistic(httpreq, &scoreData) //no need to check if exist or not,, because we can only get something if there's an id
		infoFormat := ""
		generateScoreFormat(dataInfo)
		fmt.Println(infoFormat)
		infoFormat += ("Completed: " + strconv.Itoa(dataInfo.Completed) + "\n")
		infoFormat += ("Total Users: " + strconv.Itoa(dataInfo.Total) + "\n")
		averageRating := calculateAverageRating(dataInfo)
		infoFormat += ("Current Rating: " + strconv.FormatFloat(averageRating, 'f', 2, 64))
		infoFormat += ("\nPercentage of people that dropped it: ")
		return infoFormat, image_url
	} else {
		return "No score", image_url
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
	json.Unmarshal(body, data)
	return data.Data
}

func getMangaId(data *MangaData, name string) (int, string) {
	if theManga := getManga(data, name); len(theManga) > 0 {
		first := theManga[0]
		return first.Id, first.Images.Jpg.Image_URL
	} else {
		return -1, "no image"
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
	json.Unmarshal(body, data)
	return data.Data
}
