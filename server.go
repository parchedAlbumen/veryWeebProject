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
		var name MangaName
		err := json.NewDecoder(r.Body).Decode(&name)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Println(name.MangaName, "porno")
		var data apiFolder.MangaData
		var info string = apiFolder.GiveSynopsis(&data, name.MangaName)
		io.WriteString(w, info+"\n"+name.MangaName)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/skibidiRizzlerSigmaMale", holdOn)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

//next plan is to expand on the stuff I want to do with the manga jikan api calls and then finally start with the gui
