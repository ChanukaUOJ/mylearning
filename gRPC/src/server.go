package main

import (
	"context";
	"log";
	"net";

	pb "github.com/ChanukaUOJ/gotuto/gRPC/proto_gen"

	"google.golang.org/grpc";
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, stream grpc.ServerStreamingServer[pb.Menu]) error {

	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Espresso",
		},
		&pb.Item{
			Id:   "2",
			Name: "Latte",
		},
		&pb.Item{
			Id:   "3",
			Name: "Cappuccino",
		},
	}

	menu := &pb.Menu{
		Items: items,
	}

	return stream.Send(menu)
}

func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABCD",
	}, nil
}

func (s *server) GetOrderStatus(context.Context, *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: "ABCD",
		Status:  "IN PROGRESS",
	}, nil
}

func main(){
	// setup a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")

	if err != nil{
		log.Fatalf("failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve : %v", err)
	}
}