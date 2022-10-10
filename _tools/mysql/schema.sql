CREATE TABLE `user`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name` VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`     VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utfmb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `title`　　　　 VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`　　　　VARCHAR(20) NOT NULL COMMENT 'タスクの状態',
    `created_at`　 DATETIME(6) NOT NULL COMMENT 'レコードの作成日時',
    `modified_at`　DATETIME(6) NOT NULL COMMENT 'レコードの修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';