package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CheckIfError エラーチェック
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// GetPodcastID URLからPodcasrIdを取得する
func GetPodcastID(url string) string {
	ids := strings.Split(url, "/id")
	id := ids[len(ids)-1]
	return id
}

// GetPage ページ取得
func GetPage(url string) {
	doc, err := goquery.NewDocument(url)

	CheckIfError(err)

	doc.Find("div.column a").Each(func(_ int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if !exists {
			s.Next()
		}
		podcastID := GetPodcastID(url)
		fmt.Println(podcastID)
	})
}

func main() {
	url := "https://podcasts.apple.com/jp/genre/podcast-%E3%82%A2%E3%83%BC%E3%83%88/id1301"
	GetPage(url)
}
