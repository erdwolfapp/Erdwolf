CREATE TABLE `users_groups` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `user` int NOT NULL REFERENCES users(id),
    `group` int NOT NULL REFERENCES groups(id)
);