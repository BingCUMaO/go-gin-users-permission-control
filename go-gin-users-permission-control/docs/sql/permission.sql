/*
Navicat MySQL Data Transfer

Source Database       : permission
Target Server Type    : MYSQL
Date: 2021-01-15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for permission table
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
      `id`            int(10)         unsigned    NOT NULL AUTO_INCREMENT,
      `permission_key`varchar(20)                 DEFAULT ''  COMMENT '权限字符',
      `description`   varchar(100)                DEFAULT ''  COMMENT '描述',
      `remark`        varchar(255)                DEFAULT ''  COMMENT '备注',
      `created_on`    int(10)         unsigned    DEFAULT '0' COMMENT '新建时间',
      `created_by`    varchar(100)                DEFAULT ''  COMMENT '创建人',
      `modified_on`   int(10)         unsigned    DEFAULT '0' COMMENT '修改时间',
      `modified_by`   varchar(255)                DEFAULT ''  COMMENT '修改人',
      `deleted_on`    int(10)         unsigned    DEFAULT '0' COMMENT '删除时间',

      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户权限管理';



INSERT INTO `permission` (`permission_key`, `description`) VALUES ('@@@read', '读权限字符');
INSERT INTO `permission` (`permission_key`, `description`) VALUES ('@@@modify', '修改权限字符');
INSERT INTO `permission` (`permission_key`, `description`) VALUES ('@@@delete', '删除权限字符');


