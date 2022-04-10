package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/integrii/flaggy"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	eventBusDb = 0
)

var (
	eventBusAddr = "localhost:6379"
	eventBusPass = ""
	targetUser   = "krawallklara2.0"
	cooldown     = "5m"
	database     = fmt.Sprintf(DbConnectionFormat, DbHost, DbPort, DbUser, DbPass, DbName)
)

func init() {
	flaggy.String(&eventBusAddr, "r", "redis", "Address of redis instance")
	flaggy.String(&eventBusPass, "", "password", "Password of redis instance")
	flaggy.String(&targetUser, "t", "target", "Target account name")
	flaggy.String(&database, "d", "database", "Database info")
	flaggy.String(&cooldown, "", "cooldown", "cooldown")
	flaggy.Parse()
}

func main() {
	fmt.Println("redis:", eventBusAddr)
	fmt.Println("target:", targetUser)

	hub := redis.NewClient(&redis.Options{
		Addr:     eventBusAddr,
		Password: eventBusPass,
		DB:       eventBusDb,
	})
	go Updater(hub)

	go func() {
		for {
			hub.Publish(E_ScanAccount, targetUser)
			scanAllVids(hub, targetUser)

			cd, err := time.ParseDuration(cooldown)
			if err != nil {
				panic(err)
			}
			log.Println("next scan in", cooldown)
			time.Sleep(cd)
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("App is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
	hub.Close()
}

func scanAllVids(r *redis.Client, acc string) {
	db, err := NewDbConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	log.Println("retrieving vids for", acc)
	rows, err := db.Query(`
		SELECT id
		FROM Video
		WHERE postedBy=?
	`, acc)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var vidId string
		err := rows.Scan(&vidId)
		if err != nil {
			log.Println(err)
			continue
		}

		r.Publish(E_ScanVideo, fmt.Sprintf("https://www.tiktok.com/@%s/video/%s", acc, vidId))
	}
}
