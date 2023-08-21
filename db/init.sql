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
  colorvalue varchar
);
