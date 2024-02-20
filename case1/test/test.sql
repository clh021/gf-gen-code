-- 测试 sqlite 有关表字段注释的支持程度
-- rm test.db.sqlite3; sqlite3 test.db.sqlite3 ".read test.sql"
-- sqlite3 test.db.sqlite3 ".output sqlite3.test.data.sql" ".dump"
-- 是可以完全支持的，只是不像 mysql 那样的支持方式

-- sqlite3 hotgo.db.sqlite3
-- SQLite version 3.45.1 2024-01-30 16:01:20
-- Enter ".help" for usage hints.
-- sqlite> .databases
-- main: /home/lee/Projects/compat-detect-tools/serve/resource/hotgo.db.sqlite3 r/w
-- sqlite>.tables
-- sqlite>.schema sqlite_master
-- sqlite> .header on
-- sqlite> .mode column
-- sqlite> .timer on
-- sqlite> PRAGMA table_info(users2);
-- cid  name        type          notnull  dflt_value  pk
-- ---  ----------  ------------  -------  ----------  --
-- 0    id          INTEGER       0                    1
-- 1    name__cmt   varchar(255)  0                    0
-- 2    age__cmt    INTEGER       0                    0
-- 3    email__cmt  TEXT          0                    0
-- Run Time: real 0.000 user 0.000101 sys 0.000067
-- sqlite> PRAGMA table_info(users);
-- cid  name         type                   notnull  dflt_value  pk
-- ---  -----------  ---------------------  -------  ----------  --
-- 0    id           INTEGER                0                    1
-- 1    name         varchar COMMENT '用户名'  0                    0
-- 2    age          INTEGER COMMENT '年龄'   0                    0
-- 3    email        TEXT COMMENT '邮箱'      0                    0
-- 4    category_id  bigint COMMENT '分类ID'  0                    0
-- 5    flag         json COMMENT '标签'      0                    0
-- 6    title        varchar COMMENT '标题'   0                    0
-- 7    description  varchar COMMENT '描述'   0                    0
-- 8    content      longtext COMMENT '内容'  0                    0
-- 9    image        varchar COMMENT '单图'   0                    0
-- 10   images       json COMMENT '多图'      0                    0
-- 11   attachfile   varchar COMMENT '附件'   0                    0
-- 12   attachfiles  json  COMMENT '多附件'    0                    0
-- 13   map          json COMMENT '动态键值对'   0        NULL        0
-- 14   star         decimal COMMENT '推荐星'  0        '0.0'       0
-- 15   price        decimal COMMENT '价格'   1        '0.00'      0
-- Run Time: real 0.000 user 0.000000 sys 0.000223
-- sqlite> PRAGMA table_info(credits);
-- cid  name          type           notnull  dflt_value  pk
-- ---  ------------  -------------  -------  ----------  --
-- 0    id            bigint         1                    0
-- 1    member_id     bigint         0        '0'         0
-- 2    app_id        varchar(64)    0        ''          0
-- 3    addons_name   varchar(100)   1        ''          0
-- 4    credit_type   varchar(32)    1        ''          0
-- 5    credit_group  varchar(32)    0        ''          0
-- 6    before_num    decimal(10,2)  0        '0.00'      0
-- 7    num           decimal(10,2)  0        '0.00'      0
-- 8    after_num     decimal(10,2)  0        '0.00'      0
-- 9    remark        varchar(255)   0        ''          0
-- 10   ip            varchar(20)    0        ''          0
-- 11   map_id        bigint         0        '0'         0
-- 12   status        tinyint(1)     0        '1'         0
-- 13   created_at    datetime       0        NULL        0
-- 14   updated_at    datetime       0        NULL        0
-- Run Time: real 0.000 user 0.000226 sys 0.000000
-- sqlite> SELECT * FROM sqlite_master;
-- type   name                  tbl_name              rootpage  sql
-- -----  --------------------  --------------------  --------  ------------------------------------------------------
-- table  credits  credits  2         CREATE TABLE `credits` (
--                                                                `id` bigint NOT NULL,
--                                                                `member_id` bigint DEFAULT '0',
--                                                                `app_id` varchar(64) DEFAULT '',
--                                                                `addons_name` varchar(100) NOT NULL DEFAULT '',
--                                                                `credit_type` varchar(32) NOT NULL DEFAULT '',
--                                                                `credit_group` varchar(32) DEFAULT '',
--                                                                `before_num` decimal(10,2) DEFAULT '0.00',
--                                                                `num` decimal(10,2) DEFAULT '0.00',
--                                                                `after_num` decimal(10,2) DEFAULT '0.00',
--                                                                `remark` varchar(255) DEFAULT '',
--                                                                `ip` varchar(20) DEFAULT '',
--                                                                `map_id` bigint DEFAULT '0',
--                                                                `status` tinyint(1) DEFAULT '1',
--                                                                `created_at` datetime DEFAULT NULL,
--                                                                `updated_at` datetime DEFAULT NULL
--                                                              )

-- table  users                 users                 3         CREATE TABLE users (
--                                                                  id INTEGER PRIMARY KEY,
--                                                                  name varchar COMMENT '用户名', -- User name
--                                                                  age INTEGER COMMENT '年龄',    -- 年龄
--                                                                  email TEXT COMMENT '邮箱',     -- 邮箱
--                                                                  category_id bigint COMMENT '分类ID',
--                                                                  flag json COMMENT '标签',
--                                                                  title varchar COMMENT '标题',
--                                                                  description varchar COMMENT '描述',
--                                                                  content longtext COMMENT '内容',
--                                                                  image varchar COMMENT '单图',
--                                                                  images json COMMENT '多图',
--                                                                  attachfile varchar COMMENT '附件',
--                                                                  attachfiles json  COMMENT '多附件',
--                                                                  map json COMMENT '动态键值对' DEFAULT NULL,
--                                                                  star decimal COMMENT '推荐星' DEFAULT '0.0',
--                                                                  price decimal COMMENT '价格' NOT NULL DEFAULT '0.00'
--                                                              )

