CREATE TABLE IF NOT EXISTS  `order` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `oid` bigint(11) unsigned NOT NULL COMMENT '订单id',
  `uid` bigint(11) unsigned NOT NULL COMMENT '用户id',
  `pid` bigint(11) unsigned NOT NULL COMMENT '产品id',
  `status` varchar(10) NOT NULL DEFAULT '' COMMENT '订单状态',
  `address` text NOT NULL DEFAULT '' COMMENT '收货地址',
  `details` text NOT NULL DEFAULT '' COMMENT '订单详细信息',
  `is_closed` int(1) NOT NULL DEFAULT '0' COMMENT '订单是否关闭：0否、1是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datatime NOT NULL COMMENT '最近更新时间',
  PRIMARY KEY (`uid`),
  INDEX (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单信息表';