package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	pb "wshhz.com/go-micro/Proto/Vessel"
)

type IRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// Our grpc service handler
type service struct {
	repo IRepository
}

// 实现接口方法
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	// 新建服务
	srv := micro.NewService(
		micro.Name("vessel"),
		micro.Version("latest"),
	)

	// 初始化服务
	srv.Init()

	// 注册服务
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	// 启动服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
