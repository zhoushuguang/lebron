CREATE TABLE `user_collection` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '收藏Id',
    `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
    `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
    PRIMARY KEY (`id`),
    UNIQUE KEY `UN_collection_uid_product_id`(uid,product_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏表';