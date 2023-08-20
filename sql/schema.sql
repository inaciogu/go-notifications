DROP DATABASE IF EXISTS notifications;
DROP TABLE IF EXISTS notifications;

CREATE DATABASE notifications;

CREATE TABLE notifications (
  id varchar(255) NOT NULL,
  recipient_id varchar(255) NOT NULL,
  body varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  read_at timestamp NULL,
  deleted_at timestamp NULL,
  PRIMARY KEY (id)
);