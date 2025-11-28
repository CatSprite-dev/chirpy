-- +goose Up
AlTER TABLE users
ADD COLUMN hashed_password TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
AlTER TABLE users
DROP COLUMN hashed_password;