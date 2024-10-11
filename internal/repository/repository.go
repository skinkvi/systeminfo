package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Repository struct {
	db  *sqlx.DB
	log zerolog.Logger
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:  db,
		log: zerolog.Nop(),
	}
}

func (r *Repository) LogCPUInfo(cpuInfo string) error {
	const op = "repository.LogCPUInfo"
	_, err := r.db.Exec("INSERT INTO cpu_logs (cpu_info) VALUES ($1)", cpuInfo)
	if err != nil {
		r.log.Error().Err(err).Msg(op + "Failed to log CPU info")
		return err
	}
	return nil
}

func (r *Repository) LogProccessInfo(proccessInfo string) error {
	const op = "repository.LogProccessInfo"
	_, err := r.db.Exec("INSERT INTO processes_logs (processes_info) VALUES ($1)", proccessInfo)
	if err != nil {
		r.log.Error().Err(err).Msg(op + "Failed to log proccess info")
		return err
	}
	return nil
}

func (r *Repository) GetCPUInfo() ([]map[string]interface{}, error) {
	const op = "repository.GetCPUInfo"
	rows, err := r.db.Query("SELECT * FROM cpu_logs")
	if err != nil {
		r.log.Error().Err(err).Msg(op + "Failed to get CPU info")
		return nil, err
	}
	defer rows.Close()

	var cpuLogs []map[string]interface{}
	for rows.Next() {
		var id int
		var cpuInfo string
		var timestamp time.Time
		if err := rows.Scan(&id, &cpuInfo, &timestamp); err != nil {
			r.log.Error().Err(err).Msg(op + "Failed to scan CPU info")
			return nil, err
		}
		cpuLogs = append(cpuLogs, map[string]interface{}{
			"id":        id,
			"cpu_info":  cpuInfo,
			"timestamp": timestamp,
		})
	}
	return cpuLogs, nil
}

func (r *Repository) GetProccessInfo() ([]map[string]interface{}, error) {
	const op = "repository.GetProccessInfo"
	rows, err := r.db.Query("SELECT * FROM processes_logs")
	if err != nil {
		r.log.Error().Err(err).Msg(op + "Failed to get proccess info")
		return nil, err
	}
	defer rows.Close()

	var proccessLogs []map[string]interface{}
	for rows.Next() {
		var id int
		var proccessInfo string
		var timestamp time.Time
		if err := rows.Scan(&id, &proccessInfo, &timestamp); err != nil {
			r.log.Error().Err(err).Msg(op + "Failed to scan proccess info")
			return nil, err
		}
		proccessLogs = append(proccessLogs, map[string]interface{}{
			"id":            id,
			"proccess_info": proccessInfo,
			"timestamp":     timestamp,
		})
	}
	return proccessLogs, nil
}
