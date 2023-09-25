CREATE TABLE `user` (
    `id` bigint NOT NULL,
    `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名，唯一',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密文',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
