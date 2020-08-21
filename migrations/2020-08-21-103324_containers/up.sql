CREATE TABLE `containers` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `owner` INTEGER NOT NULL REFERENCES groups(id),
    `name` varchar(255) NOT NULL,
    `path` varchar(255) NOT NULL,
    `subdomain` varchar(255)
);