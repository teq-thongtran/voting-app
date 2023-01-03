ALTER TABLE user_votes
  DROP FOREIGN KEY `user_votes_ibfk_1`,
  CHANGE COLUMN `option_id` `poll_option_id` BIGINT NOT NULL;
