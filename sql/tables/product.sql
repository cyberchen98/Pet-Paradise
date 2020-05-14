CREATE TABLE IF NOT EXISTS  `product` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
  `parent_product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '所属产品类',
  `price` varchar(10) NOT NULL DEFAULT '' COMMENT '产品价格',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '产品描述',
  `details` text NOT NULL DEFAULT '' COMMENT '产品详细信息',
  `count_remain` int(10) NOT NULL DEFAULT '0' COMMENT '产品剩余数量',
  `is_on_sale` int(1) NOT NULL DEFAULT '1' COMMENT '产品是否在售：0否、1是',
  `is_on_discount` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否在折扣：0否、1是',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '产品是否删除：0否、1是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='产品信息表';