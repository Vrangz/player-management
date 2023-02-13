CREATE TABLE players (
    id          SERIAL PRIMARY KEY,
    username    varchar(255) UNIQUE NOT NULL,
    created_at  timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
