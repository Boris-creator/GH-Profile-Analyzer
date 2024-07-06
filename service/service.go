package main

import (
	"context"
	"errors"
	"fmt"
	"ghs/analyzer"
	pb "ghs/proto"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GHServer struct {
	pb.UnimplementedGHAnalysisServiceServer
}

func fromOwnProfileInfo(info analyzer.OwnProfileInfo) pb.OwnProfileResponse {
	var developerType pb.EmploymentType
	switch info.Type {
	case analyzer.OPENSOURCE:
		developerType = pb.EmploymentType_OPENSOURCE
	case analyzer.WORK:
		developerType = pb.EmploymentType_WORK
	case analyzer.HOBBY:
		developerType = pb.EmploymentType_HOBBY
	}
	return pb.OwnProfileResponse{
		ContributionsDispersion: float32(info.ContributionsDispersion),
		Type:                    developerType,
	}
}

func (s *GHServer) OwnProfileInfo(ctx context.Context, req *pb.OwnProfileRequest) (*pb.OwnProfileResponse, error) {
	info, err := analyzer.ProfileInfo(req.Token)
	if errors.Is(err, analyzer.ErrNotEnoughData) {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := fromOwnProfileInfo(info)
	return &response, nil
}
func newServer() *GHServer {
	s := &GHServer{}
	return s
}

func main() {
	godotenv.Load("../.env")
	StartServer(os.Getenv("SERVICE_PORT"))
}

func StartServer(port string) {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGHAnalysisServiceServer(grpcServer, newServer())
	println("\033[33;1m Service started \033[0m")
	grpcServer.Serve(l)
}
