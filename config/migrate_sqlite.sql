CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    account VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT "user"
);

CREATE TABLE items (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100)
);

CREATE TABLE baskets (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL 
);


INSERT INTO users (id, account, password, role) VALUES (1, "manager", "manager", "manager");
