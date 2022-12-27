CREATE TABLE IF NOT EXISTS poll_options
(
    `id`            BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `poll_id`       BIGINT(20)   NOT NULL,
    `user_id`       BIGINT(20)   NOT NULL,
    `option_text`   VARCHAR(255) NOT NULL,
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (poll_id) REFERENCES polls(id),
    FOREIGN key (user_id) REFERENCES users(id),
    UNIQUE INDEX `poll_option_uq` (`poll_id` ASC, `option_text` ASC) VISIBLE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
