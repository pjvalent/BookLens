-- name: CreateEmbedding :one
INSERT INTO book_embeddings (id, book_id, embedding)
VALUES ($1, $2, $3)
RETURNING *;