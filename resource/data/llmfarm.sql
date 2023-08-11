

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `bot`
-- ----------------------------
DROP TABLE IF EXISTS `bot`;
CREATE TABLE `bot` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '机器人ID',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `botName` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '机器人名称',
  `chainId` int NOT NULL COMMENT '流程ID',
  `userType` enum('system','user') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'user' COMMENT '用户类型',
  `userId` int NOT NULL COMMENT '用户Id',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
  `isDefault` tinyint NOT NULL DEFAULT '0' COMMENT '是否默认选中',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`userId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='机器人表';

-- ----------------------------
--  Records of `bot`
-- ----------------------------
BEGIN;
INSERT INTO `bot` VALUES ('1', '2023-08-10 16:45:29', '2023-08-10 16:45:29', null, '万能3.5', '1', 'user', '1', '', '0');
COMMIT;

-- ----------------------------
--  Table structure for `chain`
-- ----------------------------
DROP TABLE IF EXISTS `chain`;
CREATE TABLE `chain` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '流程ID',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `userId` int NOT NULL COMMENT '用户ID',
  `chainName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '流程名称',
  `nodes` json DEFAULT NULL COMMENT '节点数据',
  `edges` json DEFAULT NULL COMMENT '边缘数据',
  `userType` enum('system','user') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'user' COMMENT '用户类型',
  `useKnowledge` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否使用知识库',
  `chainLevel` int NOT NULL DEFAULT '0' COMMENT '层级',
  `fatherId` int DEFAULT NULL COMMENT '父ID，记录所属模版',
  `fatherTemplateId` int DEFAULT NULL COMMENT '记录所属具体某个vesion的模版',
  `isUpload` tinyint NOT NULL DEFAULT '0' COMMENT '0 没上传 1 已上传',
  `chainTemplateId` int DEFAULT NULL COMMENT '上传后的模版ID',
  `isInMarket` tinyint NOT NULL DEFAULT '0' COMMENT '是否已上架 0:未上架 1:已上架',
  `chainTemplateInfoId` int DEFAULT NULL COMMENT '上架后的模版ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='流程表';

-- ----------------------------
--  Records of `chain`
-- ----------------------------
BEGIN;
INSERT INTO `chain` VALUES ('1', '2023-08-10 16:44:37', '2023-08-10 16:45:29', null, '1', '万能3.5', '[{\"id\": \"768cc05d-94aa-4f0f-abc7-f788cf1eee03\", \"data\": {\"desc\": \"用户输入内容\", \"type\": \"UserInputText\", \"title\": \"UserInputText\", \"value\": {}}, \"size\": {\"width\": 300, \"height\": 0}, \"ports\": {\"items\": [{\"id\": \"linked\", \"attrs\": {\"circle\": {\"r\": 6, \"style\": {\"stroke\": \"#52c41a\"}, \"magnet\": true}}, \"group\": \"linked\"}, {\"id\": \"text\", \"attrs\": {\"circle\": {\"r\": 6, \"style\": {\"stroke\": \"#52c41a\"}, \"magnet\": true}}, \"group\": \"text\"}], \"groups\": {\"text\": {\"position\": {\"args\": {\"dy\": 140}, \"name\": \"right\"}}, \"linked\": {\"position\": {\"args\": {\"dy\": 30}, \"name\": \"left\"}}}}, \"shape\": \"flow-node\", \"position\": {\"x\": -420, \"y\": -220}}, {\"id\": \"77d513fa-fc2e-494d-b898-a04ee5763700\", \"data\": {\"desc\": \"Call OpenAI LLM\", \"type\": \"ChatCompletion\", \"title\": \"ChatCompletion\", \"value\": {\"modelName\": \"gpt-3.5-turbo\"}}, \"size\": {\"width\": 300, \"height\": 0}, \"ports\": {\"items\": [{\"id\": \"linked\", \"attrs\": {\"circle\": {\"r\": 6, \"style\": {\"stroke\": \"#52c41a\"}, \"magnet\": true}}, \"group\": \"linked\"}, {\"id\": \"prompt\", \"attrs\": {\"circle\": {\"r\": 6, \"style\": {\"stroke\": \"#52c41a\"}, \"magnet\": true}}, \"group\": \"prompt\"}, {\"id\": \"ChatOpenAI\", \"attrs\": {\"circle\": {\"r\": 6, \"style\": {\"stroke\": \"#52c41a\"}, \"magnet\": true}}, \"group\": \"ChatOpenAI\"}], \"groups\": {\"linked\": {\"position\": {\"args\": {\"dy\": 30}, \"name\": \"left\"}}, \"prompt\": {\"position\": {\"args\": {\"dy\": 320}, \"name\": \"left\"}}, \"ChatOpenAI\": {\"position\": {\"args\": {\"dy\": 380}, \"name\": \"right\"}}}}, \"shape\": \"flow-node\", \"position\": {\"x\": -20, \"y\": -180}}]', '[{\"id\": \"43179558-baf4-4bc4-a7bb-718548ab5c07\", \"attrs\": {\"line\": {\"stroke\": \"#c2c8d5\", \"targetMarker\": {\"name\": \"block\", \"size\": 8}}}, \"shape\": \"next\", \"router\": {\"name\": \"normal\"}, \"source\": {\"cell\": \"768cc05d-94aa-4f0f-abc7-f788cf1eee03\", \"port\": \"text\"}, \"target\": {\"cell\": \"77d513fa-fc2e-494d-b898-a04ee5763700\", \"port\": \"prompt\"}, \"zIndex\": 0, \"connector\": {\"args\": {\"radius\": 4}, \"name\": \"smooth\"}}]', 'user', '0', '0', '0', '0', '0', '0', '0', '0'), ('2', '2023-08-10 16:45:11', '2023-08-10 16:46:06', '2023-08-10 16:46:06', '1', '万能3.5', '[]', '[]', 'user', '0', '0', '0', '0', '0', '0', '0', '0');
COMMIT;

