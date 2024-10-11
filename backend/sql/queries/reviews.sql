-- name: CreateReview :one
INSERT INTO reviews (id, created_at, updated_at, user_id, book_id, rating, review_text, spoiler_tag)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateReview :one
UPDATE reviews
SET 
    rating = $1,
    review_text = $2,
    updated_at = $3
WHERE 
    user_id = $4 AND book_id = $5
RETURNING *;

-- name: GetReviewByUserIDBookID :one
SELECT COUNT(*) FROM reviews WHERE user_id=$1 AND book_id=$2;

-- name: GetAllUserReviews :many
SELECT books.author, books.title, reviews.rating, reviews.review_text
FROM reviews JOIN books ON reviews.book_id = books.id
WHERE reviews.user_id = $1;