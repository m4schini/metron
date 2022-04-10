//coordinator
package main

import (
	"context"
	"coordinator-module/event"
	"coordinator-module/registry"
	"coordinator-module/server"
	"encoding/json"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"os"
	"os/signal"
	"syscall"
)

var log = kapitol.NewLogger(os.Getenv("APP_NAME"), kapitol.Debug)

func main() {
	log.Information(os.Getenv("APP_NAME"), "running...")

	// make miner registry
	reg := registry.NewMapRegistry()
	log.Debug("registry created")

	// serve grpc interface
	svr := server.NewServer(reg)
	defer svr.Close()
	go svr.Serve(os.Getenv("ADDR"))
	log.Debug("serving grpc on", os.Getenv("ADDR"))

	// connect to event bus
	events := event.NewEventBus(os.Getenv("EVENT_BUS_ADDR"), os.Getenv("EVENT_BUS_PASSWORD"))
	defer events.Close()
	log.Debug("connected to event bus on", os.Getenv("EVENT_BUS_ADDR"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	events.OnEvent(ctx, "cmd.scan.*", func(target string) {
		log.Debug("received scan command for", target)
		err := runScan(reg, events, target)
		if err != nil {
			log.Error(err)
		}
	})

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
}

func runScan(reg registry.Registrar, events event.PubSub, target string) error {
	log.Information(target, "waiting for scraper")
	scr, err := reg.Get(context.Background())
	if err != nil {
		return err
	}
	log.Information("retrieved scraper for", target)
	defer func() {
		log.Information(target, "returning scraper")
		reg.Return(scr)
	}()

	acc, vids, err := scr.GetAccount(target)
	if err != nil {
		return err
	}
	log.Information("scanned acc:", acc.Username)

	jason, err := json.Marshal(acc)
	if err != nil {
		log.Error(err)
	} else {
		events.Publish("account.scanned", string(jason))
	}

	for _, vid := range vids {
		video, err := scr.GetVideoDetails(vid.Username, vid.ID)
		if err != nil {
			log.Error(err)
			continue
		}
		video.Views = vid.Views
		log.Information("scanned vid:", video.Username, video.ID)

		jason, err := json.Marshal(video)
		if err != nil {
			log.Error(err)
			continue
		} else {
			events.Publish("video.scanned", string(jason))
		}
	}

	return nil
}
