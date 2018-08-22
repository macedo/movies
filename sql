CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v4();

CREATE TABLE movies (
  id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  title varchar(255) NOT NULL
);
