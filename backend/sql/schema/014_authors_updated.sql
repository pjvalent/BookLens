-- +goose Up
ALTER TABLE authors
ADD COLUMN average_rating DECIMAL,
ADD COLUMN author_id BIGINT,
ADD COLUMN text_review_count BIGINT,
ADD COLUMN ratings_count BIGINT;

-- +goose Down
ALTER TABLE authors
DROP COLUMN average_rating,
DROP COLUMN author_id,
DROP COLUMN text_review_count,
DROP COLUMN ratings_count;