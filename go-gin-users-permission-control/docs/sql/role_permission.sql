/*
Navicat MySQL Data Transfer

Source Database       : role_permission
Target Server Type    : MYSQL
Date: 2021-01-15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Mapping table structure for role and permission table
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
     `id`           int(10)         unsigned    NOT NULL AUTO_INCREMENT,
     `role_id`      int(10)         unsigned    DEFAULT '0' COMMENT '角色id',
     `permission_id`int(10)         unsigned    DEFAULT '0' COMMENT '权限id',

     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='role与permission的映射表';



INSERT INTO `role_permission` (`role_id`, `permission_id`) VALUES (1, 1);
INSERT INTO `role_permission` (`role_id`, `permission_id`) VALUES (1, 2);
INSERT INTO `role_permission` (`role_id`, `permission_id`) VALUES (1, 3);
INSERT INTO `role_permission` (`role_id`, `permission_id`) VALUES (2, 1);


