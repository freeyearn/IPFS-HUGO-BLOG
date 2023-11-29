SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
	`auto_id` bigint NOT NULL AUTO_INCREMENT,
	`article_id` varchar(22),
	`path` varchar(255),
	`author_id` varchar(22),
	`title` varchar(255),
	`description` text,
	`tags` text,
	`category` text,
	`keyword` text,
	`next` varchar(255),
	`prev` varchar(255),
	`status` tinyint,
	`is_deleted` tinyint(1),
	`publish_time` datetime NOT NULL DEFAULT current_timestamp(),
	`update_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`auto_id`),
	KEY `idx_article_id` (`article_id`),
	KEY `idx_author_id` (`author_id`)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COLLATE utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
	`auto_id` bigint NOT NULL AUTO_INCREMENT,
	`category_id` varchar(22),
	`user_id` varchar(22),
	`name` varchar(255),
	`is_deleted` tinyint(1),
	`create_time` datetime NOT NULL DEFAULT current_timestamp(),
	`category_name` varchar(255),
	PRIMARY KEY (`auto_id`),
	KEY `idx_category_id` (`category_id`)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COLLATE utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
	`auto_id` bigint NOT NULL AUTO_INCREMENT,
	`user_id` varchar(22),
	`wallet` varchar(255),
	`name` varchar(255),
	`account` varchar(500),
	`password` varchar(500),
	`profile` varchar(500),
	`status` tinyint,
	`is_deleted` tinyint(1),
	`create_time` datetime NOT NULL DEFAULT current_timestamp(),
	`update_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`auto_id`)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COLLATE utf8mb4_0900_ai_ci;
