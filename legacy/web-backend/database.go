package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func connect() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", DbInfo+"?parseTime=true")
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("error occured while trying to close database connection")
		}
	}, nil
}

func connectQuery(query string, args ...interface{}) (*sql.Rows, func(), error) {
	db, closeConnection, err := connect()
	if err != nil {
		return nil, nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, nil, err
	}

	return rows, func() {
		closeConnection()
		rows.Close()
	}, nil
}

func getLineData(account, query string, limit int) ([]LineDataElement, error) {
	rows, Close, err := connectQuery(fmt.Sprintf(query, limit), account)
	if err != nil {
		return nil, err
	}
	defer Close()

	lineData := make([]LineDataElement, 0)
	for rows.Next() {
		point := new(LineDataElement)
		var timestamp time.Time

		if err := rows.Scan(&point.Y, &timestamp); err == nil {
			point.X = timestamp.Format(SQLTimestampFormat)
			lineData = append(lineData, *point)
		}
	}

	return lineData, nil
}

func getPieData(query string, args ...interface{}) (PieData, error) {
	rows, Close, err := connectQuery(query, args...)
	if err != nil {
		return nil, err
	}
	defer Close()

	pieData := make(PieData, 0)
	for rows.Next() {
		var id string
		var value int
		if err := rows.Scan(&id, &value); err != nil {
			log.Println(err)
		} else {
			pieData = append(pieData, PieDataElement{
				Id:    id,
				Value: value,
			})
		}
	}

	return pieData, nil
}

func getNetworkData(query string, rootNodeName string, args ...interface{}) (*NetworkData, error) {
	rows, Close, err := connectQuery(query, args...)
	if err != nil {
		return nil, err
	}
	defer Close()

	nodes := make([]NetworkNode, 0)
	for rows.Next() {
		var id string
		var value int
		if err := rows.Scan(&id, &value); err != nil {
			log.Println(err)
		} else {
			nodes = append(nodes, NetworkNode{
				Id:     id,
				Size:   value * 10,
				Height: value,
			})
		}
	}

	links := make([]NetworkLink, 0)
	for _, node := range nodes {
		links = append(links, NetworkLink{
			Source:   rootNodeName,
			Target:   node.Id,
			Distance: node.Height * 4,
		})
	}

	nodes = append(nodes, NetworkNode{
		Id:     rootNodeName,
		Size:   len(nodes),
		Height: len(nodes),
	})

	network := NetworkData{
		Nodes: nodes,
		Links: links,
	}

	return &network, nil
}

const getAccountQuery = `
	SELECT 
		postedBy, 
		COUNT(postedBy) AS 'count',
		SUM(views) AS 'views', 
		SUM(comments) AS 'comments', 
		SUM(shares) AS 'shares', 
		SUM(likes) AS 'likes'
	FROM Video FOR SYSTEM_TIME
	AS OF NOW()
	GROUP BY postedBy=?;
`

func GetAccount(account string) (*Account, error) {
	rows, Close, err := connectQuery(getAccountQuery, account)
	if err != nil {
		return nil, err
	}
	defer Close()

	acc := new(Account)
	acc.Name = account
	acc.LastUpdate = "---"
	acc.AvatarUrl = "//"

	if rows.Next() {
		var username string
		var count int
		var views int
		var comments int
		var shares int
		var likes int

		if err := rows.Scan(&username, &count, &views, &comments, &shares, &likes); err != nil {
			return nil, err
		} else {
			acc.VideoCount = count
			acc.Summary.Views.Value = views
			acc.Summary.Likes.Value = likes
			acc.Summary.Shares.Value = shares
			acc.Summary.Comments.Value = comments
		}
	}

	log.Println("retrieved account data")
	return acc, nil
}

const getAccountLikesQuery = `
	SELECT likes, ROW_END 
	FROM Account FOR SYSTEM_TIME ALL 
	WHERE username=? 
	ORDER BY ROW_END DESC 
	LIMIT %d;
`

func GetLikesHistory(account string, limit int) (*LineData, error) {
	data, err := getLineData(account, getAccountLikesQuery, limit)
	if err != nil {
		return nil, err
	}

	return &LineData{Id: "likes", Data: data}, nil
}

const getAccountFollowersQuery = `
	SELECT followers, ROW_END 
	FROM Account FOR SYSTEM_TIME ALL 
	WHERE username=? 
	ORDER BY ROW_END DESC 
	LIMIT %d;
`

func GetFollowersHistory(account string, limit int) (*LineData, error) {
	data, err := getLineData(account, getAccountFollowersQuery, limit)
	if err != nil {
		return nil, err
	}

	return &LineData{Id: "followers", Data: data}, nil
}

func GetVideo(account, videoId string) (*Video, error) {
	rows, Close, err := connectQuery(`
		SELECT description, views, comments, shares, likes, postedBy, account_Username AS 'mentioned'
		FROM Video v
		LEFT JOIN video_mentions_Account va ON (v.id = va.post_ID)
		WHERE id=?;
	`, videoId)
	if err != nil {
		return nil, err
	}
	defer Close()

	if rows.Next() {
		vid := new(Video)
		vid.Id = videoId

		if err := rows.Scan(
			&vid.Description,
			&vid.Views,
			&vid.Comments,
			&vid.Shares,
			&vid.Likes,
			&vid.PostedBy,
			&vid.Mentioned); err != nil {
			return nil, err
		}

		return vid, nil
	}
	return nil, errors.New("video wasn't found")
}
