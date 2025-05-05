package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/parchedAlbumen/veryWeebProject/apiFolder"
)

type Manga struct {
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	Chapters int    `json:"chapters"`
}

type MangaData struct {
	Data []Manga `json:"data"` //shit is a slice? ?? ? ? ?
}

// type Anime struct {
//     Title    string `json:"title"`
//     Synopsis string `json:"synopsis"`
//     URL      string `json:"url"`
//     ID       int    `json:"mal_id"`
// }

// type SearchResponse struct {
//     Data []Anime `json:"data"` //variable data with datatype off slice Anime
// }

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP\n")
}

func getNaruto(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /getNaruto request\n")
	io.WriteString(w, "I am supposed to send the naruto jikan api data here xd\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/naruto", getNaruto)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	apiFolder.Hello()
}
