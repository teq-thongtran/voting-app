ALTER TABLE user_votes  
  ADD CONSTRAINT `user_votes_ibfk_1` FOREIGN KEY (`poll_option_id`) REFERENCES `voting`.`poll_options` (`id`);
