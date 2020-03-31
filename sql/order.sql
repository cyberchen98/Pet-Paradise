CREATE TABLE IF NOT EXISTS  `order` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(11) unsigned NOT NULL COMMENT '用户id',
  `pid` bigint(11) unsigned NOT NULL COMMENT '产品id',
  `aid` bigint(11) unsigned NOT NULL COMMENT '地址id',
  `status` varchar(10) NOT NULL DEFAULT '' COMMENT '订单状态',
  `details` text NOT NULL COMMENT '订单详细信息',
  `is_closed` int(1) NOT NULL DEFAULT '0' COMMENT '订单是否关闭：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单信息表';