-- ----------------------------
--  Table structure for `chat`
-- ----------------------------
DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '聊天ID',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `chatName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '对话名称',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `contextType` enum('单轮对话','连续对话') DEFAULT NULL COMMENT '上下文类型',
  `isTemp` tinyint NOT NULL DEFAULT '0' COMMENT '是否是临时会话 0:否 1:是',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`userId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4443 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='对话表';


-- ----------------------------
--  Table structure for `message`
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '消息ID',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `userId` int NOT NULL COMMENT '用户ID',
  `type` enum('BOT','USER') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'USER' COMMENT '消息发送者类型',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '消息内容',
  `contentType` enum('文本') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '文本' COMMENT '消息内容类型',
  `parentId` int DEFAULT NULL COMMENT '父消息ID',
  `mark` enum('未标记','好评','差评') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '未标记' COMMENT '标记',
  `feedback` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '反馈信息',
  `active` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
  `chatId` int NOT NULL COMMENT '所属对话id',
  `botId` int NOT NULL COMMENT '机器人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28548 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='聊天消息表';


-- ----------------------------
--  Table structure for `tokenUsageLimit`
-- ----------------------------
DROP TABLE IF EXISTS `tokenUsageLimit`;
CREATE TABLE `tokenUsageLimit` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用量使用额度ID',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `dailyTokens` int NOT NULL DEFAULT '4000' COMMENT 'Tokens',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`userId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2714 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用量使用额度表';

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleteTime` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号码',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `email` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码盐值',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int NOT NULL DEFAULT '0' COMMENT '生日',
  `user_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint NOT NULL DEFAULT '1' COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `is_membership` tinyint DEFAULT '0' COMMENT '是否会员',
  `member_deadline` datetime DEFAULT NULL COMMENT '会员截止时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';

-- ----------------------------
--  Records of `user`
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('1', '2023-08-10 16:15:51', '2023-08-10 16:15:54', null, 'llmfarm', '15910771917', 'bc9f63dc9681adb2fe48f4d1ed675b9b', null, '2fAIIn', '', '0', '1', '2', '', '204', '', '1', '', '', '', null, '0', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
