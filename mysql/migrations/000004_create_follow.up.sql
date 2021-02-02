CREATE TABLE IF NOT EXISTS `follows` (
  `id` int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `follow_user_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (follow_user_id) REFERENCES users(id)
)
