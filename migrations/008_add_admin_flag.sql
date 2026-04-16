-- +goose Up
ALTER TABLE users ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;

-- Make first user admin for demo (bektas)
UPDATE users SET is_admin = TRUE WHERE email = 'bektas.keldibayev@nu.edu.kz';

-- +goose Down
ALTER TABLE users DROP COLUMN IF EXISTS is_admin;
