package main

import (
	"fmt"
	"time"
)

type URL struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

func main() {
	fmt.Println("welcome to URL-shortner")
}