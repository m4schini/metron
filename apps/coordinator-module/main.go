//coordinator
package main

import (
	"context"
	"coordinator-module/comm"
	pb "coordinator-module/proto"
	"coordinator-module/registry"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var log = kapitol.NewLogger(os.Getenv("APP_NAME"), kapitol.Debug)

func main() {
	log.Information(os.Getenv("APP_NAME"), "running...")

	lis, err := net.Listen("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Critical(err)
	}

	reg := registry.NewMapRegistry()

	grpcServer := grpc.NewServer()
	pb.RegisterCoordinatorServer(grpcServer, comm.NewServer(reg))
	log.Information("serving grpc on", os.Getenv("ADDR"))
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Critical(err)
		}
	}()

	log.Information("Retrieving scraper")
	scr, err := reg.Get(context.Background())
	if err != nil {
		log.Critical(err)
	}

	log.Information("RECEIVED SCRAPER:", scr.ID())

	time.Sleep(3 * time.Second)
	reg.Return(scr)
	log.Information(reg)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shutdown.
	fmt.Println("\nshutting down...")
	lis.Close()
}
