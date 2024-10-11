package service

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/skinkvi/systeminfo/internal/repository"
)

type HTTPService struct {
	repo *repository.Repository
	log  zerolog.Logger
}

func NewHTTPService(repo *repository.Repository, log zerolog.Logger) *HTTPService {
	return &HTTPService{
		repo: repo,
		log:  log,
	}
}

func (s *HTTPService) GetCPULogs(w http.ResponseWriter, r *http.Request) {
	const op = "service.GetCPULogs"
	cpuLogs, err := s.repo.GetCPUInfo()
	if err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to get CPU info")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cpuLogs); err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to encode CPU info")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *HTTPService) GetProccessLogs(w http.ResponseWriter, r *http.Request) {
	const op = "service.GetProccessLogs"
	proccessLogs, err := s.repo.GetProccessInfo()
	if err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to get process info")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(proccessLogs); err != nil {
		s.log.Error().Err(err).Msg(op + "Failed to encode process info")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
