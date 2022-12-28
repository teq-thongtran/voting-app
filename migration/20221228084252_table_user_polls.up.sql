CREATE TABLE IF NOT EXISTS user_polls
(
    `id`            BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `poll_id`       BIGINT(20)   NOT NULL,
    `user_id`       BIGINT(20)   NOT NULL,
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (poll_id) REFERENCES polls(id),
    FOREIGN KEY (user_id)   REFERENCES users(id),
    UNIQUE INDEX `user_poll_uq` (`poll_id` ASC, `user_id` ASC) VISIBLE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
