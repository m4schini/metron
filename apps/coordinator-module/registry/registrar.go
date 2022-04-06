package registry

import (
	"context"
	"coordinator-module/scraper"
)

type Registrar interface {
	//Register stores a new scraper in the registry
	Register(scraper scraper.Scraper) error

	//Unregister removes a scraper from the registry
	Unregister(scraper scraper.Scraper) error

	//Get retrieves a scraper from the registry and locks that scraper. Blocks and waits for scrapers to become available
	Get(context.Context) (scraper.Scraper, error)

	//GetNow retrieves a scraper from the registry and locks that scraper. Returns error if no scrapers are available
	GetNow() (scraper.Scraper, error)

	//Return unlocks a scraper in the registry, so it can be used again
	Return(scraper scraper.Scraper) error
}
