-- +goose Up
CREATE TABLE posts(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(500) NOT NULL,
    url VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(1000) NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL
);

-- +goose Down
DROP TABLE posts;