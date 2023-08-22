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
  ('1C4HJXDG3MW709024', 'Jeep', 'Wrangler Unlimited', 2020, 'Sport', 'Willys', 'Junebug', 'Hellayella', 'fdb93c'),
  ('5TDYK3DC9DS368862', 'Toyota', 'Sienna', 2013, 'Limited', '', '', 'Shoreline Blue Pearl', '4e5269');
