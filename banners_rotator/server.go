package main

import (
	"context"
	"fmt"
	"net"

	"github.com/DenKuzin33/OtusProject/banners_rotator/api"
	bannersselector "github.com/DenKuzin33/OtusProject/banners_selector"
	"google.golang.org/grpc"
)

type BannersService struct {
	api.UnimplementedBannersRotatorServer
}

func (service *BannersService) AddBannerToSlot(ctx context.Context, req *api.BannerOperationRequest) (*api.BannerOperationResponse, error) {
	result := AddBannerToSlot(ctx, int(req.BannerId), int(req.SlotId))
	return &api.BannerOperationResponse{Success: result == nil}, result
}

func (service *BannersService) RemoveBannerFromSlot(ctx context.Context, req *api.BannerOperationRequest) (*api.BannerOperationResponse, error) {
	result := RemoveBannerFromSlot(ctx, int(req.BannerId), int(req.SlotId))
	return &api.BannerOperationResponse{Success: result == nil}, result
}

func (service *BannersService) SelectBannerForShow(ctx context.Context, req *api.SelectBannerForShowRequest) (*api.SelectBannerForShowResponse, error) {
	banners, totalShowsCount, err := GetSlotBanners(ctx, req.SlotId, req.DemGroupId)
	if err != nil {
		return nil, err
	}
	// for _, banner := range banners {
	// 	if banner.demGroupId == 0 {
	// 		banner.showsCount++
	// 		//todo создать запись в shows_clicks_count
	// 		return &api.SelectBannerForShowResponse{BannerId: int64(banner.bannerId)}, nil
	// 	}

	// }
	selectedBanner := bannersselector.SelectNext(banners, totalShowsCount)
	UpdateBannerData(ctx, selectedBanner, int(req.SlotId), int(req.DemGroupId))

	return &api.SelectBannerForShowResponse{BannerId: int64(selectedBanner.Id)}, nil
}

func (service *BannersService) CountBannerClick(ctx context.Context, req *api.CountBannerClickRequest) (*api.BannerOperationResponse, error) {
	err := IncrementClicksCount(ctx, int(req.Req.SlotId), int(req.Req.BannerId), int(req.DemGroupId))
	return &api.BannerOperationResponse{Success: err == nil}, err
}

func main() {
	listener, err := net.Listen("tcp", ":55011")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()
	api.RegisterBannersRotatorServer(server, new(BannersService))

	if err := server.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
