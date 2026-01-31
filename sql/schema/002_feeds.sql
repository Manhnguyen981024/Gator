-- +goose Up
CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    url VARCHAR(100) UNIQUE NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT user_id_fk
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);

-- +goose Down
DROP TABLE feeds