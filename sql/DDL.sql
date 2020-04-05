CREATE DATABASE IF NOT EXISTS `pet-paradise` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE pet-paradise;


CREATE TABLE IF NOT EXISTS  `order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL COMMENT '用户id',
  `pid` int(11) unsigned NOT NULL COMMENT '产品id',
  `aid` int(11) unsigned NOT NULL COMMENT '地址id',
  `status` varchar(10) NOT NULL DEFAULT '' COMMENT '订单状态',
  `details` text NOT NULL COMMENT '订单详细信息',
  `is_closed` int(1) NOT NULL DEFAULT '0' COMMENT '订单是否关闭：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单信息表';


CREATE TABLE IF NOT EXISTS  `product` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
  `parent_product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '所属产品类',
  `price` varchar(10) NOT NULL DEFAULT '' COMMENT '产品价格',
  `describe` varchar(100) NOT NULL DEFAULT '' COMMENT '产品描述',
  `details` text NOT NULL COMMENT '产品详细信息',
  `count_remain` int(10) NOT NULL DEFAULT '0' COMMENT '产品剩余数量',
  `is_on_sale` int(1) NOT NULL DEFAULT '1' COMMENT '产品是否在售：0否、1是',
  `is_on_discount` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否在折扣：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='产品信息表';


CREATE TABLE IF NOT EXISTS  `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) NOT NULL COMMENT '用户名',
  `user_password` varchar(20) NOT NULL COMMENT '用户密码',
  `user_email` varchar(20) COMMENT '用户邮箱',
  `user_phone` varchar(20) COMMENT '用户电话',
  `role` enum('admin', 'vip', 'common') NOT NULL DEFAULT 'common' COMMENT '用户身份',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近登陆时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE IF NOT EXISTS  `user_address` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL COMMENT '用户id',
  `province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
  `city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
  `details` varchar(20) NOT NULL COMMENT '详细地址',
  `receiver` varchar(20) NOT NULL DEFAULT '' COMMENT '收件人',
  `phone_number` varchar(11) NOT NULL DEFAULT '' COMMENT '收件人电话',
  `post_code` varchar(6) NOT NULL DEFAULT '' COMMENT '邮编',
  `is_deletet` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近登陆时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户地址表';

ALTER TABLE `order` ADD INDEX uid_index(uid);
ALTER TABLE `order` ADD INDEX uid_index(pid);
ALTER TABLE `product` ADD INDEX uid_index(uid);
ALTER TABLE `product` ADD INDEX product_index(parent_prodcut_name, product_name);
ALTER TABLE `user` ADD INDEX user_name_index(user_name);
ALTER TABLE `user` ADD INDEX user_phone_index(user_phone);
ALTER TABLE `user` ADD INDEX user_email_index(user_email);
ALTER TABLE `user_address` ADD INDEX uid_index(uid);