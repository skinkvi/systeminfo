-- +goose Up
-- +goose StatementBegin
CREATE TABLE cpu_logs (
    id SERIAL PRIMARY KEY,
    cpu_info TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cpu_logs;
-- +goose StatementEnd
