package api

import (
	"commander-module/event"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "8080"
)

func Serve(addr string, events event.PubSub) error {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "moin")
	})
	r.HandleFunc("/scan/{account}", func(w http.ResponseWriter, r *http.Request) {
		handleScanAccount(events, w, r)
	})
	r.HandleFunc("/scan/{account}/{video}", func(w http.ResponseWriter, r *http.Request) {
		handleScanVideo(events, w, r)
	})
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
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

	return http.ListenAndServe(addr, nil)
}

func handleScanAccount(events event.PubSub, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println("Request to scan account:", vars["account"])
	events.Publish("cmd.scan.account", vars["account"])

	respond(w, 501, struct{}{})
}

func handleScanVideo(events event.PubSub, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println("Request to scan video:", vars["account"], "/", vars["video"])

	respond(w, 501, struct{}{})
}

func respond(w http.ResponseWriter, status int, response any) {
	jason, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(501)
		return
	}

	w.WriteHeader(status)
	w.Write(jason)
}
