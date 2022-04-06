package comm

import (
	"context"
	pb "coordinator-module/proto"
	"coordinator-module/registry"
	"coordinator-module/scraper"
	"log"
)

type Server struct {
	pb.UnimplementedCoordinatorServer
	registry registry.Registrar
}

func NewServer(registry registry.Registrar) *Server {
	s := new(Server)
	s.registry = registry
	return s
}

func (s *Server) CheckIn(ctx context.Context, in *pb.CheckInRequest) (*pb.CheckInResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		scr := scraper.NewScraper(in.GetMinerId())
		err := s.registry.Register(scr)
		if err != nil {
			return nil, err
		}
		log.Println("New checkin:", in.MinerId, s.registry)

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
		scr := scraper.NewScraper(in.GetMinerId())
		err := s.registry.Unregister(scr)
		if err != nil {
			return nil, err
		}
		log.Println("Checked out:", in.MinerId, s.registry)

		return &pb.CheckOutResponse{}, nil
	}
}
