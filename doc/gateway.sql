

CREATE TABLE `gateway_api` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `service_id` varchar(20) NOT NULL COMMENT '服务id',
  `api_method` varchar(10) NOT NULL COMMENT 'http方法',
  `api_url` varchar(50) NOT NULL COMMENT 'api url',
  `status` varchar(2) NOT NULL DEFAULT '0' COMMENT '状态 00-正常 01-失效',
  `create_time` varchar(20) NOT NULL COMMENT '创建时间',
  `update_time` varchar(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `api_url` (`api_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
