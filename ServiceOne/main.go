package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
	pb "wshhz.com/go-micro/Proto"
)

// 仓库接口
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// 仓库对象
type Repository struct {
	consignments []*pb.Consignment
}

// 创建委托
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}

// 获取所有委托列表
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 服务对象
type service struct {
	repo IRepository
}

// 创建一个委托
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment

	return nil
}

// 获取委托
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments

	return nil
}

func main() {
	repo := &Repository{}

	// 新建一个服务对象
	srv := micro.NewService(
		micro.Name("consignment"),
		micro.Version("latest"),
	)

	// 初始化服务
	srv.Init()

	// 注册服务
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	// 启动服务
	if err := srv.Run(); err != nil {
		fmt.Println("error +:", err)
	}
}
