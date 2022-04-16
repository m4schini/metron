package main

import (
	"bufio"
	"commander-module/api"
	"commander-module/config"
	"commander-module/event"
	"commander-module/persistence"
	"context"
	"encoding/json"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"os"
	"strings"
	"time"
)

var log = kapitol.NewLogger(os.Getenv("APP_NAME"), kapitol.Debug)

func main() {
	log.Information(os.Getenv("APP_NAME"), "running...")

	events := event.NewEventBus(os.Getenv("EVENT_BUS_ADDR"), os.Getenv("EVENT_BUS_PASSWORD"))
	defer events.Close()

	targets := config.GetTargets()
	fmt.Println("Targets:", len(targets))
	for i, target := range targets {
		fmt.Printf("%d. Target: %s | Interval: %s\n", i, target.Account, target.Interval)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func(target *config.Target) {
			select {
			case <-ctx.Done():
				return
			default:
				for {
					fmt.Println("CMD > Scan Account:", target.Account)
					events.Publish("account.scan.requested", target.Account)
					fmt.Printf("CMD > Next Scan for '%s' scheduled at %s\n", target.Account, time.Now().Add(target.Interval))
					time.Sleep(target.Interval)
				}
			}
		}(target)
	}

	db, err := persistence.NewPersister()
	if err != nil {
		log.Critical(err)
	}
	persist := err == nil
	log.Information("PERSISTING DATA:", persist)

	reader := bufio.NewReader(os.Stdin)

	go func() {
		events.OnEvent(context.Background(), "video.scanned", func(payload string) {

			var data map[string]interface{}
			err := json.Unmarshal([]byte(payload), &data)
			if err != nil {
				log.Error(err)
				return
			}

			id := data["id"].(string)
			username := data["username"].(string)

			fmt.Println("VIDEO >", username, id, data["views"], data["likes"])
			if persist {
				err := db.UpdateVideo(payload)
				if err != nil {
					log.Error(err)
				}
			}
		})
	}()

	go func() {
		events.OnEvent(context.Background(), "account.scanned", func(payload string) {

			var data map[string]interface{}
			err := json.Unmarshal([]byte(payload), &data)
			if err != nil {
				log.Error(err)
				return
			}

			fmt.Println("ACCOUNT >", data["username"], data["displayname"], data["followers"], data["avatar"])
			if persist {
				err := db.UpdateAccount(payload)
				if err != nil {
					log.Error(err)
				}
			}
		})
	}()

	go func() {
		err := api.Serve(":8080", events)
		log.Critical(err)
	}()

	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		events.Publish("account.scan.requested", text)
	}
}
