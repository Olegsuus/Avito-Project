-- +goose Up
-- +goose StatementBegin
ALTER TABLE Users ADD COLUMN password VARCHAR(64) NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Users DROP COLUMN password;
-- +goose StatementEnd
