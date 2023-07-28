-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id UUID PRIMARY KEY  ,
  firs_name VARCHAR(64) NOT NULL CHECK (firs_name <> ''),
  last_name VARCHAR(64) NOT NULL CHECK (last_name <> ''),
  email VARCHAR(64) UNIQUE NOT NULL CHECK (email <> ''),
  password VARCHAR(255) NOT NULL CHECK (octet_length(password) <> 0),
  phone VARCHAR(64) UNIQUE,
  profile_picture_url VARCHAR(255),
  city VARCHAR(64),
  birthday VARCHAR(64),
  created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
  updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
  last_login_at   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sources;
-- +goose StatementEnd
