CREATE TABLE `users` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `username` varchar(255) NOT NULL,
    `password_hash` varchar(255) NOT NULL,
    `password_salt` varchar(255) NOT NULL,
    `invite_code_used` INTEGER REFERENCES invite_codes(id),
    `role` INTEGER REFERENCES roles(id)
);