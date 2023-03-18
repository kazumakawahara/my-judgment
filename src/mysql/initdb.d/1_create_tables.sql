USE my_judgment_db;

CREATE TABLE `users`
(
    `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `birthday`    DATETIME NOT NULL COMMENT '誕生日',
    `gender`      CHAR(5) NOT NULL COMMENT '性別',
    `address`     CHAR(5) NOT NULL COMMENT '所在地',
    `email`       varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'メールアドレス',
    `password`    varchar(16)  NOT NULL COMMENT 'パスワード',
    `plan`        int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '利用プラン',
    `created_at`  DATETIME NOT NULL COMMENT '作成日時',
    `created_by`  INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`  DATETIME NOT NULL COMMENT '更新日時',
    `updated_by`  INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `deleted_at`  DATETIME DEFAULT NULL COMMENT '削除日時',
    `deleted_by`  INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    KEY `index_users_1` (`name`, `email`),
    UNIQUE `uq_users_1` (`email`, `deleted_uts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'ユーザー';

CREATE TABLE `groups`
(
    `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(11) UNSIGNED NOT NULL COMMENT 'ユーザーID',
    `name`        VARCHAR(20) NOT NULL COMMENT 'グループ名',
    `created_at`  DATETIME    NOT NULL COMMENT '作成日時',
    `created_by`  INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`  DATETIME    NOT NULL COMMENT '更新日時',
    `updated_by`  INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `deleted_at`  DATETIME DEFAULT NULL COMMENT '削除日時',
    `deleted_by`  INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    UNIQUE `uq_groups_1` (`user_id`, `name`, `deleted_uts`),
    FOREIGN KEY `fk_groups_user_id` (`user_id`)
        REFERENCES `users` (`id`)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'グループ';

CREATE TABLE `categories`
(
    `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`    INT(11) UNSIGNED NOT NULL COMMENT 'グループID',
    `name`        VARCHAR(20) NOT NULL COMMENT 'カテゴリー名',
    `detail`      VARCHAR(40) DEFAULT NULL COMMENT 'カテゴリー詳細',
    `position`    INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '並び順',
    `created_at`  DATETIME    NOT NULL COMMENT '作成日時',
    `created_by`  INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`  DATETIME    NOT NULL COMMENT '更新日時',
    `updated_by`  INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `deleted_at`  DATETIME    DEFAULT NULL COMMENT '削除日時',
    `deleted_by`  INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    UNIQUE `uq_categories_1` (`group_id`, `name`, `deleted_uts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'カテゴリー';

CREATE TABLE `items`
(
    `id`               INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`         INT(11) UNSIGNED NOT NULL COMMENT 'グループID',
    `category_id`      INT(11) UNSIGNED NOT NULL COMMENT 'カテゴリーID',
    `name`             VARCHAR(20) NOT NULL COMMENT 'アイテム名',
    `area`             CHAR(5)  DEFAULT NULL COMMENT 'アイテムエリア',
    `feel`             CHAR(5)  DEFAULT NULL COMMENT '気分',
    `number_of_people` CHAR(5)  DEFAULT NULL COMMENT '人数',
    `gender_of _pair`  CHAR(5)  DEFAULT NULL COMMENT '性別ペア',
    `position`         INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '並び順',
    `created_at`       DATETIME    NOT NULL COMMENT '作成日時',
    `created_by`       INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`       DATETIME    NOT NULL COMMENT '更新日時',
    `updated_by`       INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `trashed_at`       DATETIME DEFAULT NULL COMMENT 'ゴミ箱入日時',
    `trashed_by`       INT(11) UNSIGNED DEFAULT NULL COMMENT 'ゴミ箱に入れたユーザーID',
    `restored_at`      DATETIME DEFAULT NULL COMMENT 'ゴミ箱復元日時',
    `restored_by`      INT(11) UNSIGNED DEFAULT NULL COMMENT 'ゴミ箱復元ユーザーID',
    `deleted_at`       DATETIME DEFAULT NULL COMMENT '削除日時',
    `deleted_by`       INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts`      BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    KEY `index_items_1` (`group_id`, `deleted_uts`, `category_id`),
    UNIQUE `uq_items_1` (`group_id`, `category_id`, `name`, `deleted_uts`),
    FOREIGN KEY `fk_items_category_id` (`category_id`)
        REFERENCES `categories` (`id`)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'アイテム';

CREATE TABLE `item_favorites`
(
    `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`    INT(11) UNSIGNED NOT NULL COMMENT 'グループID',
    `item_id`     INT(11) UNSIGNED NOT NULL COMMENT 'アイテムID',
    `use_count`   INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用回数',
    `created_at`  DATETIME NOT NULL COMMENT '作成日時',
    `created_by`  INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`  DATETIME NOT NULL COMMENT '更新日時',
    `updated_by`  INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `deleted_at`  DATETIME DEFAULT NULL COMMENT '削除日時',
    `deleted_by`  INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    UNIQUE `uq_categories_1` (`group_id`, `item_id`, `deleted_uts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'よく使うアイテム';

CREATE TABLE `locks`
(
    `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`    INT(11) UNSIGNED NOT NULL COMMENT 'グループID',
    `type`        VARCHAR(5)  NOT NULL COMMENT 'ロックタイプ',
    `target`      VARCHAR(30) NOT NULL COMMENT 'ロック対象',
    `target_id`   INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'ロック対象ID',
    `expires_uts` INT(11) NOT NULL DEFAULT 0 COMMENT 'ロック有効期限UNIXタイムスタンプ',
    `device`      VARCHAR(64) NOT NULL COMMENT 'ロックデバイス名',
    `created_at`  DATETIME    NOT NULL COMMENT '作成日時',
    `created_by`  INT(11) UNSIGNED NOT NULL COMMENT '作成ユーザーID',
    `updated_at`  DATETIME    NOT NULL COMMENT '更新日時',
    `updated_by`  INT(11) UNSIGNED NOT NULL COMMENT '更新ユーザーID',
    `deleted_at`  DATETIME DEFAULT NULL COMMENT '削除日時',
    `deleted_by`  INT(11) UNSIGNED DEFAULT NULL COMMENT '削除ユーザーID',
    `deleted_uts` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '削除日時UNIX NANOタイムスタンプ',
    PRIMARY KEY (`id`),
    UNIQUE `uq_locks_1` (`group_id`, `deleted_uts`, `type`, `target`, `target_id`, `expires_uts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT 'ロック';

