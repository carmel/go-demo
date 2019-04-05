package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	doc, _ := goquery.NewDocument("http://live.titan007.com/")
	log.Println(doc.Find("#team1_1527132").Text())
}
