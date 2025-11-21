-- +goose Up
CREATE TABLE IF NOT EXISTS reading_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    minutes INT NOT NULL,
    source TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reading_user ON reading_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_reading_ts ON reading_logs(timestamp);

-- +goose Down
DROP TABLE IF EXISTS reading_logs;
