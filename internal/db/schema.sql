CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE,
    password_hash TEXT,
    email TEXT,
    plan_name TEXT, -- ex: "Starter", "Pro", "Unlimited"
    disk_limit_mb INTEGER,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS domains (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    domain_name TEXT UNIQUE,
    php_version TEXT DEFAULT '8.3',
    ssl_enabled BOOLEAN DEFAULT 0,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
