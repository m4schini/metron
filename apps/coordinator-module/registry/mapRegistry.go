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
	capacity int
}

func NewMapRegistry() *mapRegistry {
	r := new(mapRegistry)
	r.scrapers = make(map[string]*registryItem)
	r.capacity = 0

	return r
}

func (r *mapRegistry) Available() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.capacity
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
	r.capacity = r.capacity + 1
	return nil
}

func (r *mapRegistry) Unregister(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.scrapers[id] = nil
	delete(r.scrapers, id)
	r.capacity = r.capacity - 1
	return nil
}

func (r *mapRegistry) Get(ctx context.Context) (miner.Miner, error) {
	select {
	case <-ctx.Done():
		// return if context is done\
		return nil, ctx.Err()

	default:
		// wait until scraper becomes available
		for {
			scr, err := r.GetNow()
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			return scr, nil
		}
	}
}

func (r *mapRegistry) GetNow() (miner.Miner, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range r.scrapers {
		if !item.locked {
			item.locked = true
			r.capacity = r.capacity - 1

			return item.scraper, nil
		}
	}

	return nil, errors.New("no scraper available")
}

func (r *mapRegistry) Return(scraper miner.Miner) error {
	if scraper == nil {
		return nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.scrapers[scraper.ID()]
	if !exists {
		return errors.New("scraper isn't registered")
	}

	r.capacity = r.capacity + 1
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
