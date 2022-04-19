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