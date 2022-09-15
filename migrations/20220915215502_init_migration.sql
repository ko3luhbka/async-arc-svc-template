-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA if NOT EXISTS test;
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
