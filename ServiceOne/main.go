package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
	pb "wshhz.com/go-micro/Proto/Consignment"
	vesselProto "wshhz.com/go-micro/Proto/Vessel"
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
	// 接口对象
	repo IRepository

	// 代销客户端对象
	vesselClient vesselProto.VesselServiceClient
}

// 创建一个委托
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	//consignment, err := s.repo.Create(req)
	//if err != nil {
	//	return err
	//}
	//
	//res.Created = true
	//res.Consignment = consignment
	//
	//return nil

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
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

	// 新建
	vesselClient := vesselProto.NewVesselServiceClient("vessel", srv.Client())

	// 初始化服务
	srv.Init()

	// 注册服务
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// 启动服务
	if err := srv.Run(); err != nil {
		log.Println("error +:", err)
	}
}
