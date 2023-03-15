DROP DATABASE IF EXISTS `sns`;
CREATE DATABASE `sns` DEFAULT CHARACTER SET utf8mb4;

USE `sns`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` ( 
  `UserID` VARCHAR(30) NOT NULL,
  `Username` VARCHAR(30) NOT NULL,
  `HashedPass` VARCHAR(200) NOT NULL,
  PRIMARY KEY (`UserID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `tweet`;
CREATE TABLE `tweet` (
  `TweetID` int NOT NULL AUTO_INCREMENT,
  `UserID` VARCHAR(30) NOT NULL,
  `Content` VARCHAR(140) NOT NULL,
  `Reply` int,
  `Quote` int,
  PRIMARY KEY (`TweetID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `FavoriteID`  int NOT NULL AUTO_INCREMENT, 
  `TweetID` int NOT NULL,
  `UserID`  VARCHAR(30) NOT NULL,
  PRIMARY KEY (`FavoriteID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` ( 
  `FollowID` int NOT NULL AUTO_INCREMENT, 
  `FollowerUserID` VARCHAR(30) NOT NULL,
  `FolloweeUserID` VARCHAR(30) NOT NULL,
  PRIMARY KEY (`FollowID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `retweet`;
CREATE TABLE `retweet` ( 
  `RetweetID` int NOT NULL AUTO_INCREMENT,
  `UserID` VARCHAR(30) NOT NULL, 
  `TweetID` int NOT NULL, 
  PRIMARY KEY (`RetweetID`)
) ENGINE = InnoDB;

