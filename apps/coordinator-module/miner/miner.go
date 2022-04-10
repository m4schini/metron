package miner

import (
	"context"
	"coordinator-module/model"
	pb "coordinator-module/proto"
	tiktokmodel "github.com/m4schini/tiktok-go/model"
	"github.com/m4schini/tiktok-go/util"
	"time"
)

type Miner interface {
	ID() string
	GetAccount(username string) (*model.Account, []*model.Video, error)
	GetVideoDetails(username, videoId string) (*model.Video, error)
}

type miner struct {
	id     string
	Addr   string
	online bool

	client pb.MinerClient
}

func NewMiner(id string, addr string) *miner {
	m := new(miner)
	m.id = id
	m.Addr = addr
	m.online = false

	client, err := Connect(addr)
	if err != nil {
		return nil
	}
	m.client = client

	return m
}

func (m *miner) ID() string {
	return m.id
}

func (m *miner) GetAccount(username string) (*model.Account, []*model.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	acc, err := m.client.GetAccount(ctx, &pb.AccountRequest{Name: username})
	cancel()
	if err != nil {
		return nil, nil, err
	}

	account := new(model.Account)
	account.Username = acc.Name
	account.DisplayName = acc.DisplayName
	account.Following = int(acc.Following)
	account.Followers = int(acc.Followers)
	account.Likes = int(acc.Likes)
	account.Bio = acc.Bio

	vids := make([]*model.Video, 0, len(acc.Videos))
	for _, video := range acc.Videos {
		vids = append(vids, &model.Video{
			URL:      video.Url,
			ID:       video.Id,
			Username: video.Username,
			Views:    int(video.Views),
		})
	}

	return account, vids, nil
}

func (m *miner) GetVideoDetails(username, videoId string) (*model.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	v, err := m.client.GetVideoDetails(ctx, &pb.VideoRequest{
		Url: tiktokmodel.ToVideoURL(username, videoId),
	})
	cancel()
	if err != nil {
		return nil, err
	}

	video := new(model.Video)
	video.URL = v.Url
	username, id := util.ExtractUsernameAndId(v.Url)
	video.ID = id
	video.Username = username
	video.VideoURL = v.VideoUrl
	video.Timestamp = v.VideoTimestamp
	video.ThumbnailURL = v.ThumbnailUrl

	video.Views = int(v.Views)
	video.Likes = int(v.Likes)
	video.Comments = int(v.Comments)
	video.Shares = int(v.Shares)
	video.Audio = v.AudioName
	video.VideoLength = 0
	video.Description = v.Description

	return video, nil
}
