package registry

import (
	"context"
	"coordinator-module/miner"
)

type Registrar interface {
	//Register stores a new scraper in the registry
	Register(scraper miner.Miner) error

	//Unregister removes a scraper from the registry
	Unregister(id string) error

	//Get retrieves a scraper from the registry and locks that scraper. Blocks and waits for scrapers to become available
	Get(context.Context) (miner.Miner, error)

	//GetNow retrieves a scraper from the registry and locks that scraper. Returns error if no scrapers are available
	GetNow() (miner.Miner, error)

	//Return unlocks a scraper in the registry, so it can be used again
	Return(scraper miner.Miner) error
}
