-- +goose Up
-- +goose StatementBegin
ALTER TABLE Tags
Rename Column name to tag;
ALTER TABLE Tags
ALTER COLUMN tag TYPE int USING tag::integer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Tags
ALTER Column tag Type varchar(255) USING tag::varchar;
ALTER TABLE Tags
RENAME COlUMN tag TO name;
-- +goose StatementEnd
