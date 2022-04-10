package comm

import (
	"context"
	"github.com/m4schini/kapitol-go"
	"github.com/m4schini/tiktok-go"
	"github.com/m4schini/tiktok-go/model"
	"github.com/m4schini/tiktok-go/util"
	pb "miner-module/proto"
	"time"
)

type Server struct {
	pb.UnimplementedMinerServer
	log    *kapitol.Logger
	source tiktok.TikTok
}

func NewServer(logger *kapitol.Logger, source tiktok.TikTok) *Server {
	s := new(Server)
	s.log = logger
	s.source = source

	return s
}

func (s Server) GetAccount(ctx context.Context, request *pb.AccountRequest) (*pb.Account, error) {
	s.log.Information("Received 'GetAccount' Request")

	select {
	case <-ctx.Done():
		s.log.Error("GetVideoDetails context deadline exceeded")
		return nil, ctx.Err()
	default:
		acc, err := s.source.GetAccount(request.GetName())
		if err != nil {
			s.log.Error(err)
			return nil, err
		}

		vids, err := s.source.GetLatestVideos(request.GetName())
		if err != nil {
			s.log.Error(err)
			s.log.Debug("using empty video previews array")
			vids = make([]*model.VideoPreview, 0)
		}

		pbVids := make([]*pb.VideoPreview, len(vids))

		for i, vid := range vids {
			username, id := util.ExtractUsernameAndId(vid.URL)

			pbVids[i] = &pb.VideoPreview{
				Url:      vid.URL,
				Username: username,
				Id:       id,
				Views:    int32(vid.Views),
			}
		}

		a := &pb.Account{
			Name:        acc.Username,
			DisplayName: acc.DisplayName,
			Bio:         acc.Bio,
			Following:   int32(acc.Following),
			Followers:   int32(acc.Followers),
			Likes:       int32(acc.Likes),
			Url:         acc.URL(),
			Videos:      pbVids,
			Timestamp:   time.Now().UnixMilli(),
		}
		s.log.Information("Returning", len(a.String()), "bytes")
		return a, nil
	}
}

func (s Server) GetVideoDetails(ctx context.Context, request *pb.VideoRequest) (*pb.VideoDetails, error) {
	s.log.Information("Received 'GetVideoDetails' Request")

	select {
	case <-ctx.Done():
		s.log.Error("GetVideoDetails context deadline exceeded")
		return nil, ctx.Err()
	default:
		vid, err := s.source.GetVideoByUrl(request.Url)
		if err != nil {
			s.log.Error(err)
			return nil, err
		}

		v := &pb.VideoDetails{
			Url:            vid.URL,
			VideoUrl:       vid.VideoURL,
			VideoTimestamp: vid.Timestamp,
			ThumbnailUrl:   vid.ThumbnailURL,
			Views:          int32(vid.Views),
			Likes:          int32(vid.Likes),
			Comments:       int32(vid.Comments),
			Shares:         int32(vid.Shares),
			AudioName:      vid.Audio,
			Description:    vid.Description,
			Timestamp:      time.Now().UnixMilli(),
		}
		s.log.Information("Returning", len(v.String()), "bytes")
		return v, nil
	}
}
