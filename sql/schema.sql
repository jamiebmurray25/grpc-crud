CREATE TABLE todos (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);
