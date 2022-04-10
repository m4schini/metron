package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/integrii/flaggy"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
)

const (
	eventBusDb = 0

	E_ScanAccount        = "request-scan-account"
	E_ScanVideo          = "request-scan-video"
	E_DiscoveredVideo    = "discovered-video"
	E_UpdateAccount      = "update-account"
	E_UpdateVideoViews   = "update-video-views"
	E_UpdateVideoDetails = "update-video-details"
	E_VideoAvailable     = "update-video-available"
	E_VideoUnAvailable   = "update-video-unavailable"
)

var (
	eventBusAddr = "localhost:6379"
	eventBusPass = ""
)

func init() {
	flaggy.String(&eventBusAddr, "r", "redis", "Address of redis instance")
	flaggy.String(&eventBusPass, "", "password", "Password of redis instance")
	flaggy.Parse()
}

func main() {
	fmt.Println("redis:", eventBusAddr)

	hub := redis.NewClient(&redis.Options{
		Addr:     eventBusAddr,
		Password: eventBusPass,
		DB:       eventBusDb,
	})
	defer hub.Close()

	go handleMessages(hub)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("App is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
}

func handleMessages(hub *redis.Client) {
	for message := range hub.PSubscribe("*").Channel() {
		log.Printf(
			"[%20s] len(Payload)=%4d | Preview: %s\n",
			message.Channel,
			len(message.Payload),
			message.Payload[:int(math.Min(128, float64(len(message.Payload))))],
		)
		switch message.Channel {
		case E_ScanAccount:
			ScanAccount(hub, message)
		case E_ScanVideo:
			ScanVideo(hub, message)
		default:
			//log.Printf("[%s] NOT HANDELED\n", message.Channel)
		}
	}
}
