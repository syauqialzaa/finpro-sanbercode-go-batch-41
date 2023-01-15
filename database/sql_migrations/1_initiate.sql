-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(256),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);


CREATE TABLE books (
  id INT PRIMARY KEY NOT NULL,
  title VARCHAR(256),
  description VARCHAR(256),
  image_url VARCHAR(256),
  release_year INT,
  price VARCHAR(256),
  total_page INT,
  thickness VARCHAR(256),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  category_id INT
);
-- +migrate StatementEnd