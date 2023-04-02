package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aldinofrizal/soompi-colly-scrap/config"
	"github.com/gocolly/colly"
)

func Scrap() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.soompi.com"),
	)

	client := config.InitDatabase()
	newsCol := client.Database("soompi").Collection("tvnews")

	c.OnHTML(".thumbnail-wrapper a", func(h *colly.HTMLElement) {
		title := h.Attr("title")
		link := "https://www.soompi.com" + h.Attr("href")
		imgUrl := h.ChildAttr("img", "src")

		news := News{title, link, imgUrl, time.Now()}
		_, err := newsCol.InsertOne(context.Background(), news)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("news inserted")
			fmt.Println(news)
		}
	})

	c.Visit("https://www.soompi.com/category/tvfilm")
}

func PullData() {

}
