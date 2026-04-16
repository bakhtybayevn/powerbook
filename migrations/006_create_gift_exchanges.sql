-- +goose Up
CREATE TABLE gift_exchanges (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    competition_id UUID NOT NULL REFERENCES competitions(id) ON DELETE CASCADE,
    giver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    gift_description TEXT NOT NULL DEFAULT '',
    giver_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    receiver_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(competition_id, giver_id)
);
CREATE INDEX idx_gift_comp ON gift_exchanges(competition_id);
CREATE INDEX idx_gift_giver ON gift_exchanges(giver_id);
CREATE INDEX idx_gift_receiver ON gift_exchanges(receiver_id);

-- +goose Down
DROP TABLE IF EXISTS gift_exchanges;
