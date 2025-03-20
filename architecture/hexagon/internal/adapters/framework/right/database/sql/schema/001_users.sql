-- +goose Up
Create table users (
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  email TEXT NOT NULL UNIQUE,
  hashed_password TEXT NOT NULL
);

-- +goose Down
Drop table users;
