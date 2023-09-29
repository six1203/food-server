CREATE TABLE `collection_shop` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `category` varchar(30) NOT NULL DEFAULT '' COMMENT '美食分类',
   `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
   `logo` varchar(2000) NOT NULL DEFAULT '' COMMENT 'logo',
   `address` varchar(1000) NOT NULL DEFAULT '' COMMENT '地址',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `deleted_at` datetime DEFAULT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `ux_category_name` (`category`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '火锅', '九鼎轩脆毛肚火锅(虹梅路店)', 'https://qcloud.dpfile.com/pc/06XQadZ452oUJ-bn7sE_u0T58aofcdajrBhhlusz9aKE1dUpvbW1gcNw26qKBJshl0cm-Lf9tDMlLZpO7rb3bg.jpg', '虹梅路3279号（华光花园东门右侧,10号线龙溪路2号出口右拐直行）', '2023-09-30 01:20:18', '2023-09-30 01:20:45', NULL);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '烤肉', '西塔老太太泥炉烤肉(长宁大融城黑金店)', 'https://img.meituan.net/msmerchant/84221a51a699dac298bc104f5bc49707505475.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '淞虹路377号长宁大融城LG2层98室', '2023-09-30 01:21:09', '2023-09-30 01:21:09', NULL);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, '川菜', '天辣(日月光店) ', 'https://img.meituan.net/msmerchant/7d55bd6fd59513f17468cbac9d606729590465.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '徐家汇路618号日月光广场4楼徐家汇区4F—XJH14', '2023-09-30 01:21:45', '2023-09-30 01:21:45', NULL);
