#员工信息表
CREATE TABLE `t_employee_inf` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT '员工姓名',
  `gender` char(4) NOT NULL COMMENT '性别',
  `id_card` char(18) NOT NULL COMMENT '身份证号',
  `birthday` varchar(20) DEFAULT NULL COMMENT '出生日期',
  `phone` varchar(11) DEFAULT NULL COMMENT '电话号码',
  `address` varchar(64) DEFAULT NULL COMMENT '联系地址',
  `is_head` varchar(4) DEFAULT NULL COMMENT '是否部长',
  `department` enum('机修部','美容部','保养部','运营部') DEFAULT NULL COMMENT '部门',
  `salary` int(11) DEFAULT NULL COMMENT '薪水',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


#绩效表
CREATE TABLE `t_employee_performance` (
  `id` int(11) NOT NULL,
  `name` varchar(10) NOT NULL COMMENT '员工姓名',
  `service_item` varchar(45) DEFAULT NULL,
  `performance` int(11) DEFAULT '0' COMMENT '奖金',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '服务项目',
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


#客户表
CREATE TABLE `t_guest` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(10) NOT NULL COMMENT '用户名字',
  `user_phone` varchar(11) NOT NULL COMMENT '用户电话',
  `car_license` varchar(10) NOT NULL COMMENT '车牌号',
  `car_brand` varchar(10) DEFAULT NULL COMMENT '品牌',
  `service_nums` int(11) DEFAULT '0' COMMENT '服务次数',
  `consumption` int(11) DEFAULT '0' COMMENT '总消费',
  `is_vip` varchar(4) NOT NULL DEFAULT '否',
  `created_at` datetime(3) DEFAULT NULL COMMENT '否',
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `car_license_UNIQUE` (`car_license`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


#订单表
CREATE TABLE `t_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(10) NOT NULL COMMENT '用户名字',
  `user_phone` varchar(11) NOT NULL COMMENT '用户电话',
  `car_license` varchar(10) NOT NULL COMMENT '车牌号',
  `car_brand` varchar(10) DEFAULT NULL COMMENT '品牌',
  `service_item` varchar(30) NOT NULL COMMENT '服务项目',
  `material` varchar(45) DEFAULT NULL COMMENT '配件',
  `cost` int(11) NOT NULL DEFAULT '0' COMMENT '金额',
  `employee` varchar(10) NOT NULL COMMENT '施工人',
  `operator` varchar(10) DEFAULT NULL COMMENT '操作人',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


#成本表
CREATE TABLE `t_service_cost` (
  `id` int(11) NOT NULL COMMENT '订单号',
  `car_license` varchar(20) NOT NULL COMMENT '车牌',
  `service_item` varchar(30) NOT NULL COMMENT '服务项目',
  `material` varchar(20) NOT NULL COMMENT '配件',
  `material_cost` int(11) DEFAULT NULL COMMENT '成本',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


#进货表
CREATE TABLE `t_store` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `material_name` varchar(20) NOT NULL COMMENT '配件名字',
  `unit_price` int(11) NOT NULL COMMENT '单价',
  `number` int(11) NOT NULL COMMENT '数量',
  `total_price` int(11) DEFAULT NULL COMMENT '总价',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


#user表
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_name` longtext,
  `password_digest` longtext,
  `nickname` longtext,
  `status` longtext,
  `avatar` varchar(1000) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
