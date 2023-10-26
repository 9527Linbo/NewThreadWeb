/*
 Navicat Premium Data Transfer

 Source Server         : 官网项目
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 8.130.105.144:3306
 Source Schema         : NewThread

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 26/10/2023 09:04:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_article
-- ----------------------------
DROP TABLE IF EXISTS `t_article`;
CREATE TABLE `t_article`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `type` int(11) NOT NULL COMMENT '文章类型（1、知识分享，2、新闻。3、文化活动）',
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文章正文',
  `user_id` bigint(20) NOT NULL COMMENT '作者id（user表中的id）',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1001 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_awards
-- ----------------------------
DROP TABLE IF EXISTS `t_awards`;
CREATE TABLE `t_awards`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `type_id` int(11) NOT NULL COMMENT '类型表id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '获奖名次',
  `student_id` bigint(20) NOT NULL COMMENT '负责人id（student表中的id）',
  `teacher_id` bigint(20) NULL DEFAULT NULL COMMENT '指导老师id（teacher表中的id）',
  `TeamName` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '获奖团队',
  `number` bigint(20) NOT NULL COMMENT '获奖人数',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_awardsName
-- ----------------------------
DROP TABLE IF EXISTS `t_awardsName`;
CREATE TABLE `t_awardsName`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类型名',
  `description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_group
-- ----------------------------
DROP TABLE IF EXISTS `t_group`;
CREATE TABLE `t_group`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组名',
  `description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `teacher_id` bigint(20) NOT NULL COMMENT '指导老师id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_image
-- ----------------------------
DROP TABLE IF EXISTS `t_image`;
CREATE TABLE `t_image`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `url` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'URL',
  `article_id` bigint(20) NOT NULL COMMENT '文章id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_position
-- ----------------------------
DROP TABLE IF EXISTS `t_position`;
CREATE TABLE `t_position`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '职位名称',
  `description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '职位描述',
  `user_id` bigint(20) NOT NULL COMMENT '所属人id（user表中的id）',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_student
-- ----------------------------
DROP TABLE IF EXISTS `t_student`;
CREATE TABLE `t_student`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
  `group_id` bigint(20) NOT NULL COMMENT 'group表中的id',
  `user_id` bigint(20) NOT NULL COMMENT 'user表中的id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 101 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_studentGraduate
-- ----------------------------
DROP TABLE IF EXISTS `t_studentGraduate`;
CREATE TABLE `t_studentGraduate`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
  `description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT 'user表中的id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_teacher
-- ----------------------------
DROP TABLE IF EXISTS `t_teacher`;
CREATE TABLE `t_teacher`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名',
  `user_id` bigint(20) NOT NULL COMMENT 'user表中的id',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '唯一标识id',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `account` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `update_Time` datetime NOT NULL COMMENT '修改时间',
  `create_Time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 301 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
