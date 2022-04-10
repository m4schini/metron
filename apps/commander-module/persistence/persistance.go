package persistence

import (
	"commander-module/model"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var (
	DbConnectionFormat = "%s:%s@tcp(%s:%s)/%s"
)

func getConnectionString(host, port, password, username, database string) string {
	return fmt.Sprintf(DbConnectionFormat, username, password, host, port, database)
}

func getConnection() (*sql.DB, error) {
	return sql.Open("mysql", getConnectionString(
		os.Getenv("METRON_DB_HOST"),
		os.Getenv("METRON_DB_PORT"),
		os.Getenv("METRON_DB_PASS"),
		os.Getenv("METRON_DB_USER"),
		os.Getenv("METRON_DB_NAME"),
	))
}

type Persistence struct {
	db *sql.DB
}

func NewPersister() (*Persistence, error) {
	p := new(Persistence)
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	p.db = db
	return p, nil
}

func (p *Persistence) UpdateVideo(payload string) error {
	var video model.Video
	err := json.Unmarshal([]byte(payload), &video)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(`
			INSERT INTO tiktok.Video(id, postedBy, added, available)
				VALUES(?, ?, ?, ?)
		`,
		video.ID, video.Username, time.Now(), true,
	)

	_, err = p.db.Exec(`
			UPDATE tiktok.Video
			SET description=?, comments=?, views=?, shares=?, likes=?
			WHERE id=?
	`, video.Description, video.Comments, video.Views, video.Shares, video.Likes, video.ID)

	return err
}

func (p *Persistence) UpdateAccount(payload string) error {
	var acc model.Account
	err := json.Unmarshal([]byte(payload), &acc)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(
		`INSERT INTO tiktok.Account(username, displayname, bio, following, followers, likes)
						VALUES(?, ?, ?, ?, ?, ?)
					ON DUPLICATE KEY
	   				UPDATE displayname=?, bio=?, following=?, followers=?, likes=?`,
		acc.Username, acc.DisplayName, acc.Bio, acc.Following, acc.Followers, acc.Likes,
		acc.DisplayName, acc.Bio, acc.Following, acc.Followers, acc.Likes,
	)

	return err
}
