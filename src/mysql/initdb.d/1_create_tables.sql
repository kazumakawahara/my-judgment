USE my_judgment_db;

CREATE TABLE `users`
(
    `id`         INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `gender`     CHAR(5) DEFAULT NULL COMMENT '性別',
    `address`    CHAR(5)     NOT NULL COMMENT '所在地',
    `plan`       int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '利用プラン',
    `created_at` DATETIME    NOT NULL COMMENT '作成日時',
    `created_by` INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at` DATETIME    NOT NULL COMMENT '更新日時',
    `updated_by` INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    PRIMARY KEY (`id`),
    UNIQUE `uq_users_1` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'ユーザー';
