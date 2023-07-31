-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id serial PRIMARY KEY,
  first_name VARCHAR(64) NOT NULL CHECK (first_name <> ''),
  last_name VARCHAR(64) NOT NULL CHECK (last_name <> ''),
  email VARCHAR(64) UNIQUE NOT NULL CHECK (email <> ''),
  password VARCHAR(255) NOT NULL CHECK (octet_length(password) <> 0),
  phone VARCHAR(64) UNIQUE,
  profile_picture_url VARCHAR(255),
  city VARCHAR(64),
  birthday VARCHAR(64),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  last_login_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
  post_id serial PRIMARY KEY,
  user_id serial NOT NULL REFERENCES users (user_id),
  content TEXT NOT NULL CHECK (content <> ''),
  image_url VARCHAR(1024) CHECK (image_url <> ''),
  likes_count INTEGER NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE post_likes (
  like_id serial PRIMARY KEY,
  post_id int NOT NULL,
  user_id int NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  UNIQUE (post_id, user_id),
  FOREIGN KEY (user_id) REFERENCES users (user_id),
  FOREIGN KEY (post_id) REFERENCES posts (post_id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_likes;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd