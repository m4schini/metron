package scraper

import "github.com/google/uuid"

type testScraper struct {
	id string
}

func NewTestScraper() *testScraper {
	scr := new(testScraper)
	scr.id = uuid.New().String()

	return scr
}

func (t *testScraper) ID() string {
	return t.id
}

