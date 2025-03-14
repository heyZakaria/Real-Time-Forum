-- Création de la table users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL CHECK(LENGTH(username) < 20),
    age TEXT NOT NULL ,
    gender TEXT CHECK(gender IN ('M', 'F')),
    first_name TEXT NOT NULL ,
    last_name TEXT NOT NULL ,
    email TEXT UNIQUE NOT NULL CHECK(LENGTH(email) < 70),
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Création de la table posts
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    creator TEXT,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Création de la table comments
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Création de la table categories
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name_category TEXT NOT NULL
);

-- Création de la table post_categories (table de liaison)
CREATE TABLE IF NOT EXISTS post_categories (

    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(category_id) REFERENCES categories(id),
    PRIMARY KEY (post_id, category_id)
);

-- Création de la table likes_dislikes
CREATE TABLE IF NOT EXISTS likes_dislikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    user_id INTEGER NOT NULL,
    thetype TEXT CHECK(thetype IN ('LIKE', 'DISLIKE')),
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
-- Creation de la table session
CREATE TABLE IF NOT EXISTS session (
id INTEGER PRIMARY KEY AUTOINCREMENT,
id_users INTEGER,
code TEXT NOT NULL,
FOREIGN KEY(id_users) REFERENCES users(id)
);


CREATE TABLE IF NOT EXISTS react_comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment_id INTEGER,
    user_id INTEGER NOT NULL,
    thetype TEXT CHECK(thetype IN ('LIKE', 'DISLIKE')),
    FOREIGN KEY(comment_id) REFERENCES comments(id)
    FOREIGN KEY(user_id) REFERENCES users(id)
);
