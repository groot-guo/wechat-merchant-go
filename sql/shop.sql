/*------------------------------------------------
------------- shop_info_tab   --------------
------------------------------------------------
 */
CREATE TABLE `shop_info_tab`
(
    `id`        bigint(64) unsigned primary key auto_increment not null,
    `shop_id`   int(32) unsigned unique key                    not null default 0,
    `shop_name` varchar(256)                                   not null default '',
    `ctime`     int(32) unsigned                               not null default 0,
    `mtime`     int(32) unsigned                               not null default 0
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;

