/*
------------------------------------------------
------------- sku_info_tab   --------------
------------------------------------------------
 */
CREATE TABLE `sku_info_tab`
(
    `id`            bigint(64) unsigned auto_increment primary key not null,
    `sku_id`        varchar(64) unique key                         not null default '',
    `sku_name`      varchar(256)                                   not null default '',
    `item_id`       int(32)                                        not null default 0,
    `item_name`     varchar(128)                                   not null default '',
    `product_id`    int(32) unsigned unique key                    not null default 0,
    `product_name`  varchar(128)                                   not null default '',
    `category_id`   mediumint(16) unsigned                         not null default 0,
    `category_name` char(32)                                       not null default '',
    `shop_id`       int(32) unsigned                               not null default 0,
    `ctime`         int(32) unsigned                               not null default 0,
    `mtime`         int(32) unsigned                               not null default 0
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;

/*
------------------------------------------------
------------- item_info_tab   --------------
------------------------------------------------
 */
CREATE TABLE `item_info_tab`
(
    `id`            bigint(64) unsigned auto_increment primary key not null,
    `item_id`       int(32) unsigned unique key                    not null default 0,
    `item_name`     varchar(128)                                   not null default '',
    `category_id`   mediumint(16) unsigned                         not null default 0, /* 品类可以按层级展开，目前定位一级*/
    `category_name` char(32)                                       not null default '', /*使用char，确保品类是一个固定字段，不能设置过长字段来展示*/
    `shop_id`       int(32) unsigned                               not null default 0,
    `ctime`         int(32) unsigned                               not null default 0,
    `mtime`         int(32) unsigned                               not null default 0
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;
/*
------------------------------------------------
------------- sku_inventory_tab   --------------
------------------------------------------------
 */
CREATE TABLE `sku_inventory_tab`
(
    `id`        bigint(64) unsigned auto_increment primary key not null,
    `sku_id`    varchar(64) unique key                         not null default '',
    `inventory` int(32) unsigned                               not null default 0,
    `damage`    int(32) unsigned                               not null default 0,
    `ctime`     int(32) unsigned                               not null default 0,
    `mtime`     int(32) unsigned                               not null default 0
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;