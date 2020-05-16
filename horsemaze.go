package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"github.com/lakkinzimusic/horse_maze/api/version"
	pb "github.com/lakkinzimusic/horse_maze/proto/consignment"
	vesselProto "github.com/lakkinzimusic/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

func main() {
	repo := &Repository{}
	log.Printf(
		"Starting the service on port %s...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	port := os.Getenv("PORT")
	port = ":50051"
	if port == "" {
		log.Fatal("Port is not set.")
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

	// This name must match the package name given in your protobuf definition
	// micro.Name("horse-maze"),
	)

	// Init will parse the command line flags.
	srv.Init()
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", srv.Client())
	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	// app := server.NewApp()
	// if err := app.Run(port); err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
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

	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
