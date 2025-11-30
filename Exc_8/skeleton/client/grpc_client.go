package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.Dial(
		"localhost:4000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	fmt.Println("Requesting drinks")

	drinksResp, _ := c.client.GetDrinks(context.Background(), &emptypb.Empty{})

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%d  description:%q\n", d.Id, d.Name, d.Price, d.Description)
	}

	order := func(q int32) {
		req := &pb.OrderRequest{
			Items: []*pb.OrderItem{
				{DrinkId: 1, Quantity: q},
				{DrinkId: 2, Quantity: q},
				{DrinkId: 3, Quantity: q},
			},
		}
		c.client.OrderDrink(context.Background(), req)
	}

	fmt.Println("Ordering drinks")
	order(2)

	fmt.Println("Ordering another round")
	order(6)

	fmt.Println("Getting the bill...")
	totals, _ := c.client.GetOrders(context.Background(), &emptypb.Empty{})

	for _, t := range totals.Totals {
		name := drinksResp.Drinks[t.DrinkId-1].Name
		fmt.Printf("\t> Total: %d x %s\n", t.Quantity, name)
	}

	return nil
}
