-- +goose up
-- +goose StatementBegin
CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    name          TEXT        NOT NULL,
    login         TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL
);

CREATE TABLE transactions
(
    id         UUID PRIMARY KEY,
    user_id    UUID      NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    item       TEXT      NOT NULL,
    price      INTEGER   NOT NULL,
    class      TEXT,
    status     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE outcomes
(
    id       UUID PRIMARY KEY,
    user_id  UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    analysis TEXT NOT NULL
)

CREATE INDEX transactions_user_id_idx ON transactions (user_id);

-- +goose StatementEnd

-- +goose down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE transactions;
-- +goose StatementEnd