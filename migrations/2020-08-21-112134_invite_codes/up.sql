CREATE TABLE `invite_codes` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    `code` varchar(255) NOT NULL,
    `times_used` INTEGER NOT NULL,
    `max_usages` INTEGER NOT NULL,
    `generated` INTEGER NOT NULL REFERENCES users(id)
);