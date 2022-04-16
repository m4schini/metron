package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"web-api/data"
	"web-api/model"
)

func handleResponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.WriteHeader(500)
		log.Println("[RESPONSE ERROR]", err)

		w.Header().Set("Cache-Control", "max-age=0")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf("{ error: \"%s\" }", err.Error()))
		return
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		log.Println("[RESPONSE ERROR]", err)

		w.Header().Set("Cache-Control", "max-age=0")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, fmt.Sprintf("{ error: \"%s\" }", err.Error()))
		return
	}

	w.Header().Set("Cache-Control", "max-age=150")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(jsonStr))
}

func AccountHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	acc, err := data.GetAccount(vars["account"])

	handleResponse(writer, acc, err)
}

func VideoHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	vid, err := data.GetVideo(vars["account"], vars["video"])
	handleResponse(writer, vid, err)
}

func AccountLikesHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	limit := 0
	limitStr := request.URL.Query().Get("limit")
	if limitStr == "" {
		limit = model.AccHistoricDataDefAmount
	} else {
		i, err := strconv.ParseInt(limitStr, 10, 32)
		if err != nil || i < 0 {
			limit = model.AccHistoricDataDefAmount
		} else {
			limit = int(i)
		}
	}

	likes, err := data.GetLikesHistory(vars["account"], limit)
	handleResponse(writer, likes, err)
}

func AccountFollowersHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	limit := 0
	limitStr := request.URL.Query().Get("limit")
	if limitStr == "" {
		limit = model.AccHistoricDataDefAmount
	} else {
		i, err := strconv.ParseInt(limitStr, 10, 32)
		if err != nil || i < 0 {
			limit = model.AccHistoricDataDefAmount
		} else {
			limit = int(i)
		}
	}

	followers, err := data.GetFollowersHistory(vars["account"], limit)
	handleResponse(writer, followers, err)
}

func AccountTagsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	log.Println("tagHandler", vars["account"])

	pie, err := data.GetPieData(`
		SELECT name, COUNT(*) AS amount
		FROM Video v
		JOIN (
		    SELECT DISTINCT *
		    FROM Video_has_Tag
		    ) vt ON v.id = vt.post_ID
		JOIN Tag t ON vt.tag_ID = t.id
		WHERE postedBy=?
		GROUP BY name
		ORDER BY amount DESC;
	`, vars["account"])
	handleResponse(writer, pie, err)
}

func AccountMentionsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	network, err := data.GetNetworkData(`	
		SELECT account_Username, COUNT(*) AS 'count'
		FROM Video v
		JOIN (
    		SELECT DISTINCT *
    		FROM video_mentions_Account
		) va ON (v.id = va.post_ID)
		WHERE postedBy = ?
		GROUP BY account_Username;
	`, vars["account"], vars["account"])

	handleResponse(writer, network, err)
}

func getTimeRangeData(account, query string) (model.TimeRangeData, error) {
	db, err := sql.Open("mysql", model.DbInfo+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	log.Println(account)
	rows, err := db.Query(query, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activity := make(map[string]int)

	for rows.Next() {
		var likes int
		var timestamp time.Time
		if err := rows.Scan(&likes, &timestamp); err != nil {
			log.Println(err)
		} else {
			activity[timestamp.Format("2006-01-02")]++
		}
	}

	timeRangeData := make([]model.TimeRangeElement, 0)
	for k, v := range activity {
		timeRangeData = append(timeRangeData, model.TimeRangeElement{
			Day:   k, //"2019-07-04T13:33:03.969Z"
			Value: v,
		})
	}

	return timeRangeData, nil
}

func AccountActivtyHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	data, err := getTimeRangeData(vars["account"], `
		SELECT id, MIN(ROW_END) AS 'timestamp'
		FROM Video FOR SYSTEM_TIME ALL
		WHERE postedBy=?
		GROUP BY id;
	`)
	handleResponse(writer, data, err)
}

func getAccountVideos(account, query string) ([]*model.Video, error) {
	db, err := sql.Open("mysql", model.DbInfo+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	log.Println(account)
	rows, err := db.Query(query, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vids := make([]*model.Video, 0)
	for rows.Next() {
		vid := new(model.Video)
		if err := rows.Scan(
			&vid.Id,
			&vid.Description,
			&vid.Views,
			&vid.Comments,
			&vid.Shares,
			&vid.Likes,
			&vid.PostedBy,
		); err != nil {
			log.Println(err)
		} else {
			vids = append(vids, vid)
		}
	}

	return vids, nil
}
func AccountVideosHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	data, err := getAccountVideos(vars["account"], `
		SELECT *
		FROM Video
		WHERE postedBy=?
	`)
	handleResponse(writer, data, err)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "moin")
	})
	r.HandleFunc("/{account}", AccountHandler)
	r.HandleFunc("/{account}/likes", AccountLikesHandler)
	r.HandleFunc("/{account}/followers", AccountFollowersHandler)
	r.HandleFunc("/{account}/tags", AccountTagsHandler)
	r.HandleFunc("/{account}/mentions", AccountMentionsHandler)
	r.HandleFunc("/{account}/activity", AccountActivtyHandler)
	r.HandleFunc("/{account}/videos", AccountVideosHandler)
	r.HandleFunc("/{account}/{video}", VideoHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = model.DefaultPort
	}

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Serving on :%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
