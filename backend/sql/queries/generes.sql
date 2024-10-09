-- name: CreateGenere :one
INSERT INTO generes (genere_id, genere_name)
VALUES ($1, $2)
RETURNING *;


-- name: GetGenereByName :one
SELECT genere_id, genere_name FROM generes WHERE genere_name=$1;