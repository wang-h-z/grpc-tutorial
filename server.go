package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wang_h_z/grpc-tutorial/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(req *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Espresso",
		},
		&pb.Item{
			Id:   "3",
			Name: "Cappuccino",
		},
	}
	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	return nil
}

func (s *server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}
func (s *server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	//set up a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
