/*
Navicat MySQL Data Transfer

Source Database       : user_group
Target Server Type    : MYSQL
Date: 2021-01-15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for role table
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
    `id`            int(10)         unsigned    NOT NULL AUTO_INCREMENT,
    `name`          varchar(20)                 DEFAULT ''  COMMENT '角色名称',
    `description`   varchar(100)                DEFAULT ''  COMMENT '描述',
    `remark`        varchar(255)                DEFAULT ''  COMMENT '备注',
    `created_on`    int(10)         unsigned    DEFAULT '0' COMMENT '新建时间',
    `created_by`    varchar(100)                DEFAULT ''  COMMENT '创建人',
    `modified_on`   int(10)         unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by`   varchar(255)                DEFAULT ''  COMMENT '修改人',
    `deleted_on`    int(10)         unsigned    DEFAULT '0' COMMENT '删除时间',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户组列表';



INSERT INTO `role` (`name`, `description`) VALUES ('BinGCU_Role', '描述：天天向上');
INSERT INTO `role` (`name`, `description`) VALUES ('bingcu_role', '描述：天天向下');


