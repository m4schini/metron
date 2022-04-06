//scraper
package main

import (
	"context"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"math/rand"
	"miner-module/comm"
	"miner-module/scraper"
	"os"
	"time"
)

var log = kapitol.NewLogger(os.Getenv("APP_NAME"), kapitol.Information)

var scr scraper.Scraper

var maxRetryCooldown = 2 * time.Minute
var minRetryCooldown = 2 * time.Second

func init() {
	s, err := scraper.NewScraper()
	if err != nil {
		log.Critical(err)
	}
	scr = s
}

func main() {
	log.Information(os.Getenv("APP_NAME"), "running...")
	log.Information("ID:", scr.ID())
	log.Information("ADDR:", os.Getenv("ADDR"))

	c := comm.NewClient(log, scr.ID())
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := c.KeepAlive(
			ctx,
			os.Getenv("COORDINATOR_ADDR"),
			os.Getenv("ADDR"),
		)
		if err != nil {
			log.Error(err)
		}
	}()

	d, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int()%60))
	log.Information("alive for another", d)
	time.Sleep(d)

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
}
