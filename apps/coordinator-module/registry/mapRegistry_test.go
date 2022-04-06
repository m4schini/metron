package registry

import (
	"context"
	"coordinator-module/scraper"
	"testing"
	"time"
)

func TestMapRegistry_Get(t *testing.T) {
	// arrange
	reg := NewMapRegistry()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)

	//act
	_, err := reg.Get(ctx)
	cancel()

	// assert
	if err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestMapRegistry_GetNow_EmptyRegistry(t *testing.T) {
	// arrange
	reg := NewMapRegistry()

	// act
	_, err := reg.GetNow()

	// assert
	if err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestMapRegistry_Register(t *testing.T) {
	// arrange
	reg := NewMapRegistry()

	// act
	err := reg.Register(scraper.NewTestScraper())

	// assert
	if err != nil {
		t.Fatal(err)
	}

	if reg.Len() != 1 {
		t.Fail()
	}
}

func TestMapRegistry_Unregister(t *testing.T) {
	// arrange
	scr := scraper.NewTestScraper()
	reg := NewMapRegistry()
	err := reg.Register(scr)
	if err != nil {
		t.Fatal(err)
	}

	// act
	err = reg.Unregister(scr)

	// assert
	if err != nil {
		t.Fatal(err)
	}

	if reg.Len() != 0 {
		t.Log("Expected: 0 | Actual:", reg.Len())
		t.Fail()
	}
}

func TestMapRegistry_GetNow(t *testing.T) {
	// arrange
	scr := scraper.NewTestScraper()
	reg := NewMapRegistry()
	err := reg.Register(scr)
	if err != nil {
		t.Fatal(err)
	}

	// act
	scr2, err := reg.GetNow()

	// assert
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if scr.ID() != scr2.ID() {
		t.Log("Expected:", scr.ID(), "| Actual:", scr2.ID())
		t.Fail()
	}
}
