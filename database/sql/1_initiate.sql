-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE admins (
  id INT PRIMARY KEY NOT NULL,
  username VARCHAR(256),
  password VARCHAR(256),
  token VARCHAR(256),
  token_exp TIMESTAMP
);


CREATE TABLE members (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(256),
  email VARCHAR(256),
  team_id INT
);

CREATE TABLE projects (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(256),
  description TEXT,
  status VARCHAR(256),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE TABLE tasks (
  id INT PRIMARY KEY NOT NULL,
  project_id INT,
  team_id INT,
  name VARCHAR(256),
  description TEXT,
  assigned_member_id INT,
  status VARCHAR(256),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE TABLE teams (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(256)
);

-- +migrate StatementEnd