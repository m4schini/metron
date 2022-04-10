package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis"
	tiktok "github.com/m4schini/tiktok-go"
	"log"
	"time"
)

const (
	E_ScanAccount        = "request-scan-account"
	E_ScanVideo          = "request-scan-video"
	E_DiscoveredVideo    = "discovered-video"
	E_UpdateAccount      = "update-account"
	E_UpdateVideoViews   = "update-video-views"
	E_UpdateVideoDetails = "update-video-details"
)

func prepareVideo(db *sql.DB, r *redis.Client, url string) {
	username, id := tiktok.ExtractUsernameAndId(url)

	_, err := db.Exec(`
			INSERT INTO tiktok.Video(id, postedBy, added, available)
				VALUES(?, ?, ?, ?)
		`,
		id, username, time.Now(), true,
	)

	if err != nil {
		r.Publish(E_DiscoveredVideo, url)
	}
}

func Updater(red *redis.Client) {
	db, err := NewDbConn()
	if err != nil {
		panic(err)
	}

	for update := range red.PSubscribe("update-*").Channel() {

		switch update.Channel {
		case E_UpdateVideoViews:
			updateVideoViews(db, red, update.Payload)
		case E_UpdateVideoDetails:
			updateVideoDetails(db, red, update.Payload)
		case E_UpdateAccount:
			updateAccount(db, update.Payload)
		}
	}
}

func updateVideoViews(db *sql.DB, r *redis.Client, payload string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return
	}

	value, ok := data["url"]
	if !ok {
		log.Println("url doesn't exist")
		return
	}
	url := value.(string)
	prepareVideo(db, r, url)
	_, id := tiktok.ExtractUsernameAndId(url)

	value, ok = data["views"]
	if !ok {
		log.Println("views doesn't exist")
		return
	}
	views := int(value.(float64))

	_, err = db.Exec(`
			UPDATE tiktok.Video
			SET views=?
			WHERE id=?
		`,
		views, id,
	)
	if err != nil {
		log.Println(err)
		return
	}
}

func updateVideoDetails(db *sql.DB, r *redis.Client, payload string) {
	var vid tiktok.Video
	err := json.Unmarshal([]byte(payload), &vid)
	if err != nil {
		return
	}

	prepareVideo(db, r, vid.URL)

	_, err = db.Exec(
		`UPDATE tiktok.Video
			   SET description=?, comments=?, shares=?, likes=?, available=?
			   WHERE id=?
				`,
		vid.Description, vid.CommentCount, vid.ShareCount, vid.LikeCount, vid.Available,
		vid.ID,
	)
}

func updateAccount(db *sql.DB, payload string) {
	var acc tiktok.Account
	err := json.Unmarshal([]byte(payload), &acc)
	if err != nil {
		return
	}

	_, err = db.Exec(
		`INSERT INTO tiktok.Account(username, displayname, bio, following, followers, likes)
						VALUES(?, ?, ?, ?, ?, ?)
					ON DUPLICATE KEY
	   				UPDATE displayname=?, bio=?, following=?, followers=?, likes=?`,
		acc.Username, acc.DisplayName, acc.Bio, acc.Following, acc.Followers, acc.Likes,
		acc.DisplayName, acc.Bio, acc.Following, acc.Followers, acc.Likes,
	)
}
