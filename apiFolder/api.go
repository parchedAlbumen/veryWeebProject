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

func getRec(data []Manga) string {
	if len(data) > 0 {
		first := data[0]
		fmt.Println("from api:", first.Title)
		return first.Title + "\n"
	} else {
		return "No manga exist with this name bruh\n"
	}
}

func DoSomething(data *MangaData) string {
	resp, err := http.Get("https://api.jikan.moe/v4/anime?q=naruto")
	if err == nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err == nil {
		fmt.Println("Error:", err)
	}

	json.Unmarshal(body, &data)

	return getRec(data.Data)
}
