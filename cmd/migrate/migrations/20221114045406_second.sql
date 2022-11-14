-- modify "users" table
ALTER TABLE `users` DROP CONSTRAINT `users_chk_1`, ADD CHECK (age < 10);
