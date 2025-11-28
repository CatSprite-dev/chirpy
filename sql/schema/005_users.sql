-- +goose Up
AlTER TABLE users
ADD COLUMN is_chirpy_red BOOL NOT NULL DEFAULT false;

-- +goose Down
AlTER TABLE users
DROP COLUMN is_chirpy_red;

