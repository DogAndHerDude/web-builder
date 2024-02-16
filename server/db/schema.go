package db

var schema = `
  CREATE TABLE IF NOT EXISTS user (
    id TEXT PRIMARY KEY NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
  ); 

  CREATE TABLE IF NOT EXISTS site_dom_tree (
    id TEXT PRIMARY KEY NOT NULL,
    tree JSON NOT NULL
  );

  CREATE TABLE IF NOT EXISTS site (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    settings JSON NOT NULL,
    tree_id TEXT NOT NULL,
    user_id TEXT NOT NULL,

    CONSTRAINT site_site_dom_tree_fkey FOREIGN KEY ("tree_id") REFERENCES site_dom_tree ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT site_user_fkey FOREIGN KEY ("user_id") REFERENCES user ("id") ON DELETE RESTRICT ON UPDATE CASCADE
  );

  CREATE INDEX IF NOT EXISTS user_idx on user ("id");

  CREATE INDEX IF NOT EXISTS site_idx on site ("id");

  CREATE INDEX IF NOT EXISTS site_dom_trees_idx on site_dom_tree ("id");
`

