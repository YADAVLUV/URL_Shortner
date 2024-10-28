package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	URL          string    `json:"url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(OriginalURL))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("hash:", hash)
	return hash[:8]

}
func CreateURl(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		URL:          OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL

}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found with ID")
	}
	return url, nil

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	allowCORS(w)
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid req body", http.StatusBadRequest)
		return
	}
	sortURL := CreateURl(data.URL)
	fmt.Fprintf(w, "Short URL: %s", sortURL)

}
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.URL, http.StatusFound)
}

func main() {
	//fmt.Println("welcome to URL-shortner")
	//OriginalURL := "https://github.com/YADAVLUV?tab=repositories"
	//generateShortURL(OriginalURL)

	//handler func
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// start the http server on port 8080
	fmt.Println("start the http server on port 8080")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting the server", err)
	}
}
