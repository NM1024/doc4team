 

/*Table structure for table `apidoc` */

DROP TABLE IF EXISTS `apidoc`;

CREATE TABLE `apidoc` (
  `api_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 DEFAULT NULL COMMENT '名称',
  `address` varchar(100) CHARACTER SET utf8 DEFAULT NULL COMMENT 'API地址',
  `method` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '类型',
  `header` varchar(1000) CHARACTER SET utf8 DEFAULT NULL COMMENT '请求头',
  `parameters` varchar(1000) CHARACTER SET utf8 DEFAULT NULL,
  `body` varchar(1000) CHARACTER SET utf8 DEFAULT NULL,
  `response` varchar(1000) CHARACTER SET utf8 DEFAULT NULL,
  `describe` varchar(300) CHARACTER SET utf8 DEFAULT NULL COMMENT '描述',
  `version` varchar(20) CHARACTER SET utf8 DEFAULT NULL COMMENT '版本',
  `remark` varchar(300) CHARACTER SET utf8 DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`api_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;

/*Table structure for table `doctag` */

DROP TABLE IF EXISTS `doctag`;

CREATE TABLE `doctag` (
  `tag_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Table structure for table `doctagmap` */

DROP TABLE IF EXISTS `doctagmap`;

CREATE TABLE `doctagmap` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `api_id` bigint(20) unsigned NOT NULL,
  `tag_id` bigint(20) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

 