CREATE TABLE `orders` (
  `id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
  `userid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `shoppingid` bigint(20) NOT NULL DEFAULT '0' COMMENT '收货信息表id',
  `payment` decimal(20,2) DEFAULT '0.00' COMMENT '实际付款金额,单位是元,保留两位小数',
  `paymenttype` tinyint(4) NOT NULL DEFAULT '1' COMMENT '支付类型,1-在线支付',
  `postage` int(10) NOT NULL DEFAULT '0' COMMENT '运费,单位是元',
  `status` smallint(6) NOT NULL DEFAULT '10' COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';