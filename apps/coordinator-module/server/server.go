package server

import (
	"context"
	"coordinator-module/miner"
	pb "coordinator-module/proto"
	"coordinator-module/registry"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedCoordinatorServer
	registry registry.Registrar
	closeF   func()
}

func NewServer(registry registry.Registrar) *Server {
	s := new(Server)
	s.registry = registry
	s.closeF = func() {}
	return s
}

func (s *Server) Serve(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoordinatorServer(grpcServer, s)

	s.closeF = func() {
		listener.Close()
	}
	return grpcServer.Serve(listener)
}

func (s *Server) Close() {
	s.closeF()
}

func (s *Server) CheckIn(ctx context.Context, in *pb.CheckInRequest) (*pb.CheckInResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		m := miner.NewMiner(in.GetMinerId(), in.GetAddress())
		err := s.registry.Register(m)
		if err != nil {
			return &pb.CheckInResponse{
				Success: false,
			}, err
		}
		log.Println("New checkin:", in.MinerId)

		return &pb.CheckInResponse{
			Success: true,
		}, nil
	}
}

func (s *Server) CheckOut(ctx context.Context, in *pb.CheckOutRequest) (*pb.CheckOutResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		err := s.registry.Unregister(in.GetMinerId())
		if err != nil {
			return nil, err
		}
		log.Println("Checked out:", in.MinerId)

		return &pb.CheckOutResponse{}, nil
	}
}
