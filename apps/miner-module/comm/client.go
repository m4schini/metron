package comm

import (
	"context"
	"errors"
	"fmt"
	"github.com/m4schini/kapitol-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "miner-module/proto"
	"time"
)

const (
	minRetryCooldown = 2 * time.Second
	maxRetryCooldown = 2 * time.Minute
)

type Client struct {
	id          string
	coordinator pb.CoordinatorClient
	logger      *kapitol.Logger
	closeFunc   func()
}

func NewClient(logger *kapitol.Logger, id string) *Client {
	c := new(Client)
	c.id = id
	c.logger = logger
	c.closeFunc = func() {}
	return c
}

func (c *Client) KeepAlive(ctx context.Context, targetAddr, callbackAddr string) error {
	if targetAddr == "" {
		return errors.New("target addr is required")
	}
	if callbackAddr == "" {
		return errors.New("callback addr is required")
	}

	done := false

	select {
	case <-ctx.Done():
		done = true
		break
	default:
		cooldown := minRetryCooldown
		for !done {
			closed, err := c.CheckIn(ctx, targetAddr, callbackAddr, c.id)
			if err != nil {
				c.logger.Error(err)
				c.logger.Debug("Trying to reconnect in", cooldown)
				time.Sleep(cooldown)

				cooldown = cooldown * 2
				if cooldown > maxRetryCooldown {
					cooldown = maxRetryCooldown
				}
				continue
			}

			<-closed
		}
	}

	return ctx.Err()
}

//CheckIn tries to check in with coordinator, returns channel that receives if connection is closing
func (c *Client) CheckIn(ctx context.Context, targetAddr, callbackAddr, id string) (<-chan struct{}, error) {
	c.logger.Debug("dialing grpc server on", targetAddr)
	conn, err := grpc.Dial(targetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c.logger.Debug("creating coordinator client")
	client := pb.NewCoordinatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	c.logger.Debug("checking in with coordinator")
	_, err = client.CheckIn(ctx, &pb.CheckInRequest{
		MinerId: id,
		Address: callbackAddr,
	})
	cancel()

	c.coordinator = client

	closeCh := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			for {
				if conn.WaitForStateChange(context.Background(), conn.GetState()) {
					state := conn.GetState()

					fmt.Println("CONNECTION STATE CHANGED TO", state)
					// if state changed to 'shutdown' or 'idle'
					if state == 4 || state == 0 {
						closeCh <- struct{}{}
						close(closeCh)
						break
					}
				}
			}
		}

	}()
	c.closeFunc = func() {
		c.CheckOut()
		conn.Close()
		closeCh <- struct{}{}
	}
	return closeCh, err
}

func (c *Client) CheckOut() error {
	if c.coordinator == nil {
		return errors.New("coordinator is closed")
	}

	c.coordinator.CheckOut(context.Background(), &pb.CheckOutRequest{MinerId: c.id})
	return nil
}

func (c *Client) Close() {
	if c != nil {
		c.closeFunc()
		c.coordinator = nil
	}
}
