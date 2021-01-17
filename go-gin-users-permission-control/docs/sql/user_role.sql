/*
Navicat MySQL Data Transfer

Source Database       : user_role
Target Server Type    : MYSQL
Date: 2021-01-15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Mapping table structure for user and role table
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
    `id`           int(10)         unsigned    NOT NULL AUTO_INCREMENT,
    `user_id`      int(10)         unsigned    DEFAULT '0' COMMENT '用户id',
    `role_id`      int(10)         unsigned    DEFAULT '0' COMMENT '角色id',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user与role的映射表';



INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (1, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (1, 2);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (2, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (3, 2);


