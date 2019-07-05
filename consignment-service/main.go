package main

import (
	// 导如 protoc 自动生成的包
	"context"
	"errors"
	"fmt"
	pb "github.com/snake/test/consignment-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":50051"
)

//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment
}

//
// 定义微服务
//
type service struct {
	repo Repository
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
}

func CheckRepoConsignments(consignments []*pb.Consignment, consignment *pb.Consignment) bool {

	for _, num := range consignments {
		if num.Id == consignment.Id {
			return true
		}
	}
	return false
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	if CheckRepoConsignments(repo.consignments, consignment) {
		fmt.Println("Contains item")
		return nil, errors.New("already contains")
	}
	fmt.Println("新来一批货物:", consignment)
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{Created: true, Consignment: consignment}
	return resp, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	allConsignments := s.repo.GetAll()
	resp := &pb.Response{Consignments: allConsignments}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := Repository{}

	// 向 rRPC 服务器注册微服务
	// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterShippingServiceServer(server, &service{repo})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
