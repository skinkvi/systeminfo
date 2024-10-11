package main

import (
	"context"
	"time"

	"github.com/skinkvi/systeminfo/pkg/logger"
	"google.golang.org/grpc"

	pb "github.com/skinkvi/protosinfo/gen/go/info"
)

func main() {
	logger := logger.GetLogger()

	conn, err := grpc.Dial("server:11011", grpc.WithInsecure())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect")
	}
	defer conn.Close()
	logger.Info().Msg("Connected to server")

	client := pb.NewSystemInfoServiceClient(conn)
	logger.Info().Msg("Client connected")

	// пример отправки инфы о CPU
	cpuInfo := "CPU Info"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cpuResponse, err := client.SendCPU(ctx, &pb.CPURequest{CpuInfo: cpuInfo})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to send CPU info")
	}
	logger.Info().Msg(cpuResponse.Message)

	// пример отправки инфы о процессах
	processesInfo := "Processes Info"
	processesResponse, err := client.SendCurrentProcesses(ctx, &pb.ProcessesRequest{ProcessesInfo: processesInfo})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to send processes info")
	}
	logger.Info().Msg(processesResponse.Message)
}
