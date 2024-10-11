package app

import (
	"net"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/skinkvi/systeminfo/internal/config"
	"github.com/skinkvi/systeminfo/internal/repository"
	"github.com/skinkvi/systeminfo/internal/service"
	"github.com/skinkvi/systeminfo/pkg/db"
	"google.golang.org/grpc"

	pb "github.com/skinkvi/protosinfo/gen/go/info"
)

type App struct {
	cfg     *config.Config
	repo    *repository.Repository
	log     zerolog.Logger
	grpcSrv *grpc.Server
	httpSrv *http.Server
}

func NewApp(cfg *config.Config, log *zerolog.Logger) (*App, error) {
	dbConn, err := db.NewDB(cfg, log)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(dbConn)
	grpcService := service.NewGRPCService(repo, *log)

	httpService := service.NewHTTPService(repo, *log)

	grpcSrv := grpc.NewServer()
	pb.RegisterSystemInfoServiceServer(grpcSrv, grpcService)

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/cpu", httpService.GetCPULogs)
	httpMux.HandleFunc("/proccess", httpService.GetProccessLogs)

	httpSrv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: httpMux,
	}

	return &App{
		cfg:     cfg,
		repo:    repo,
		log:     *log,
		grpcSrv: grpcSrv,
		httpSrv: httpSrv,
	}, nil
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", ":"+a.cfg.GRPCPort)
	if err != nil {
		a.log.Error().Err(err).Msg("Failed to listen")
		return err
	}

	go func() {
		if err := a.grpcSrv.Serve(lis); err != nil {
			a.log.Error().Err(err).Msg("Failed to serve")
		}
	}()

	if err := a.httpSrv.ListenAndServe(); err != nil {
		a.log.Error().Err(err).Msg("Failed to serve")
		return err
	}

	return nil
}
