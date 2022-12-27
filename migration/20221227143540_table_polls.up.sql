CREATE TABLE IF NOT EXISTS polls
(
    `id`              BIGINT(20)      NOT NULL AUTO_INCREMENT,
    `user_id`         BIGINT(20)      NOT NULL,
    `poll_policy`     BIGINT(1)       NOT NULL DEFAULT 0,
    `poll_title`      VARCHAR(255)    NOT NULL UNIQUE,
    `poll_vote_type`  BIGINT(1)       NOT NULL DEFAULT 0,
    `created_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`      TIMESTAMP       NULL     DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
