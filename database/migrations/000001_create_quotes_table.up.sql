BEGIN;
CREATE TABLE quotes (
  id VARCHAR(32),
  name VARCHAR(50) NOT NULL,
  service VARCHAR(50) NOT NULL,
  deadline INT NOT NULL,
  price DECIMAL(12, 2) NOT NULL,
  PRIMARY KEY(id)
);
COMMIT;