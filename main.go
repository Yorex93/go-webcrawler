package main

import (
	"fmt"
	"go-webcrawler/crawler"
)

func main() {
	fmt.Println("Hello GO!")
	result := crawler.BuildSiteMap("https://monzo.com")
	fmt.Print(result)
}
