package miner

import (
	"context"
	pb "coordinator-module/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(addr string) (pb.MinerClient, error) {
	grpcConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewMinerClient(grpcConn)

	go waitForDeadConnection(grpcConn, func() {})

	return grpcClient, nil
}

func waitForDeadConnection(conn *grpc.ClientConn, onDead func()) {
	for func() bool {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if conn.WaitForStateChange(ctx, conn.GetState()) {
			state := conn.GetState()

			// if state changed to 'shutdown' or 'idle'
			if state == 4 || state == 0 {

				onDead()
				return false
			}
		}

		return true
	}() {
	}
}
