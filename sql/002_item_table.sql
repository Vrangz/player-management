CREATE TABLE items (
    id          SERIAL PRIMARY KEY,
    player_id   int NOT NULL,
    name        varchar(255) NOT NULL,
    quantity    int NOT NULL,
    CONSTRAINT UQ_player_item UNIQUE(player_id, name),
    CONSTRAINT FK_players_items FOREIGN KEY (player_id) REFERENCES players (id) ON DELETE CASCADE
);

CREATE INDEX player_id_idx ON items (player_id);
CREATE INDEX name_idx ON items (name);
