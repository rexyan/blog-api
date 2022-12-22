/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 22/12/2022 17:05:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for about
-- ----------------------------
DROP TABLE IF EXISTS `about`;
CREATE TABLE `about`  (
                          `id` bigint(20) NOT NULL,
                          `name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `name_zh` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of about
-- ----------------------------
INSERT INTO `about` VALUES (1, 'title', '标题', '111111111111');
INSERT INTO `about` VALUES (2, 'musicId', '网易云歌曲ID', '423015580');
INSERT INTO `about` VALUES (3, 'content', '正文Markdown', '1111111111111');
INSERT INTO `about` VALUES (4, 'commentEnabled', '评论开关', 'false');

-- ----------------------------
-- Table structure for blog
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
                         `first_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章首图，用于随机文章展示',
                         `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章正文',
                         `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
                         `is_published` bit(1) NOT NULL COMMENT '公开或私密',
                         `is_recommend` bit(1) NOT NULL COMMENT '推荐开关',
                         `is_appreciation` bit(1) NOT NULL COMMENT '赞赏开关',
                         `is_comment_enabled` bit(1) NOT NULL COMMENT '评论开关',
                         `create_time` datetime(0) NOT NULL COMMENT '创建时间',
                         `update_time` datetime(0) NOT NULL COMMENT '更新时间',
                         `views` int(11) NOT NULL COMMENT '浏览次数',
                         `words` int(11) NOT NULL COMMENT '文章字数',
                         `read_time` int(11) NOT NULL COMMENT '阅读时长(分钟)',
                         `category_id` bigint(20) NOT NULL COMMENT '文章分类',
                         `is_top` bit(1) NOT NULL COMMENT '是否置顶',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码保护',
                         `user_id` bigint(20) NULL DEFAULT NULL COMMENT '文章作者',
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `type_id`(`category_id`) USING BTREE,
                         INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog
-- ----------------------------
INSERT INTO `blog` VALUES (1, 'xxxxxxxxxxxxxxxxxxxx', '12123', '1213', '12331', b'1', b'0', b'0', b'0', '2022-12-12 10:39:24', '2022-12-12 10:41:57', 0, 1231, 6, 1, b'0', '1231', 1);
INSERT INTO `blog` VALUES (2, '1111111111111111111', '1111111111111111', '1111111111111111', '1111111111111', b'1', b'0', b'0', b'0', '2022-12-12 10:47:21', '2022-12-12 10:47:21', 0, 131313, 657, 2, b'1', '', 1);
INSERT INTO `blog` VALUES (3, '自知则知之11222', '1111111111111111', '11111111111111111', '1111111111111', b'1', b'1', b'0', b'0', '2022-12-12 10:47:21', '2022-12-19 10:30:12', 0, 131313, 657, 1, b'1', '', 1);
INSERT INTO `blog` VALUES (4, '12313123', '12313', '12313', '123131', b'1', b'0', b'0', b'0', '2022-12-19 14:33:31', '2022-12-19 14:33:31', 0, 123, 1, 3, b'0', '', 1);

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
                             `blog_id` bigint(20) NOT NULL,
                             `tag_id` bigint(20) NOT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (0, 2);
INSERT INTO `blog_tag` VALUES (2, 1);
INSERT INTO `blog_tag` VALUES (0, 1);
INSERT INTO `blog_tag` VALUES (4, 1);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
                             `id` bigint(20) NOT NULL AUTO_INCREMENT,
                             `category_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (3, '嘻嘻信息');
INSERT INTO `category` VALUES (4, '12313');
INSERT INTO `category` VALUES (5, '2313');
INSERT INTO `category` VALUES (6, '玉兔');

-- ----------------------------
-- Table structure for city_visitor
-- ----------------------------
DROP TABLE IF EXISTS `city_visitor`;
CREATE TABLE `city_visitor`  (
                                 `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '城市名称',
                                 `uv` int(11) NOT NULL COMMENT '独立访客数量',
                                 PRIMARY KEY (`city`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
                            `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
                            `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像(图片路径)',
                            `create_time` datetime(0) NULL DEFAULT NULL COMMENT '评论时间',
                            `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '评论者ip地址',
                            `is_published` bit(1) NOT NULL COMMENT '公开或回收站',
                            `is_admin_comment` bit(1) NOT NULL COMMENT '博主回复',
                            `page` int(11) NOT NULL COMMENT '0普通文章，1关于我页面，2友链页面',
                            `is_notice` bit(1) NOT NULL COMMENT '接收邮件提醒',
                            `blog_id` bigint(20) NULL DEFAULT NULL COMMENT '所属的文章',
                            `parent_comment_id` bigint(20) NOT NULL COMMENT '父评论id，-1为根评论',
                            `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人网站',
                            `qq` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '如果评论昵称为QQ号，则将昵称和头像置为QQ昵称和QQ头像，并将此字段置为QQ号备份',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for exception_log
-- ----------------------------
DROP TABLE IF EXISTS `exception_log`;
CREATE TABLE `exception_log`  (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
                                  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
                                  `param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求参数',
                                  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作描述',
                                  `error` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '异常信息',
                                  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
                                  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
                                  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
                                  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
                                  `create_time` datetime(0) NOT NULL COMMENT '操作时间',
                                  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user-agent用户代理',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exception_log
-- ----------------------------
INSERT INTO `exception_log` VALUES (1, '1', '1', '1', '1', '1', '1', '1', '1', '1', '2022-12-22 14:57:49', '1');

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend`  (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                           `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
                           `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
                           `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '站点',
                           `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
                           `is_published` bit(1) NOT NULL COMMENT '公开或隐藏',
                           `views` int(11) NOT NULL COMMENT '点击次数',
                           `create_time` datetime(0) NOT NULL COMMENT '创建时间',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friend
-- ----------------------------
INSERT INTO `friend` VALUES (1, '55', '44', '33', '22', b'1', 55, '2022-12-20 07:56:59');
INSERT INTO `friend` VALUES (2, '阿呆', '而', '33', '44', b'0', 55, '2022-12-20 07:56:59');
INSERT INTO `friend` VALUES (3, '兔儿童', '345算法算法', '33', '44', b'0', 55, '2022-12-19 23:56:59');
INSERT INTO `friend` VALUES (4, '1', '1', '1', '1', b'1', 0, '2022-12-20 09:05:05');

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log`  (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
                              `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
                              `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
                              `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
                              `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
                              `status` bit(1) NULL DEFAULT NULL COMMENT '登录状态',
                              `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作描述',
                              `create_time` datetime(0) NOT NULL COMMENT '登录时间',
                              `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user-agent用户代理',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_log
-- ----------------------------
INSERT INTO `login_log` VALUES (1, '1', '1', '1', '1', '1', b'1', '1', '2022-12-22 14:57:30', '1');

-- ----------------------------
-- Table structure for moment
-- ----------------------------
DROP TABLE IF EXISTS `moment`;
CREATE TABLE `moment`  (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                           `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '动态内容',
                           `create_time` datetime(0) NOT NULL COMMENT '创建时间',
                           `likes` int(11) NULL DEFAULT NULL COMMENT '点赞数量',
                           `is_published` bit(1) NOT NULL COMMENT '是否公开',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of moment
-- ----------------------------
INSERT INTO `moment` VALUES (2, '211111111', '2022-12-11 16:00:00', 111, b'0');
INSERT INTO `moment` VALUES (3, '131313', '2022-12-28 16:00:00', 12313, b'1');
INSERT INTO `moment` VALUES (4, '仄仄仄仄仄仄仄仄仄做着做着', '2022-12-29 16:00:00', 0, b'1');
INSERT INTO `moment` VALUES (5, '345345请求嗯群', '2022-12-19 05:45:53', 111, b'0');
INSERT INTO `moment` VALUES (6, '让他用疼痛', '2022-12-19 05:46:14', 111, b'0');

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作者用户名',
                                  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
                                  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
                                  `param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求参数',
                                  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作描述',
                                  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
                                  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
                                  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
                                  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
                                  `times` int(11) NOT NULL COMMENT '请求耗时（毫秒）',
                                  `create_time` datetime(0) NOT NULL COMMENT '操作时间',
                                  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user-agent用户代理',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operation_log
-- ----------------------------
INSERT INTO `operation_log` VALUES (1, '1', '1', '1', '1', '1', '1', '1', '1', '1', 1, '2022-12-22 14:58:26', '1');

-- ----------------------------
-- Table structure for schedule_job
-- ----------------------------
DROP TABLE IF EXISTS `schedule_job`;
CREATE TABLE `schedule_job`  (
                                 `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务id',
                                 `bean_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'spring bean名称',
                                 `method_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '方法名',
                                 `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数',
                                 `cron` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'cron表达式',
                                 `status` tinyint(4) NULL DEFAULT NULL COMMENT '任务状态',
                                 `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                                 `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                 PRIMARY KEY (`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of schedule_job
-- ----------------------------
INSERT INTO `schedule_job` VALUES (1, 'redisSyncScheduleTask', 'SyncBlogViewsToDatabase', '1', '0 */30 * * * *', 1, '每天凌晨一点，从Redis将博客浏览量同步到数据库', '2020-11-17 07:45:42');
INSERT INTO `schedule_job` VALUES (2, 'visitorSyncScheduleTask', 'SyncVisitInfoToDatabase', '', '0 */30 * * * *', 1, '', '2021-02-04 08:14:28');
INSERT INTO `schedule_job` VALUES (10, '奇尔韦尔翁', 'SyncBlogViewsToDatabase', '', '0 */30 * * * *', 1, '', '2022-12-21 02:15:59');

-- ----------------------------
-- Table structure for schedule_job_log
-- ----------------------------
DROP TABLE IF EXISTS `schedule_job_log`;
CREATE TABLE `schedule_job_log`  (
                                     `log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志id',
                                     `job_id` bigint(20) NOT NULL COMMENT '任务id',
                                     `bean_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'spring bean名称',
                                     `method_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '方法名',
                                     `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数',
                                     `status` tinyint(4) NOT NULL COMMENT '任务执行结果',
                                     `error` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '异常信息',
                                     `times` int(11) NOT NULL COMMENT '耗时（单位：毫秒）',
                                     `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                     PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 301 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of schedule_job_log
-- ----------------------------
INSERT INTO `schedule_job_log` VALUES (17, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:00');
INSERT INTO `schedule_job_log` VALUES (18, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:01');
INSERT INTO `schedule_job_log` VALUES (19, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:02');
INSERT INTO `schedule_job_log` VALUES (21, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:04');
INSERT INTO `schedule_job_log` VALUES (22, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:05');
INSERT INTO `schedule_job_log` VALUES (23, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:06');
INSERT INTO `schedule_job_log` VALUES (24, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:07');
INSERT INTO `schedule_job_log` VALUES (25, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:08');
INSERT INTO `schedule_job_log` VALUES (26, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:09');
INSERT INTO `schedule_job_log` VALUES (27, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:10');
INSERT INTO `schedule_job_log` VALUES (28, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:11');
INSERT INTO `schedule_job_log` VALUES (29, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:12');
INSERT INTO `schedule_job_log` VALUES (30, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:13');
INSERT INTO `schedule_job_log` VALUES (31, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:14');
INSERT INTO `schedule_job_log` VALUES (32, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:15');
INSERT INTO `schedule_job_log` VALUES (33, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:16');
INSERT INTO `schedule_job_log` VALUES (34, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:17');
INSERT INTO `schedule_job_log` VALUES (35, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:18');
INSERT INTO `schedule_job_log` VALUES (36, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:19');
INSERT INTO `schedule_job_log` VALUES (37, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:20');
INSERT INTO `schedule_job_log` VALUES (38, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:21');
INSERT INTO `schedule_job_log` VALUES (39, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:22');
INSERT INTO `schedule_job_log` VALUES (40, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:23');
INSERT INTO `schedule_job_log` VALUES (41, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:24');
INSERT INTO `schedule_job_log` VALUES (42, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:16:25');
INSERT INTO `schedule_job_log` VALUES (43, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:31');
INSERT INTO `schedule_job_log` VALUES (44, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:32');
INSERT INTO `schedule_job_log` VALUES (45, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:33');
INSERT INTO `schedule_job_log` VALUES (46, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:34');
INSERT INTO `schedule_job_log` VALUES (47, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:35');
INSERT INTO `schedule_job_log` VALUES (48, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:36');
INSERT INTO `schedule_job_log` VALUES (49, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:37');
INSERT INTO `schedule_job_log` VALUES (50, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:38');
INSERT INTO `schedule_job_log` VALUES (51, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:39');
INSERT INTO `schedule_job_log` VALUES (52, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:40');
INSERT INTO `schedule_job_log` VALUES (53, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:41');
INSERT INTO `schedule_job_log` VALUES (54, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:42');
INSERT INTO `schedule_job_log` VALUES (55, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:43');
INSERT INTO `schedule_job_log` VALUES (56, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:43');
INSERT INTO `schedule_job_log` VALUES (57, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:44');
INSERT INTO `schedule_job_log` VALUES (58, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:44');
INSERT INTO `schedule_job_log` VALUES (59, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:45');
INSERT INTO `schedule_job_log` VALUES (60, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:45');
INSERT INTO `schedule_job_log` VALUES (61, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:46');
INSERT INTO `schedule_job_log` VALUES (62, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:46');
INSERT INTO `schedule_job_log` VALUES (63, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:47');
INSERT INTO `schedule_job_log` VALUES (64, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:47');
INSERT INTO `schedule_job_log` VALUES (65, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:48');
INSERT INTO `schedule_job_log` VALUES (66, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:48');
INSERT INTO `schedule_job_log` VALUES (67, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:49');
INSERT INTO `schedule_job_log` VALUES (68, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:49');
INSERT INTO `schedule_job_log` VALUES (69, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:50');
INSERT INTO `schedule_job_log` VALUES (70, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:50');
INSERT INTO `schedule_job_log` VALUES (71, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:51');
INSERT INTO `schedule_job_log` VALUES (73, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:52');
INSERT INTO `schedule_job_log` VALUES (74, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:52');
INSERT INTO `schedule_job_log` VALUES (75, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:53');
INSERT INTO `schedule_job_log` VALUES (76, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:53');
INSERT INTO `schedule_job_log` VALUES (77, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:54');
INSERT INTO `schedule_job_log` VALUES (78, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:54');
INSERT INTO `schedule_job_log` VALUES (79, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:55');
INSERT INTO `schedule_job_log` VALUES (80, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:55');
INSERT INTO `schedule_job_log` VALUES (81, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:56');
INSERT INTO `schedule_job_log` VALUES (82, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:56');
INSERT INTO `schedule_job_log` VALUES (83, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:57');
INSERT INTO `schedule_job_log` VALUES (84, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:57');
INSERT INTO `schedule_job_log` VALUES (85, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:58');
INSERT INTO `schedule_job_log` VALUES (86, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:58');
INSERT INTO `schedule_job_log` VALUES (87, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:23:59');
INSERT INTO `schedule_job_log` VALUES (88, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:23:59');
INSERT INTO `schedule_job_log` VALUES (89, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:00');
INSERT INTO `schedule_job_log` VALUES (90, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:00');
INSERT INTO `schedule_job_log` VALUES (91, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:01');
INSERT INTO `schedule_job_log` VALUES (92, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:01');
INSERT INTO `schedule_job_log` VALUES (93, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:02');
INSERT INTO `schedule_job_log` VALUES (94, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:02');
INSERT INTO `schedule_job_log` VALUES (95, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:03');
INSERT INTO `schedule_job_log` VALUES (96, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:03');
INSERT INTO `schedule_job_log` VALUES (97, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:04');
INSERT INTO `schedule_job_log` VALUES (98, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:04');
INSERT INTO `schedule_job_log` VALUES (99, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:05');
INSERT INTO `schedule_job_log` VALUES (100, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:05');
INSERT INTO `schedule_job_log` VALUES (101, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:06');
INSERT INTO `schedule_job_log` VALUES (102, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 8, '2022-12-21 10:24:06');
INSERT INTO `schedule_job_log` VALUES (103, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:07');
INSERT INTO `schedule_job_log` VALUES (104, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:07');
INSERT INTO `schedule_job_log` VALUES (105, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:08');
INSERT INTO `schedule_job_log` VALUES (106, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:08');
INSERT INTO `schedule_job_log` VALUES (107, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:09');
INSERT INTO `schedule_job_log` VALUES (108, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:09');
INSERT INTO `schedule_job_log` VALUES (109, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:10');
INSERT INTO `schedule_job_log` VALUES (110, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:10');
INSERT INTO `schedule_job_log` VALUES (111, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:11');
INSERT INTO `schedule_job_log` VALUES (112, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:11');
INSERT INTO `schedule_job_log` VALUES (113, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:12');
INSERT INTO `schedule_job_log` VALUES (114, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:12');
INSERT INTO `schedule_job_log` VALUES (115, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:13');
INSERT INTO `schedule_job_log` VALUES (116, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:13');
INSERT INTO `schedule_job_log` VALUES (117, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:14');
INSERT INTO `schedule_job_log` VALUES (118, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:14');
INSERT INTO `schedule_job_log` VALUES (119, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:15');
INSERT INTO `schedule_job_log` VALUES (120, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:15');
INSERT INTO `schedule_job_log` VALUES (121, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:16');
INSERT INTO `schedule_job_log` VALUES (122, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:16');
INSERT INTO `schedule_job_log` VALUES (123, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:17');
INSERT INTO `schedule_job_log` VALUES (124, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:17');
INSERT INTO `schedule_job_log` VALUES (125, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:18');
INSERT INTO `schedule_job_log` VALUES (126, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:18');
INSERT INTO `schedule_job_log` VALUES (127, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:19');
INSERT INTO `schedule_job_log` VALUES (128, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:19');
INSERT INTO `schedule_job_log` VALUES (129, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:20');
INSERT INTO `schedule_job_log` VALUES (130, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:20');
INSERT INTO `schedule_job_log` VALUES (131, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:21');
INSERT INTO `schedule_job_log` VALUES (132, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:21');
INSERT INTO `schedule_job_log` VALUES (133, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:22');
INSERT INTO `schedule_job_log` VALUES (134, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:22');
INSERT INTO `schedule_job_log` VALUES (135, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:23');
INSERT INTO `schedule_job_log` VALUES (136, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:23');
INSERT INTO `schedule_job_log` VALUES (137, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:24');
INSERT INTO `schedule_job_log` VALUES (138, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:24');
INSERT INTO `schedule_job_log` VALUES (139, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:25');
INSERT INTO `schedule_job_log` VALUES (140, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:25');
INSERT INTO `schedule_job_log` VALUES (141, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:26');
INSERT INTO `schedule_job_log` VALUES (142, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:26');
INSERT INTO `schedule_job_log` VALUES (143, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:27');
INSERT INTO `schedule_job_log` VALUES (144, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:27');
INSERT INTO `schedule_job_log` VALUES (145, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:28');
INSERT INTO `schedule_job_log` VALUES (146, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:28');
INSERT INTO `schedule_job_log` VALUES (147, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:29');
INSERT INTO `schedule_job_log` VALUES (148, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:29');
INSERT INTO `schedule_job_log` VALUES (149, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:30');
INSERT INTO `schedule_job_log` VALUES (150, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:30');
INSERT INTO `schedule_job_log` VALUES (151, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:31');
INSERT INTO `schedule_job_log` VALUES (152, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:31');
INSERT INTO `schedule_job_log` VALUES (153, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:32');
INSERT INTO `schedule_job_log` VALUES (154, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:32');
INSERT INTO `schedule_job_log` VALUES (155, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:33');
INSERT INTO `schedule_job_log` VALUES (156, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:33');
INSERT INTO `schedule_job_log` VALUES (157, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:34');
INSERT INTO `schedule_job_log` VALUES (158, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:34');
INSERT INTO `schedule_job_log` VALUES (159, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:35');
INSERT INTO `schedule_job_log` VALUES (160, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:35');
INSERT INTO `schedule_job_log` VALUES (161, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:36');
INSERT INTO `schedule_job_log` VALUES (162, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:36');
INSERT INTO `schedule_job_log` VALUES (163, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:37');
INSERT INTO `schedule_job_log` VALUES (164, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:37');
INSERT INTO `schedule_job_log` VALUES (165, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:38');
INSERT INTO `schedule_job_log` VALUES (166, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:38');
INSERT INTO `schedule_job_log` VALUES (167, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:39');
INSERT INTO `schedule_job_log` VALUES (168, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:39');
INSERT INTO `schedule_job_log` VALUES (169, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:40');
INSERT INTO `schedule_job_log` VALUES (170, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:40');
INSERT INTO `schedule_job_log` VALUES (171, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:41');
INSERT INTO `schedule_job_log` VALUES (172, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:41');
INSERT INTO `schedule_job_log` VALUES (173, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:42');
INSERT INTO `schedule_job_log` VALUES (174, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:42');
INSERT INTO `schedule_job_log` VALUES (175, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:24:43');
INSERT INTO `schedule_job_log` VALUES (176, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-21 10:24:43');
INSERT INTO `schedule_job_log` VALUES (177, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:41');
INSERT INTO `schedule_job_log` VALUES (178, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:42');
INSERT INTO `schedule_job_log` VALUES (179, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:43');
INSERT INTO `schedule_job_log` VALUES (180, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:44');
INSERT INTO `schedule_job_log` VALUES (181, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:45');
INSERT INTO `schedule_job_log` VALUES (182, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:46');
INSERT INTO `schedule_job_log` VALUES (183, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-21 10:27:47');
INSERT INTO `schedule_job_log` VALUES (184, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:19');
INSERT INTO `schedule_job_log` VALUES (185, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:19');
INSERT INTO `schedule_job_log` VALUES (186, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:19');
INSERT INTO `schedule_job_log` VALUES (187, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:20');
INSERT INTO `schedule_job_log` VALUES (188, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:20');
INSERT INTO `schedule_job_log` VALUES (189, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:20');
INSERT INTO `schedule_job_log` VALUES (190, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:21');
INSERT INTO `schedule_job_log` VALUES (191, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:21');
INSERT INTO `schedule_job_log` VALUES (192, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:21');
INSERT INTO `schedule_job_log` VALUES (193, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:22');
INSERT INTO `schedule_job_log` VALUES (194, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:22');
INSERT INTO `schedule_job_log` VALUES (195, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:22');
INSERT INTO `schedule_job_log` VALUES (196, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:23');
INSERT INTO `schedule_job_log` VALUES (197, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:23');
INSERT INTO `schedule_job_log` VALUES (198, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:23');
INSERT INTO `schedule_job_log` VALUES (199, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:24');
INSERT INTO `schedule_job_log` VALUES (200, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:24');
INSERT INTO `schedule_job_log` VALUES (201, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:24');
INSERT INTO `schedule_job_log` VALUES (202, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:25');
INSERT INTO `schedule_job_log` VALUES (203, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:25');
INSERT INTO `schedule_job_log` VALUES (204, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:25');
INSERT INTO `schedule_job_log` VALUES (205, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:26');
INSERT INTO `schedule_job_log` VALUES (206, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:26');
INSERT INTO `schedule_job_log` VALUES (207, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:26');
INSERT INTO `schedule_job_log` VALUES (208, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:27');
INSERT INTO `schedule_job_log` VALUES (209, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:27');
INSERT INTO `schedule_job_log` VALUES (210, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:27');
INSERT INTO `schedule_job_log` VALUES (211, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:28');
INSERT INTO `schedule_job_log` VALUES (212, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:28');
INSERT INTO `schedule_job_log` VALUES (213, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:28');
INSERT INTO `schedule_job_log` VALUES (214, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:29');
INSERT INTO `schedule_job_log` VALUES (215, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:29');
INSERT INTO `schedule_job_log` VALUES (216, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:29');
INSERT INTO `schedule_job_log` VALUES (217, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:30');
INSERT INTO `schedule_job_log` VALUES (218, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:30');
INSERT INTO `schedule_job_log` VALUES (219, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:30');
INSERT INTO `schedule_job_log` VALUES (220, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:31');
INSERT INTO `schedule_job_log` VALUES (221, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:31');
INSERT INTO `schedule_job_log` VALUES (222, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:31');
INSERT INTO `schedule_job_log` VALUES (223, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:32');
INSERT INTO `schedule_job_log` VALUES (224, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:32');
INSERT INTO `schedule_job_log` VALUES (225, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:32');
INSERT INTO `schedule_job_log` VALUES (226, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:33');
INSERT INTO `schedule_job_log` VALUES (227, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:33');
INSERT INTO `schedule_job_log` VALUES (228, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:33');
INSERT INTO `schedule_job_log` VALUES (229, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:34');
INSERT INTO `schedule_job_log` VALUES (230, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:34');
INSERT INTO `schedule_job_log` VALUES (231, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:34');
INSERT INTO `schedule_job_log` VALUES (232, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:35');
INSERT INTO `schedule_job_log` VALUES (233, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:35');
INSERT INTO `schedule_job_log` VALUES (234, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:35');
INSERT INTO `schedule_job_log` VALUES (235, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:36');
INSERT INTO `schedule_job_log` VALUES (236, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:36');
INSERT INTO `schedule_job_log` VALUES (237, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:36');
INSERT INTO `schedule_job_log` VALUES (238, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:37');
INSERT INTO `schedule_job_log` VALUES (239, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:37');
INSERT INTO `schedule_job_log` VALUES (240, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:37');
INSERT INTO `schedule_job_log` VALUES (241, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:38');
INSERT INTO `schedule_job_log` VALUES (242, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:38');
INSERT INTO `schedule_job_log` VALUES (243, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:38');
INSERT INTO `schedule_job_log` VALUES (244, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:39');
INSERT INTO `schedule_job_log` VALUES (245, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:39');
INSERT INTO `schedule_job_log` VALUES (246, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:39');
INSERT INTO `schedule_job_log` VALUES (247, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:40');
INSERT INTO `schedule_job_log` VALUES (248, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:40');
INSERT INTO `schedule_job_log` VALUES (249, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:40');
INSERT INTO `schedule_job_log` VALUES (250, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:41');
INSERT INTO `schedule_job_log` VALUES (251, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:41');
INSERT INTO `schedule_job_log` VALUES (252, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:41');
INSERT INTO `schedule_job_log` VALUES (253, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:42');
INSERT INTO `schedule_job_log` VALUES (254, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:42');
INSERT INTO `schedule_job_log` VALUES (255, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:42');
INSERT INTO `schedule_job_log` VALUES (256, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:43');
INSERT INTO `schedule_job_log` VALUES (257, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:43');
INSERT INTO `schedule_job_log` VALUES (258, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:43');
INSERT INTO `schedule_job_log` VALUES (259, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-21 10:37:44');
INSERT INTO `schedule_job_log` VALUES (260, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-21 10:37:44');
INSERT INTO `schedule_job_log` VALUES (261, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-21 10:37:44');
INSERT INTO `schedule_job_log` VALUES (262, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 03:00:01');
INSERT INTO `schedule_job_log` VALUES (263, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 03:00:01');
INSERT INTO `schedule_job_log` VALUES (264, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 03:00:01');
INSERT INTO `schedule_job_log` VALUES (265, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 03:30:01');
INSERT INTO `schedule_job_log` VALUES (266, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 03:30:01');
INSERT INTO `schedule_job_log` VALUES (267, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 03:30:01');
INSERT INTO `schedule_job_log` VALUES (268, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 04:00:01');
INSERT INTO `schedule_job_log` VALUES (269, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 04:00:01');
INSERT INTO `schedule_job_log` VALUES (270, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 04:00:01');
INSERT INTO `schedule_job_log` VALUES (271, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 04:30:01');
INSERT INTO `schedule_job_log` VALUES (272, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 04:30:01');
INSERT INTO `schedule_job_log` VALUES (273, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 04:30:01');
INSERT INTO `schedule_job_log` VALUES (274, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 05:00:01');
INSERT INTO `schedule_job_log` VALUES (275, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 05:00:01');
INSERT INTO `schedule_job_log` VALUES (276, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 05:00:01');
INSERT INTO `schedule_job_log` VALUES (277, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 05:30:01');
INSERT INTO `schedule_job_log` VALUES (278, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 05:30:01');
INSERT INTO `schedule_job_log` VALUES (279, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 05:30:01');
INSERT INTO `schedule_job_log` VALUES (280, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 06:00:01');
INSERT INTO `schedule_job_log` VALUES (281, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 06:00:01');
INSERT INTO `schedule_job_log` VALUES (282, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 06:00:01');
INSERT INTO `schedule_job_log` VALUES (283, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 06:30:01');
INSERT INTO `schedule_job_log` VALUES (284, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 06:30:01');
INSERT INTO `schedule_job_log` VALUES (285, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 06:30:01');
INSERT INTO `schedule_job_log` VALUES (286, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 07:00:01');
INSERT INTO `schedule_job_log` VALUES (287, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 07:00:01');
INSERT INTO `schedule_job_log` VALUES (288, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 07:00:01');
INSERT INTO `schedule_job_log` VALUES (289, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 08:00:01');
INSERT INTO `schedule_job_log` VALUES (290, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 08:00:01');
INSERT INTO `schedule_job_log` VALUES (291, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 08:00:01');
INSERT INTO `schedule_job_log` VALUES (292, 2, '', 'SyncVisitInfoToDatabase', '2', 1, '', 0, '2022-12-22 08:30:01');
INSERT INTO `schedule_job_log` VALUES (293, 10, '', 'SyncBlogViewsToDatabase', '10', 1, '', 0, '2022-12-22 08:30:01');
INSERT INTO `schedule_job_log` VALUES (294, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 08:30:01');
INSERT INTO `schedule_job_log` VALUES (295, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-22 08:45:30');
INSERT INTO `schedule_job_log` VALUES (296, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-22 08:45:31');
INSERT INTO `schedule_job_log` VALUES (297, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 08:45:33');
INSERT INTO `schedule_job_log` VALUES (298, 1, '', 'SyncBlogViewsToDatabase', '1', 1, '', 0, '2022-12-22 08:48:10');
INSERT INTO `schedule_job_log` VALUES (299, 2, '', 'SyncVisitInfoToDatabase', '', 1, '', 0, '2022-12-22 08:48:11');
INSERT INTO `schedule_job_log` VALUES (300, 10, '', 'SyncBlogViewsToDatabase', '', 1, '', 0, '2022-12-22 08:48:13');

-- ----------------------------
-- Table structure for site_setting
-- ----------------------------
DROP TABLE IF EXISTS `site_setting`;
CREATE TABLE `site_setting`  (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                 `name_zh` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                 `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
                                 `type` int(11) NULL DEFAULT NULL COMMENT '1基础设置，2页脚徽标，3资料卡，4友链信息',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of site_setting
-- ----------------------------
INSERT INTO `site_setting` VALUES (1, 'blogName', '博客名称', 'Naccl\'s Blog', 1);
INSERT INTO `site_setting` VALUES (2, 'webTitleSuffix', '网页标题后缀', ' - Naccl\'s Blog', 1);
INSERT INTO `site_setting` VALUES (3, 'footerImgTitle', '页脚图片标题', '手机看本站', 1);
INSERT INTO `site_setting` VALUES (4, 'footerImgUrl', '页脚图片路径', '/img/qr.png', 1);
INSERT INTO `site_setting` VALUES (5, 'copyright', 'Copyright', '{\"title\":\"Copyright © 2019 - 2022\",\"siteName\":\"NACCL\'S BLOG\"}', 1);
INSERT INTO `site_setting` VALUES (6, 'beian', 'ICP备案号', '', 1);
INSERT INTO `site_setting` VALUES (7, 'reward', '赞赏码', '/img/reward.jpg', 1);
INSERT INTO `site_setting` VALUES (8, 'commentAdminFlag', '博主评论标识', '咕咕', 1);
INSERT INTO `site_setting` VALUES (9, 'avatar', '头像', '/img/avatar.jpg', 2);
INSERT INTO `site_setting` VALUES (10, 'name', '昵称', 'Naccl', 2);
INSERT INTO `site_setting` VALUES (11, 'rollText', '滚动个签', '\"云鹤当归天，天不迎我妙木仙；\",\"游龙当归海，海不迎我自来也11。\"', 2);
INSERT INTO `site_setting` VALUES (12, 'github', 'GitHub', 'https://github.com/Naccl', 2);
INSERT INTO `site_setting` VALUES (13, 'telegram', 'Telegram', 'https://t.me/NacclOfficial', 2);
INSERT INTO `site_setting` VALUES (14, 'qq', 'QQ', 'http://sighttp.qq.com/authd?IDKEY=', 2);
INSERT INTO `site_setting` VALUES (15, 'bilibili', 'bilibili', 'https://space.bilibili.com/', 2);
INSERT INTO `site_setting` VALUES (16, 'netease', '网易云音乐', 'https://music.163.com/#/user/home?id=', 2);
INSERT INTO `site_setting` VALUES (17, 'email', 'email', 'mailto:you@example.com', 2);
INSERT INTO `site_setting` VALUES (18, 'favorite', '自定义', '{\"title\":\"最喜欢的动漫 ?\",\"content\":\"异度侵入、春物语、NO GAME NO LIFE、实力至上主义的教室、辉夜大小姐、青春猪头少年不会梦到兔女郎学姐、路人女主、Re0、魔禁、超炮、俺妹、在下坂本、散华礼弥、OVERLORD、慎勇、人渣的本愿、白色相簿2、死亡笔记、DARLING in the FRANXX、鬼灭之刃\"}', 2);
INSERT INTO `site_setting` VALUES (19, 'favorite', '自定义', '{\"title\":\"最喜欢我的女孩子们 ?\",\"content\":\"芙兰达、土间埋、食蜂操祈、佐天泪爷、樱岛麻衣、桐崎千棘、02、亚丝娜、高坂桐乃、五更琉璃、安乐冈花火、一色彩羽、英梨梨、珈百璃、时崎狂三、可儿那由多、和泉纱雾、早坂爱\"}', 2);
INSERT INTO `site_setting` VALUES (20, 'favorite', '自定义', '{\"title\":\"最喜欢玩的游戏 ?\",\"content\":\"Stellaris、巫师、GTA、荒野大镖客、刺客信条、魔兽争霸、LOL、PUBG\"}', 2);
INSERT INTO `site_setting` VALUES (21, 'badge', '徽标', '{\"title\":\"本博客已开源于 GitHub\",\"url\":\"https://github.com/Naccl/NBlog\",\"subject\":\"NBlog\",\"value\":\"Open Source\",\"color\":\"brightgreen\"}', 3);
INSERT INTO `site_setting` VALUES (22, 'badge', '徽标', '{\"title\":\"由 Spring Boot 强力驱动\",\"url\":\"https://spring.io/projects/spring-boot/\",\"subject\":\"Powered\",\"value\":\"Spring Boot\",\"color\":\"blue\"}', 3);
INSERT INTO `site_setting` VALUES (23, 'badge', '徽标', '{\"title\":\"Vue.js 客户端渲染\",\"url\":\"https://cn.vuejs.org/\",\"subject\":\"SPA\",\"value\":\"Vue.js\",\"color\":\"brightgreen\"}', 3);
INSERT INTO `site_setting` VALUES (24, 'badge', '徽标', '{\"title\":\"UI 框架 Semantic-UI\",\"url\":\"https://semantic-ui.com/\",\"subject\":\"UI\",\"value\":\"Semantic-UI\",\"color\":\"semantic-ui\"}', 3);
INSERT INTO `site_setting` VALUES (25, 'badge', '徽标', '{\"title\":\"阿里云提供服务器及域名相关服务\",\"url\":\"https://www.aliyun.com/\",\"subject\":\"VPS & DNS\",\"value\":\"Aliyun\",\"color\":\"blueviolet\"}', 3);
INSERT INTO `site_setting` VALUES (26, 'badge', '徽标', '{\"title\":\"静态资源托管于 GitHub\",\"url\":\"https://github.com/\",\"subject\":\"OSS\",\"value\":\"GitHub\",\"color\":\"github\"}', 3);
INSERT INTO `site_setting` VALUES (29, 'friendContent', '友链页面信息', '{\"value\":\"{\\\"value\\\":\\\"随机排序，不分先后。欢迎交换友链~(￣▽￣)~*\\\\n\\\\n* 昵称：Naccl\\\\n* 一句话：游龙当归海，海不迎我自来也。\\\\n* 网址：[https://naccl.top](https://naccl.top)\\\\n* 头像URL：[https://naccl.top/img/avatar.jpg](https://naccl.top/img/avatar.jpg)\\\\n\\\\n仅凭个人喜好添加友链，请在收到我的回复邮件后再于贵站添加本站链接。原则上已添加的友链不会删除，如果你发现自己被移除了，恕不另行通知，只需和我一样做就好。\\\\n\\\\n\\\"}}', 4);
INSERT INTO `site_setting` VALUES (30, 'friendCommentEnabled', '友链页面评论开关', '1', 4);
INSERT INTO `site_setting` VALUES (33, 'favorite', '自定义', '{\"title\":\"111111111\",\"content\":\"222222222222\"}', 2);

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `tag_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签颜色(可选)',
                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, '1111111111111111111', 'blue');
INSERT INTO `tag` VALUES (2, '2222', 'grey');
INSERT INTO `tag` VALUES (3, '213123', 'orange');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                         `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
                         `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像地址',
                         `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
                         `create_time` datetime(0) NOT NULL COMMENT '创建时间',
                         `update_time` datetime(0) NOT NULL COMMENT '更新时间',
                         `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色访问权限',
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '8d5f8aeeb64e3ce20b537d04c486407eaf489646617cfcf493e76f5b794fa080', 'Admin', '/img/avatar.jpg', 'admin@naccl.top', '2020-09-21 16:47:18', '2020-09-21 16:47:22', 'ROLE_admin');

-- ----------------------------
-- Table structure for visit_log
-- ----------------------------
DROP TABLE IF EXISTS `visit_log`;
CREATE TABLE `visit_log`  (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访客标识码',
                              `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
                              `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
                              `param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求参数',
                              `behavior` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问行为',
                              `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问内容',
                              `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                              `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
                              `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
                              `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
                              `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
                              `times` int(11) NOT NULL COMMENT '请求耗时（毫秒）',
                              `create_time` datetime(0) NOT NULL COMMENT '访问时间',
                              `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user-agent用户代理',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visit_log
-- ----------------------------
INSERT INTO `visit_log` VALUES (1, '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', 1, '2022-12-22 14:58:44', '1');

-- ----------------------------
-- Table structure for visit_record
-- ----------------------------
DROP TABLE IF EXISTS `visit_record`;
CREATE TABLE `visit_record`  (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `pv` int(11) NOT NULL COMMENT '访问量',
                                 `uv` int(11) NOT NULL COMMENT '独立用户',
                                 `date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期\"02-23\"',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for visitor
-- ----------------------------
DROP TABLE IF EXISTS `visitor`;
CREATE TABLE `visitor`  (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '访客标识码',
                            `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
                            `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
                            `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
                            `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
                            `create_time` datetime(0) NOT NULL COMMENT '首次访问时间',
                            `last_time` datetime(0) NOT NULL COMMENT '最后访问时间',
                            `pv` int(11) NULL DEFAULT NULL COMMENT '访问页数统计',
                            `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user-agent用户代理',
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE INDEX `idx_uuid`(`uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitor
-- ----------------------------
INSERT INTO `visitor` VALUES (1, '1', '1', '1', '1', '1', '2022-12-22 16:40:24', '2022-12-22 16:40:30', 1, '1');
INSERT INTO `visitor` VALUES (2, '2', '1', '1', '1', '1', '2022-12-22 16:40:24', '2022-12-22 16:40:30', 1, '1');

SET FOREIGN_KEY_CHECKS = 1;
