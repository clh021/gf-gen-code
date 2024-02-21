-- 创建一个书本管理的表格
CREATE TABLE IF NOT EXISTS `books` (
  `id` bigint NOT NULL,                             -- 编号
  `author_id` bigint DEFAULT '0',                   -- 会员编号
  `class` varchar(64) DEFAULT '',                   -- 应用编号
  `name` varchar(255) DEFAULT '',                   -- 应用编号
  `price` decimal(10,2) DEFAULT '0.00',             -- 年龄
  `cover_fid` bigint DEFAULT '0',                   -- 封面图片|上传图片文件的编号
  `author_fid` varchar(100) NOT NULL DEFAULT '',    -- 作者图片|上传图片文件的编号
  `status` tinyint(1) DEFAULT '1',                  -- 状态|1:正常|0:删除
  `created_at` datetime DEFAULT NULL,               -- 添加时间
  `updated_at` datetime DEFAULT NULL,               -- 更新时间
  `deleted_at` datetime DEFAULT NULL                -- 删除时间
);