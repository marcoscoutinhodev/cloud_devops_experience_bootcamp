ALTER SYSTEM SET max_connections = 500;

CREATE TABLE IF NOT EXISTS comments (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(70) NOT NULL,
  email VARCHAR(320) NOT NULL,
  comment VARCHAR(512) NOT NULL,
  PRIMARY KEY (id)
);
