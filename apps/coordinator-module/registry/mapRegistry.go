package registry

import (
	"context"
	"coordinator-module/miner"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type registryItem struct {
	scraper miner.Miner
	locked  bool
}

type mapRegistry struct {
	scrapers map[string]*registryItem
	mu       sync.Mutex
}

func NewMapRegistry() *mapRegistry {
	r := new(mapRegistry)
	r.scrapers = make(map[string]*registryItem)

	return r
}

func (r *mapRegistry) Register(scraper miner.Miner) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.scrapers[scraper.ID()]
	if exists {
		return errors.New(fmt.Sprintf("scraper \"%s\" already registered", scraper.ID()))
	}

	r.scrapers[scraper.ID()] = &registryItem{
		scraper: scraper,
		locked:  false,
	}
	return nil
}

func (r *mapRegistry) Unregister(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.scrapers[id] = nil
	delete(r.scrapers, id)
	return nil
}

func (r *mapRegistry) Get(ctx context.Context) (miner.Miner, error) {
	scrCh := make(chan miner.Miner)
	stop := false

	go func() {
		for !stop {
			r.mu.Lock()
			for _, item := range r.scrapers {
				if !item.locked {
					item.locked = true

					scrCh <- item.scraper
					r.mu.Unlock()
					return
				}
			}
			r.mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-ctx.Done():
		// return if context is done\
		stop = true
		return nil, ctx.Err()

	case scr := <-scrCh:
		// wait until scraper becomes available
		return scr, nil
	}
}

func (r *mapRegistry) GetNow() (miner.Miner, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.scrapers {
		if !item.locked {
			item.locked = true

			return item.scraper, nil
		}
	}

	return nil, errors.New("no scraper available")
}

func (r *mapRegistry) Return(scraper miner.Miner) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.scrapers[scraper.ID()]
	if !exists {
		return errors.New("scraper isn't registered")
	}

	item.locked = false
	return nil
}

func (r *mapRegistry) String() string {
	lines := make([]string, 1)

	lines[0] = fmt.Sprintf("\nRegistry(%d):", len(r.scrapers))

	for id, item := range r.scrapers {
		lines = append(lines, fmt.Sprintf("+ [%s] LOCKED=%v", id, item.locked))
	}

	return strings.Join(lines, "\n")
}

func (r *mapRegistry) Len() int {
	return len(r.scrapers)
}
