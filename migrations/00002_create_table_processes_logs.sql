-- +goose Up
-- +goose StatementBegin
CREATE TABLE processes_logs (
    id SERIAL PRIMARY KEY,
    processes_info TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE processes_logs;
-- +goose StatementEnd
