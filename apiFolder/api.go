package apiFolder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Manga struct {
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	Chapters int    `json:"chapters"`
}

type MangaData struct {
	Data []Manga `json:"data"`
}

func getRec(manga []Manga) string {
	// getManga(data, name)
	if len(manga) > 0 {
		first := manga[0]
		fmt.Println("giving user informations")
		return first.Title + " " + first.Synopsis + "\n"
	} else {
		return "No manga exist with this name bruh\n"
	}
}

func GiveSynopsis(data *MangaData, name string) string {
	theManga := getManga(data, name)
	if len(theManga) > 0 {
		first := theManga[0]
		return first.Title + "\n" + first.Synopsis + "\n"
	} else {
		return "lol xd\n"
	}
}

func getManga(data *MangaData, name string) []Manga {
	theName := "https://api.jikan.moe/v4/anime?q=" + name
	resp, err := http.Get(theName)
	if err != nil {
		fmt.Println("(1)Error:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("(2)Error:", err)
	}

	json.Unmarshal(body, &data)

	return data.Data
}
