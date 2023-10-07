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


INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (1, '火锅', '九鼎轩脆毛肚火锅(虹梅路店)', 'https://qcloud.dpfile.com/pc/06XQadZ452oUJ-bn7sE_u0T58aofcdajrBhhlusz9aKE1dUpvbW1gcNw26qKBJshl0cm-Lf9tDMlLZpO7rb3bg.jpg', '虹梅路3279号（华光花园东门右侧,10号线龙溪路2号出口右拐直行）', '2023-09-30 01:20:18', '2023-10-06 00:18:15', '2023-10-06 00:18:15', 5);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (2, '烤肉', '西塔老太太泥炉烤肉(长宁大融城黑金店)', 'https://img.meituan.net/msmerchant/84221a51a699dac298bc104f5bc49707505475.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '淞虹路377号长宁大融城LG2层98室', '2023-09-30 01:21:09', '2023-10-06 10:33:50', NULL, 5);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (3, '川菜', '天辣(日月光店) ', 'https://img.meituan.net/msmerchant/7d55bd6fd59513f17468cbac9d606729590465.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '徐家汇路618号日月光广场4楼徐家汇区4F—XJH14', '2023-09-30 01:21:45', '2023-10-04 11:05:58', NULL, 2);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (17, '烧烤', '破店', 'https://img.meituan.net/msmerchant/c51adf941f3312bd6fa224c980515f9f828470.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '北京西路1700号-6（静安中华大厦正对面）', '2023-10-02 12:09:26', '2023-10-04 11:05:58', NULL, 5);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (19, '海鲜自助', '万岛', 'http://p0.meituan.net/mogu/3371d407b69266088094ec56589d186327923.jpg%40280w_212h_1e_1c_1l%7Cwatermark%3D0', '大渡河路1428-1438号信泰中心(商场)一层T102A-T102D单元', '2023-10-03 10:59:27', '2023-10-04 21:19:01', NULL, 5);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (38, '222', '3333333', '3333333', '3333333', '2023-10-04 20:09:40', '2023-10-06 00:18:20', '2023-10-06 00:18:20', 5);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (39, '火锅', '左庭右院鲜牛肉火锅(南翔印象城MEGA店) ', 'https://img.meituan.net/content/4875486db7aeec58a8e9ff309b3e08e7512434.png%40340w_255h_1e_1c_1l%7Cwatermark%3D0', '陈翔公路2299号南翔印象城4楼', '2023-10-06 00:28:44', '2023-10-06 00:28:44', NULL, 3);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (40, '火锅', '八合里牛肉火锅(中山公园龙之梦店) ', 'https://img.meituan.net/content/40f4886458828ad00c4a658b4f2228ac59122.jpg%40340w_255h_1e_1c_1l%7Cwatermark%3D0', '长宁路1018号龙之梦购物公园8楼8059-8063铺面', '2023-10-06 00:32:10', '2023-10-06 00:32:10', NULL, 2);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (41, '串串', '付小姐在成都(合生汇店) ', 'http://p0.meituan.net/biztone/150307802_1691376381748.jpeg%40340w_255h_1e_1c_1l%7Cwatermark%3D0', '翔殷路1099号合生汇商场L5-11', '2023-10-06 00:33:47', '2023-10-06 00:33:47', NULL, 4);
INSERT INTO `food`.`collection_shop` (`id`, `category`, `name`, `logo`, `address`, `created_at`, `updated_at`, `deleted_at`, `star`) VALUES (42, '川菜', '成都鬼瘾食(美罗城店) ', 'http://p1.meituan.net/biztone/1662238277_1690275547818.jpeg%40340w_255h_1e_1c_1l%7Cwatermark%3D0', '肇嘉浜路1111号美罗城5楼', '2023-10-06 00:34:47', '2023-10-06 00:35:05', NULL, 3);
