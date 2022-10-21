package demo

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func Goquery() {
	doc, _ := goquery.NewDocument("http://live.titan007.com/")
	log.Println(doc.Find("#team1_1527132").Text())
}
