-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    user_login VARCHAR(255) PRIMARY KEY,
    pass_hash  VARCHAR,
    secret_word VARCHAR(255)
);

CREATE TABLE playlist
(   
    id SERIAL PRIMARY KEY,
    user_login VARCHAR(255) NOT NULL,
    playing BOOLEAN DEFAULT FALSE,
    prev_id INT UNIQUE,
    next_id INT UNIQUE,
    title TEXT NOT NULL,
    article TEXT NOT NULL,
    duration BIGINT NOT NULL,
    FOREIGN KEY (user_login) REFERENCES users(user_login) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE playlist;
-- +goose StatementEnd
