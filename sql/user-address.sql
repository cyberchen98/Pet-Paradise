CREATE TABLE IF NOT EXISTS  `user_address` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `aid` bigint(11) unsigned NOT NULL COMMENT '用户地址id',
  `uid` bigint(11) unsigned NOT NULL COMMENT '用户id',
  `province` varchar(20) NOT NULL DEFAULT '' COMMENT '省份',
  `city` varchar(20) NOT NULL DEFAULT '' COMMENT '城市',
  `details` varchar(20) NOT NULL DEFAULT '' COMMENT '详细地址',
  `receiver` varchar(20) NOT NULL DEFAULT '' COMMENT '收件人',
  `phone_number` varchar(11) NOT NULL DEFAULT '' COMMENT '收件人电话',
  `post_code` varchar(6) NOT NULL DEFAULT '' COMMENT '邮编',
  `is_deletet` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datatime NOT NULL COMMENT '最近登陆时间',
  PRIMARY KEY (`id`),
  INDEX (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户地址表';