package main

import (
	"context"
	"fmt"

	"github.com/DenKuzin33/OtusProject/banners_rotator/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":55011")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := api.NewBannersRotatorClient(conn)

	req := &api.BannerOperationRequest{
		SlotId:   1,
		BannerId: 1,
	}
	client.AddBannerToSlot(context.TODO(), req)
}
