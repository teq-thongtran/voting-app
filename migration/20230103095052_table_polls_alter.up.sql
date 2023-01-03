ALTER TABLE `polls` 
  CHANGE COLUMN `poll_policy` `poll_policy` VARCHAR(45) NOT NULL DEFAULT '0' ,
  CHANGE COLUMN `poll_vote_type` `poll_vote_type` VARCHAR(45) NOT NULL DEFAULT '0' ;
