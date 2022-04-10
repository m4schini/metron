//scraper
package main

import (
	"context"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"google.golang.org/grpc"
	"miner-module/comm"
	pb "miner-module/proto"
	"miner-module/scraper"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var log = kapitol.NewLogger(os.Getenv("APP_NAME"), kapitol.Information)

var scr scraper.Scraper

func init() {
	s, err := scraper.NewScraper()
	if err != nil {
		log.Critical(err)
	}
	scr = s
}

func main() {
	log.Information(os.Getenv("APP_NAME"), "running...")
	log.Information("ID:", scr.ID())
	log.Information("ADDR:", os.Getenv("ADDR"))

	c := comm.NewClient(log, scr.ID())
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := c.KeepAlive(
			ctx,
			os.Getenv("COORDINATOR_ADDR"),
			os.Getenv("ADDR"),
		)
		if err != nil {
			log.Error(err)
		}
	}()

	lis, err := net.Listen("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Critical(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMinerServer(grpcServer, comm.NewServer(log, scr.TikTok()))
	log.Information("serving grpc on", os.Getenv("ADDR"))
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Critical(err)
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
	lis.Close()
}
