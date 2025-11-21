-- +goose Up
CREATE TABLE IF NOT EXISTS participants (
    competition_id UUID NOT NULL REFERENCES competitions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    points INT NOT NULL DEFAULT 0,
    days_read INT NOT NULL DEFAULT 0,
    minutes_total INT NOT NULL DEFAULT 0,
    last_log_date DATE NULL,

    PRIMARY KEY (competition_id, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS participants;
