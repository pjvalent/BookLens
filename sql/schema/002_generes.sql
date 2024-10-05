-- +goose Up
CREATE TABLE generes (
    genere_id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);



-- +goose Down
DROP TABLE generes;