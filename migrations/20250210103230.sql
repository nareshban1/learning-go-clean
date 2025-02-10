-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `role`, ADD COLUMN `role_id` bigint unsigned NULL, ADD INDEX `fk_users_role` (`role_id`), ADD CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE ON DELETE SET NULL;
