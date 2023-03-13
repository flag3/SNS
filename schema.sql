DROP DATABASE IF EXISTS `sns`;
CREATE DATABASE `sns` DEFAULT CHARACTER SET utf8mb4;

USE `sns`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` ( 
  `UserID` VARCHAR(30) NOT NULL,
  `Username` VARCHAR(30) NOT NULL , 
  `HashedPass` VARCHAR(200) NOT NULL , 
  PRIMARY KEY (`UserID`)
) ENGINE = InnoDB;

DROP TABLE IF EXISTS `tweet`;
CREATE TABLE `tweet` (
  `TweetID` int NOT NULL AUTO_INCREMENT,
  `UserID`  VARCHAR(30) NOT NULL,
  `Body`    VARCHAR(140) NOT NULL,
  PRIMARY KEY (`TweetID`)
) ENGINE = InnoDB;

INSERT INTO `tweet` VALUES (1,'sobaya007','は？俺は全強');
INSERT INTO `tweet` VALUES (2,'sobaya007','は？俺はホモ');

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `FavoriteID`  int NOT NULL AUTO_INCREMENT, 
  `TweetID` int NOT NULL,
  `UserID`  VARCHAR(30) NOT NULL,
  PRIMARY KEY (`FavoriteID`)
) ENGINE = InnoDB;

INSERT INTO `favorite` VALUES (1,8,'sobaya007');
INSERT INTO `favorite` VALUES (2,1,'sobaya007');

DROP TABLE IF EXISTS `follow`
CREATE TABLE `follow` ( `FollowID` int NOT NULL AUTO_INCREMENT, 
  `FollowerUserID` VARCHAR(30) NOT NULL,
  `FolloweeUserID` VARCHAR(30) NOT NULL,
  PRIMARY KEY (`FollowID`)
) ENGINE = InnoDB;

INSERT INTO `follow` VALUES (1,'_ymdtr_', 'sobaya007');
