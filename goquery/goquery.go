package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func BaiduHotSearch() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Text())
	//doc.Find(".s-hotsearch-content .hotsearch-item").Each(func(i int, s *goquery.Selection) {
	doc.Find(".s-news-rank-content .news-meta-item").Each(func(i int, s *goquery.Selection) {
		content := s.Find("title-content-title").Text()
		fmt.Printf("%d: %s\n", i, content)
	})
}

func main() {
	BaiduHotSearch()
}
