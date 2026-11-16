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

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	getReq, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)

	if err != nil {
		return nil, err
	}

	getReq.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(getReq)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var xmlResp RSSFeed
	err = xml.Unmarshal(data, &xmlResp)

	if err != nil {
		return nil, err
	}

	xmlResp.Channel.Title = html.UnescapeString(xmlResp.Channel.Title)
	xmlResp.Channel.Description = html.UnescapeString(xmlResp.Channel.Description)

	for i := range xmlResp.Channel.Item {
		xmlResp.Channel.Item[i].Title = html.UnescapeString(xmlResp.Channel.Item[i].Title)
		xmlResp.Channel.Item[i].Description = html.UnescapeString(xmlResp.Channel.Item[i].Description)
	}

	return &xmlResp, nil
}
