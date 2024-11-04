-- name: CreateAuthor :one
INSERT INTO authors (id, name, average_rating, author_id, text_review_count, ratings_count)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;