package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"fmt"
	"log"
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

func main() {
	fmt.Println("welcome to URL-shortner")
	OriginalURL := "https://github.com/YADAVLUV?tab=repositories"
	generateShortURL(OriginalURL)

}
