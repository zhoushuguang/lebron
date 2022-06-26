CREATE TABLE `user_receive_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人名称',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否为默认地址',
  `post_code` varchar(100) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份/直辖市',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
  `region` varchar(100) NOT NULL DEFAULT '' COMMENT '区',
  `detail_address` varchar(128) NOT NULL DEFAULT '' COMMENT '详细地址(街道)',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户收货地址表';