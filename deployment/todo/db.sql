-- 创建数据库
CREATE DATABASE IF NOT EXISTS `newb` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `newb`;

-- 创建 todo 表
CREATE TABLE `todos` (
  `id` CHAR(36) NOT NULL PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `completed` BOOLEAN NOT NULL DEFAULT FALSE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建 task 表
CREATE TABLE `tasks` (
  `id` CHAR(36) NOT NULL PRIMARY KEY,
  `todo_id` CHAR(36) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `completed` BOOLEAN NOT NULL DEFAULT FALSE,
  CONSTRAINT `fk_tasks_todo` FOREIGN KEY (`todo_id`) REFERENCES `todos`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
