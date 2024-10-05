-- +goose Up
ALTER TABLE generes RENAME COLUMN name to genere_name;


-- +goose Down
ALTER TABLE generes RENAME COLUMN genere_name to name;