package scraper

import (
	"github.com/google/uuid"
	tiktok_go "github.com/m4schini/tiktok-go"
	"sync"
)

type Scraper interface {
	ID() string
	TikTok() tiktok_go.TikTok
}

type scraper struct {
	id string
	tt tiktok_go.TikTok

	mu sync.Mutex
}

func NewScraper() (*scraper, error) {
	scr := new(scraper)
	scr.id = uuid.New().String()
	scr.tt = tiktok_go.NewTikTok()

	return scr, nil
}

func (s *scraper) TikTok() tiktok_go.TikTok {
	return s.tt
}

func (s *scraper) ID() string {
	return s.id
}
