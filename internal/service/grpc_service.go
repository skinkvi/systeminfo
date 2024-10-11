package service

import (
	"context"

	"github.com/rs/zerolog"
	pb "github.com/skinkvi/protosinfo/gen/go/info"
	"github.com/skinkvi/systeminfo/internal/repository"
)

type GRPCService struct {
	pb.UnimplementedSystemInfoServiceServer
	repo *repository.Repository
	log  zerolog.Logger
}

func NewGRPCService(repo *repository.Repository, log zerolog.Logger) *GRPCService {
	return &GRPCService{
		repo: repo,
		log:  log,
	}
}

func (s *GRPCService) SendCPU(ctx context.Context, in *pb.CPURequest) (*pb.CPUResponse, error) {
	const op = "service.SendCPU"
	err := s.repo.LogCPUInfo(in.CpuInfo)
	if err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to log CPU info")
		return &pb.CPUResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.CPUResponse{Success: true, Message: "CPU info logged successfully"}, nil
}

func (s *GRPCService) SendCurrentProcesses(ctx context.Context, in *pb.ProcessesRequest) (*pb.ProcessesResponse, error) {
	const op = "service.SendCurrentProcesses"
	err := s.repo.LogProccessInfo(in.ProcessesInfo)
	if err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to log process info")
		return &pb.ProcessesResponse{Success: false, Message: err.Error()}, nil
	}

	return &pb.ProcessesResponse{Success: true, Message: "Process info logged successfully"}, nil
}
