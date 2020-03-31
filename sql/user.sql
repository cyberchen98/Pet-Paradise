CREATE TABLE IF NOT EXISTS  `user` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
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