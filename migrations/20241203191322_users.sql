-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    user_login VARCHAR(255) UNIQUE NOT NULL,
    pass_hash  VARCHAR
    secret_word VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
