-- +goose Up
ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
