package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aldinofrizal/soompi-colly-scrap/config"
	pb "github.com/aldinofrizal/soompi-colly-scrap/news"
	"go.mongodb.org/mongo-driver/bson"
)

type News struct {
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func GetNews() ([]*pb.Item, error) {
	fmt.Println("get news is called ----")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var newsList []*pb.Item
	defer cancel()
	newsCol := config.DB.Database("soompi").Collection("tvnews")

	cur, err := newsCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("failed to fetch colleciton")
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var news pb.Item
		if err = cur.Decode(&news); err != nil {
			return nil, errors.New("failed to fetch collection")
		}
		newsList = append(newsList, &news)
	}

	return newsList, nil
}