-- table  users2                users2                4         CREATE TABLE users2 (
--                                                                  id INTEGER PRIMARY KEY,
--                                                                  name__cmt varchar(255),  -- User name
--                                                                  age__cmt INTEGER,        -- User age
--                                                                  email__cmt TEXT          -- User email
--                                                              )
-- Run Time: real 0.001 user 0.000233 sys 0.000166

-- 思路:
-- 通过 `--` 的方式编写注释，不要尝试使用  COMMENT ，会破坏 column 的 type。
-- 然后可以很容易的获得 表创建语句，分析创建语句，得到字段的备注信息。
-- 进而可以进行下一步操作了。

CREATE TABLE IF NOT EXISTS `credits` (
  `id` bigint NOT NULL,
  `member_id` bigint DEFAULT '0',
  `app_id` varchar(64) DEFAULT '',
  `addons_name` varchar(100) NOT NULL DEFAULT '',
  `credit_type` varchar(32) NOT NULL DEFAULT '',
  `credit_group` varchar(32) DEFAULT '',
  `before_num` decimal(10,2) DEFAULT '0.00',
  `num` decimal(10,2) DEFAULT '0.00',
  `after_num` decimal(10,2) DEFAULT '0.00',
  `remark` varchar(255) DEFAULT '',
  `ip` varchar(20) DEFAULT '',
  `map_id` bigint DEFAULT '0',
  `status` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
);

-- COMMENT 和 DEFAULT NULL 默认值设置一起使用时，要注意 默认值设置排在最后
-- COMMENT 不能和 varchar(255), decimal(5,1) 括号之类一起使用，但是可以在之后添加注释
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name varchar COMMENT '用户名', -- User name
    age INTEGER COMMENT '年龄',    -- 年龄
    email TEXT COMMENT '邮箱',     -- 邮箱
    category_id bigint COMMENT '分类ID',
    flag json COMMENT '标签',
    title varchar COMMENT '标题',
    description varchar COMMENT '描述',
    content longtext COMMENT '内容',
    image varchar COMMENT '单图',
    images json COMMENT '多图',
    attachfile varchar COMMENT '附件',
    attachfiles json  COMMENT '多附件',
    map json COMMENT '动态键值对' DEFAULT NULL,
    star decimal COMMENT '推荐星' DEFAULT '0.0',
    price decimal COMMENT '价格' NOT NULL DEFAULT '0.00'
);
-- COMMENT ON COLUMN users.name IS '用户名';
-- COMMENT ON COLUMN users.price IS '价格';
CREATE TABLE users2 (
    id INTEGER PRIMARY KEY,
    name__cmt varchar(255),  -- User name
    age__cmt INTEGER,        -- User age
    email__cmt TEXT          -- User email
);

-- CREATE TABLE IF NOT EXISTS `hg_addon_hgexample_table` (
--   `id` bigint NOT NULL COMMENT 'ID',
--   `category_id` bigint NOT NULL COMMENT '分类ID',
--   `flag` json DEFAULT NULL COMMENT '标签',
--   `title` varchar(255) NOT NULL COMMENT '标题',
--   `description` varchar(255) NOT NULL COMMENT '描述',
--   `content` longtext NOT NULL COMMENT '内容',
--   `image` varchar(255) DEFAULT NULL COMMENT '单图',
--   `images` json DEFAULT NULL COMMENT '多图',
--   `attachfile` varchar(255) DEFAULT NULL COMMENT '附件',
--   `attachfiles` json DEFAULT NULL COMMENT '多附件',
--   `map` json DEFAULT NULL COMMENT '动态键值对',
--   `star` decimal(5,1) DEFAULT '0.0' COMMENT '推荐星',
--   `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
--   `views` bigint DEFAULT NULL COMMENT '浏览次数',
--   `activity_at` date DEFAULT NULL COMMENT '活动时间',
--   `start_at` datetime DEFAULT NULL COMMENT '开启时间',
--   `end_at` datetime DEFAULT NULL COMMENT '结束时间',
--   `switch` tinyint(1) DEFAULT NULL COMMENT '开关',
--   `sort` int NOT NULL COMMENT '排序',
--   `avatar` varchar(255) DEFAULT '' COMMENT '头像',
--   `sex` tinyint(1) DEFAULT NULL COMMENT '性别',
--   `qq` varchar(20) DEFAULT '' COMMENT 'qq',
--   `email` varchar(60) DEFAULT '' COMMENT '邮箱',
--   `mobile` varchar(20) DEFAULT '' COMMENT '手机号码',
--   `hobby` json DEFAULT NULL COMMENT '爱好',
--   `channel` int NOT NULL DEFAULT '1' COMMENT '渠道',
--   `city_id` bigint DEFAULT '0' COMMENT '所在城市',
--   `pid` bigint NOT NULL COMMENT '上级ID',
--   `level` int DEFAULT '1' COMMENT '树等级',
--   `tree` varchar(512) NOT NULL COMMENT '关系树',
--   `remark` varchar(255) DEFAULT NULL COMMENT '备注',
--   `status` tinyint(1) DEFAULT '1' COMMENT '状态',
--   `created_by` bigint DEFAULT '0' COMMENT '创建者',
--   `updated_by` bigint DEFAULT '0' COMMENT '更新者',
--   `created_at` datetime DEFAULT NULL COMMENT '创建时间',
--   `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
--   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间'
-- ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='插件_案例_表格';