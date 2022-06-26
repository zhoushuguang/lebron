CREATE TABLE `orderitem` (
     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单子表id',
     `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
     `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
     `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
     `proname` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
     `proimage` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片地址',
     `currentunitprice` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '生成订单时的商品单价，单位是元,保留两位小数',
     `quantity` int(10) NOT NULL DEFAULT 0 COMMENT '商品数量',
     `totalprice` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '商品总价,单位是元,保留两位小数',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `ix_orderid` (`orderid`),
     KEY `ix_userid` (`userid`),
     KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';