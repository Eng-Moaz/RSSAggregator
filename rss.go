package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil{
		return nil, fmt.Errorf("Error in requesting: %v", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		return nil, fmt.Errorf("Error in requesting: %v", err)
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil{
		return nil, fmt.Errorf("Error in reading: %v", err)
	}
	
	var rssfeed RSSFeed
	err = xml.Unmarshal(data, &rssfeed)
	if err != nil{
		return nil, fmt.Errorf("Error in Unmarshaling: %v", err)
	}
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)		
	
	return &rssfeed, nil
}
