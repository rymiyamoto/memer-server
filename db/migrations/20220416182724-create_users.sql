
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id`              BIGINT       UNSIGNED AUTO_INCREMENT,
  `is_authed`       BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '認証済みであるか',
  `name`            VARCHAR(64)                         COMMENT 'ユーザー名',
  `deleted_at`      DATETIME                            COMMENT '退会日',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

-- +migrate Down
DROP TABLE IF EXISTS `users`;
