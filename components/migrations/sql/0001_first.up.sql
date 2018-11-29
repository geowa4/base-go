CREATE TABLE IF NOT EXISTS foos (
  id serial,
  name text,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bars (
  id serial,
  foo_id integer,
  value integer,

  PRIMARY KEY (id),
  FOREIGN KEY (foo_id) REFERENCES foos(id) ON DELETE CASCADE
);
