-- +goose Up
CREATE TABLE generes (
    genere_id SERIAL PRIMARY KEY,
    name TEXT UNIQUE
);



-- +goose Down
DROP TABLE generes;