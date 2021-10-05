CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL ,
                        `score` bigint(20) NOT NULL ,
                        `chances` tinyint(8) NOT NULL DEFAULT '0',
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                        `gift` varchar(64) COLLATE utf8mb4_general_ci,
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
                                    CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;