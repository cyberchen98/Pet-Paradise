CREATE TABLE IF NOT EXISTS  `user` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(11) unsigned NOT NULL COMMENT '用户id',
  `user_name` varchar(20) NOT NULL COMMENT '用户名',
  `user_password` varchar(20) NOT NULL COMMENT '用户密码',
  `user_email` varchar(20) COMMENT '用户邮箱',
  `user_phone` varchar(20) COMMENT '用户电话',
  `role` enum('admin', 'vip', 'common') NOT NULL COMMENT '用户身份',
  `is_deletet` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datatime NOT NULL COMMENT '最近登陆时间',
  PRIMARY KEY (`uid`),
  INDEX (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';