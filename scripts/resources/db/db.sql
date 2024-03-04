CREATE TABLE IF NOT EXISTS `user` (                 -- 用户管理
  `id` INTEGER,                                     -- 编号
  `name` VARCHAR(50),                               -- 名称
  `email` VARCHAR(255),                             -- 邮箱
  `address` TEXT,                                   -- 地址
  `salt` VARCHAR(50),                               -- 盐
  `password` VARCHAR(50),                           -- 密码
  `mark` VARCHAR(255),                              -- 备注
  `permission` TEXT,                                -- 权限
  `created_user_id` INTEGER,                        -- 创建者编号
  `created_at` DATETIME,                            -- 创建时间
  `updated_at` DATETIME,                            -- 更新时间
  `deleted_at` DATETIME,                            -- 删除时间
  PRIMARY KEY(`id`)
);
INSERT INTO user VALUES(1,'admin','admin@admin.com','中国北京','2jnb22','admin','管理员','["create","history","env","user"]',NULL,'2024-01-01 01:01:01',NULL,NULL);
INSERT INTO user VALUES(2,'user','user@user.com','中国上海','2a33ce','user','普通用户','["create","history","env"]',NULL,'2024-01-01 01:01:01',NULL,NULL);
INSERT INTO user VALUES(3,'guest','guest@guest.com','中国武汉','242s9e','guest','访客用户','["history"]',NULL,'2024-01-01 01:01:01',NULL,NULL);
CREATE TABLE IF NOT EXISTS `book` (                 -- 书本管理
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
  `deleted_at` datetime DEFAULT NULL,               -- 删除时间
  PRIMARY KEY(`id`)
);
CREATE TABLE IF NOT EXISTS `file` (                 -- 文件管理
  `id` INTEGER,                                     -- 编号
  `name` VARCHAR(55),                               -- 名称
  `content_type` VARCHAR(255),                      -- 文件类型
  `size` INTEGER,                                   -- 大小
  `filehash` VARCHAR(255),                          -- 文件hash
  `filepath` VARCHAR(255),                          -- 文件路径
  `url` VARCHAR(255),                               -- 文件url
  `user_id` INTEGER,                                -- 用户编号
  `create_at` DATETIME,                             -- 创建时间
  `updated_at` DATETIME,                            -- 更新时间
  `deleted_at` DATETIME,                            -- 删除时间
  PRIMARY KEY(`id`)
);