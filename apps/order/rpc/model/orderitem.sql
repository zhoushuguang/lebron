CREATE TABLE `orderitem` (
     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单子表id',
     `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
     `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
     `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
     `product_name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
     `product_image` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片地址',
     `current_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '生成订单时的商品单价，单位是元,保留两位小数',
     `quantity` int(10) NOT NULL DEFAULT 0 COMMENT '商品数量',
     `total_price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品总价,单位是元,保留两位小数',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `ix_orderid` (`order_id`),
     KEY `ix_userid` (`user_id`),
     KEY `ix_proid` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';