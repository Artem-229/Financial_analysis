-- +goose up
-- +goose StatementBegin
CREATE TABLE users
(
    id       UUID PRIMARY KEY,
    name     TEXT NOT NULL,
    login    TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE transactions
(
    id    UUID PRIMARY KEY,
    item  TEXT    NOT NULL,
    price INTEGER NOT NULL,
    class TEXT
);
-- +goose StatementEnd

-- +goose down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE transactions;
-- +goose StatementEnd