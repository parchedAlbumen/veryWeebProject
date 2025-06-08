package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/parchedAlbumen/veryWeebProject/apiFolder"
)

type MangaName struct {
	MangaName string `json:"mangaName"`
}

func getScore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var name MangaName
		if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		var mangaData apiFolder.MangaData
		var info string = apiFolder.GetMangaScore(&mangaData, name.MangaName)
		io.WriteString(w, info+"\n")
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP\n")
}

func holdOn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var name MangaName //from request
		if err := json.NewDecoder(r.Body).Decode(&name); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		var data apiFolder.MangaData
		info, url := apiFolder.GetSynopsis(&data, name.MangaName)

		json.NewEncoder(w).Encode(apiFolder.ResponseData{Response: info, ImageUrl: url})
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/skibidiRizzlerSigmaMale", holdOn)
	http.HandleFunc("/getScore", getScore)

	if err := http.ListenAndServe(":3333", nil); errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
