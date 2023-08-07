CREATE TABLE user (
  id       INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  name     TEXT NOT NULL,
  password TEXT NOT NULL,
  role     TEXT NOT NULL,
  created  TEXT NOT NULL,
  modified TEXT NOT NULL
);

CREATE TABLE task (
  id       INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  title    TEXT NOT NULL,
  status   TEXT NOT NULL,
  created  TEXT NOT NULL,
  modified TEXT NOT NULL
);

