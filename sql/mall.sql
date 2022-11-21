/*========================>database user <===================================*/
CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `question` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码问题',
  `answer` varchar(100) NOT NULL DEFAULT '' COMMENT '找回密码答案',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_phone` (`phone`),
  UNIQUE KEY `uniq_username` (`username`),
  KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

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


/*========================>database product <===================================*/
CREATE DATABASE product;
USE product;

CREATE TABLE `product` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `cateid` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别Id',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
    `subtitle` varchar(200) NOT NULL DEFAULT '' COMMENT '商品副标题',
    `images` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址,逗号分隔',
    `detail` varchar(1024) NOT NULL DEFAULT '' COMMENT '商品详情',
    `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '价格,单位-元保留两位小数',
    `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
    `status` int(6) NOT NULL DEFAULT 1 COMMENT '商品状态.1-在售 2-下架 3-删除',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_cateid` (`cateid`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';


CREATE TABLE `category` (
    `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `parentid` smallint(6) NOT NULL DEFAULT 0 COMMENT '父类别id当id=0时说明是根节点,一级类别',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '类别状态1-正常,2-已废弃',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';

CREATE TABLE `product_operation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品id',
  `status` int NOT NULL DEFAULT '1' COMMENT '运营商品状态 0-下线 1-上线',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='商品运营表';

/*========================>database cart <===================================*/
CREATE DATABASE cart;
USE cart;

CREATE TABLE `cart` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '购物车id',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
    `quantity` int(11) NOT NULL DEFAULT 0 COMMENT '数量',
    `checked` int(11) NOT NULL DEFAULT 0 COMMENT '是否选择,1=已勾选,0=未勾选',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_userid` (`userid`),
    KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';


/*========================>database orders <===================================*/
CREATE DATABASE orders;
USE orders;

CREATE TABLE `orders` (
    `id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `shoppingid` bigint(20) NOT NULL DEFAULT 0 COMMENT '收货信息表id',
    `payment` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT '实际付款金额,单位是元,保留两位小数',
    `paymenttype` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付类型,1-在线支付',
    `postage` int(10)  NOT NULL DEFAULT 0 COMMENT '运费,单位是元',
    `status` smallint(6) NOT NULL DEFAULT 10 COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

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

CREATE TABLE `shipping` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '收货信息表id',
    `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
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


/*========================>database pay <===================================*/
CREATE DATABASE pay;
USE pay;

CREATE TABLE `payinfo` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付信息表id',
    `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `payplatform` tinyint(4) NOT NULL DEFAULT 0 COMMENT '支付平台:1-支付宝,2-微信',
    `platformnumber` varchar(200) NOT NULL DEFAULT '' COMMENT '支付流水号',
    `platformstatus` varchar(20) NOT NULL DEFAULT '' COMMENT '支付状态',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_orderid` (`orderid`),
    KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付信息表';


/*========================>database reply <===================================*/
CREATE DATABASE reply;
USE reply;

CREATE TABLE `reply`(
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论表id',
    `business` varchar(64) NOT NULL DEFAULT '' COMMENT '评论业务类型',
    `targetid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论目标id',
    `reply_userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复用户id',
    `be_reply_userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '被回复用户id',
    `parentid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论id',
    `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
    `image` varchar(255) NOT NULL DEFAULT '' COMMENT '评论图片',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_targetid` (`targetid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论列表';