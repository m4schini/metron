package main

import (
	tiktok "github.com/m4schini/tiktok-go"
	"github.com/m4schini/tiktok-go/scraper"
)

type Video struct {
	Url   string `json:"url"`
	Views int    `json:"views"`
}

func GetAccountInfo(account string) (*tiktok.Account, []Video, error) {
	scr, err := scraper.NewChromedpScraper()
	if err != nil {
		return nil, nil, err
	}
	defer scr.Close()

	acc, err := tiktok.GetAccountByUsername(scr, account)
	if err != nil {
		return nil, nil, err
	}

	urls, views, err := acc.GetLatestVideoURLs(scr)
	videos := make([]Video, len(urls))
	for i := 0; i < len(urls); i++ {
		videos[i] = Video{
			Url:   urls[i],
			Views: views[i],
		}
	}

	return acc, videos, nil
}

func GetVideoDetails(url string) (*tiktok.Video, error) {
	scr, err := scraper.NewChromedpScraper()
	if err != nil {
		return nil, err
	}
	defer scr.Close()

	return tiktok.GetVideoByUrl(scr, url)
}
