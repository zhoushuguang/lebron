CREATE TABLE `shipping` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '收货信息表id',
  `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
  `userid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `receiver_name` varchar(20) NOT NULL DEFAULT '' COMMENT '收货姓名',
  `receiver_phone` varchar(20) NOT NULL DEFAULT '' COMMENT '收货固定电话',
  `receiver_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '收货移动电话',
  `receiver_province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
  `receiver_city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
  `receiver_district` varchar(20) NOT NULL DEFAULT '' COMMENT '区/县',
  `receiver_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `ix_orderid` (`orderid`),
  KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收货信息表';