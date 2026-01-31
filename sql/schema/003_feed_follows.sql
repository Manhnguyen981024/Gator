-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    CONSTRAINT user_id_fk FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT feed_id_fk FOREIGN KEY(feed_id) REFERENCES feeds(id),
    CONSTRAINT unique_user_feed UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;