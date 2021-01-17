/*
Navicat MySQL Data Transfer

Source Database       : user
Target Server Type    : MYSQL
Date: 2021-01-15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user table
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id`            int(10)         unsigned    NOT NULL AUTO_INCREMENT,
    `username`      varchar(100)                DEFAULT ''  COMMENT '用户名',
    `password`      varchar(100)                DEFAULT ''  COMMENT '密码',
    `email`         varchar(50)                 DEFAULT ''  COMMENT '邮箱',
    `phone`         varchar(20)                 DEFAULT ''  COMMENT '电话号码',
    `sex`           tinyint(3)                  DEFAULT '0' COMMENT '用户性别（1男 2女）',
    `remark`        varchar(255)                DEFAULT ''  COMMENT '备注',
    `created_on`    int(10)         unsigned    DEFAULT '0' COMMENT '新建时间',
    `created_by`    varchar(100)                DEFAULT ''  COMMENT '创建人',
    `modified_on`   int(10)         unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by`   varchar(255)                DEFAULT ''  COMMENT '修改人',
    `deleted_on`    int(10)         unsigned    DEFAULT '0' COMMENT '删除时间',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户列表';



INSERT INTO `user` (`username`, `password`) VALUES ('bingcu', '123456');
INSERT INTO `user` (`username`, `password`) VALUES ('bingcu2', '123456');
INSERT INTO `user` (`username`, `password`) VALUES ('bingcu3', '123456');


