DROP DATABASE IF EXISTS `sns`;
CREATE DATABASE `sns` DEFAULT CHARACTER SET utf8mb4;

USE `sns`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` ( 
  `UserID` int NOT NULL AUTO_INCREMENT,
  `Username` VARCHAR(15) NOT NULL,
  `DisplayName` VARCHAR(50) NOT NULL,
  `HashedPass` VARCHAR(200) NOT NULL,
  `Bio` VARCHAR(160),
  `Location` VARCHAR(30),
  `Website` VARCHAR(100),
  PRIMARY KEY (`UserID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `tweet`;
CREATE TABLE `tweet` (
  `TweetID` int NOT NULL AUTO_INCREMENT,
  `UserID` int NOT NULL,
  `Content` VARCHAR(140) NOT NULL,
  `Reply` int,
  `Quote` int,
  PRIMARY KEY (`TweetID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `fav`;
CREATE TABLE `fav` (
  `LikeID`  int NOT NULL AUTO_INCREMENT, 
  `TweetID` int NOT NULL,
  `UserID`  int NOT NULL,
  PRIMARY KEY (`LikeID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` ( 
  `FollowID` int NOT NULL AUTO_INCREMENT, 
  `FollowerUserID` int NOT NULL,
  `FolloweeUserID` int NOT NULL,
  PRIMARY KEY (`FollowID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `retweet`;
CREATE TABLE `retweet` ( 
  `RetweetID` int NOT NULL AUTO_INCREMENT,
  `UserID` int NOT NULL, 
  `TweetID` int NOT NULL, 
  PRIMARY KEY (`RetweetID`)
) ENGINE = InnoDB;

