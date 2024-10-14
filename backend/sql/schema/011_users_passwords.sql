-- +goose Up
ALTER TABLE users ADD COLUMN user_password VARCHAR(255) NOT NULL DEFAULT 'admin';


-- +goose Down
ALTER TABLE users DROP COLUMN user_password;
