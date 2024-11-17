-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user".users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    phone VARCHAR(15),
    birthdate TIMESTAMP,
    avatar VARCHAR(255),
    biodata TEXT,
    city VARCHAR(255),
    country VARCHAR(255),
    followers_total INT DEFAULT 0,
    following_total INT DEFAULT 0,
    private_account BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user".users;
-- +goose StatementEnd
