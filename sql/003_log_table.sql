CREATE TABLE logs (
    id          SERIAL PRIMARY KEY,
    created_at  timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    msg         JSON NOT NULL
);
