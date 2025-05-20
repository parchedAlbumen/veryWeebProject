package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/parchedAlbumen/veryWeebProject/apiFolder"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP\n")
}

func holdOn(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "What manga would you like to look at? (make sure if you want to type with spaces,, instead of spaces use dashes, '-'")
	var manga string = ""
	sentence := "The manga mentioned: " + manga
	fmt.Println(sentence)
	io.WriteString(w, sentence)
}

func getNaruto(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /getNaruto request\n")
	var data apiFolder.MangaData
	var info string = apiFolder.DoSomething(&data)
	fmt.Println("from server:", info)
	io.WriteString(w, info)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/naruto", getNaruto)
	http.HandleFunc("/wassup", holdOn)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

//next plan is to expand on the stuff I want to do with the manga jikan api calls and then finally start with the gui
