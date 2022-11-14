-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `age` bigint NOT NULL, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`), CHECK (age < 10)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
