package scraper

import "coordinator-module/model"

type Scraper interface {
	ID() string
	GetAccount(username string) (*model.Account, error)
	GetVideoDetails(username, videoId string) (*model.Account, error)
}

type scraper struct {
	id string
}

func (s scraper) ID() string {
	return s.id
}

func NewScraper(id string) *scraper {
	scr := new(scraper)
	scr.id = id

	return scr
}
