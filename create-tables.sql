DROP TABLE IF EXISTS usertype;
CREATE TABLE usertype (
  id         INT AUTO_INCREMENT NOT NULL,
  user      VARCHAR(128) NOT NULL,
  administrator      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users 
  (username, command) 
VALUES
  ('Example action 1', '!dosomething'),
  ('Example action 1', '!dosomethingelse');

DROP TABLE IF EXISTS actions;
CREATE TABLE actions (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  command     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO actions 
  (title, command) 
VALUES 
  ('Example action 1', '!dosomething'),
  ('Example action 2', '!dosomethingelse');

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id    INT AUTO_INCREMENT NOT NULL,
  username  VARCHAR(128) NOT NULL,
  usertype  FOREIGN KEY REFERENCES users(id),
  PRIMARY KEY (`id`)
);

INSERT INTO users 
  (username, usertype) 
VALUES
  ('Example user 1', ''),
  ('Example user 2', '!dosomethingelse');