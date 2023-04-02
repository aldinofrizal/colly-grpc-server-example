package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aldinofrizal/soompi-colly-scrap/config"
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)

func Scrap() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.soompi.com"),
	)

	newsCol := config.DB.Database("soompi").Collection("tvnews")

	_, err := newsCol.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

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
		}
	})

	c.Visit("https://www.soompi.com/category/tvfilm")
}
