CREATE TABLE IF NOT EXISTS `posts` (
  `id` int unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `image_url` varchar(128) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
)
