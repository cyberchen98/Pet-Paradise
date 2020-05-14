CREATE TABLE IF NOT EXISTS  `order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL COMMENT '用户id',
  `pid` int(11) unsigned NOT NULL COMMENT '产品id',
  `aid` int(11) unsigned NOT NULL COMMENT '地址id',
  `status` varchar(10) NOT NULL DEFAULT '' COMMENT '订单状态',
  `count_to_buy` int(10) NOT NULL DEFAULT '0' COMMENT '数量',
  `details` text NOT NULL COMMENT '订单详细信息',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单信息表';
