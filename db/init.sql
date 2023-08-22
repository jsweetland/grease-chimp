CREATE DATABASE gc;
GRANT ALL PRIVILEGES ON DATABASE gc TO gcuser;

\c gc;

CREATE TABLE vehicles (
  vin varchar primary key,
  make varchar,
  model varchar,
  year int,
  trim varchar,
  package varchar,
  nickname varchar,
  colorname varchar,
  colorhex varchar
);

/* insert test data */
INSERT INTO vehicles (vin, make, model, year, trim, package, nickname, colorname, colorhex) VALUES
  ('abc', 'Jeep', 'Wrangler Unlimited', 2020, 'Sport', 'Willys', 'Junebug', 'Hellayella', 'fdb93c'),
  ('def', 'Toyota', 'Sienna', 2013, 'Limited', NULL, NULL, NULL, NULL);
