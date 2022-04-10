package main

import (
	"github.com/go-redis/redis"
)

func ScanAccount(hub *redis.Client, message *redis.Message) {
	account := message.Payload

	acc, vids, err := GetAccountInfo(account)
	if err != nil {
		return
	}
	hub.Publish(E_UpdateAccount, JSON(acc))

	//db, err := NewDbConn()
	if err != nil {
		return
	}

	for _, vid := range vids {
		//id, username := tiktok.ExtractUsernameAndId(vid.Url)

		//TODO database connection in other service?
		//var count int
		//row := db.QueryRow(`SELECT COUNT(*) FROM tiktok.Video WHERE id=?`, id)
		//err := row.Scan(&count)
		//
		//if err != nil || count == 0 {
		//	hub.Publish(E_DiscoveredVideo, vid.Url)
		//}
		hub.Publish(E_UpdateVideoViews, JSON(vid))
	}
}

func ScanVideo(hub *redis.Client, message *redis.Message) {
	url := message.Payload

	vid, err := GetVideoDetails(url)
	if err != nil {
		return
	}

	hub.Publish(E_UpdateVideoDetails, JSON(vid))
}
