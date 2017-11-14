-- Adminer 4.3.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP DATABASE IF EXISTS `hms`;
CREATE DATABASE `hms` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `hms`;

DROP TABLE IF EXISTS `addl_guests`;
CREATE TABLE `addl_guests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `reservation_id` int(11) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `addl_guests`;
INSERT INTO `addl_guests` (`id`, `reservation_id`, `name`) VALUES
(1, 2,  'Kris K. Lee'),
(2, 2,  'Barbara G. Myles');

DROP TABLE IF EXISTS `guests`;
CREATE TABLE `guests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `reservation_id` int(11) NOT NULL,
  `govt_id` varchar(255) NOT NULL,
  `id_number` varchar(255) NOT NULL,
  `check_in` datetime NOT NULL,
  `check_out` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `guests`;
INSERT INTO `guests` (`id`, `reservation_id`, `govt_id`, `id_number`, `check_in`, `check_out`) VALUES
(1, 1,  'driver\'s license',    '3295074775',   '2017-10-08 22:45:00',  '2017-10-10 11:00:00'),
(2, 2,  'driver\'s license',    '3295435431',   '2017-07-08 20:00:00',  '2017-07-14 12:00:00');

DROP TABLE IF EXISTS `hotels`;
CREATE TABLE `hotels` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `description` text,
  `rooms` int(11) NOT NULL,
  `floors` int(11) NOT NULL,
  `logo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `hotels`;
INSERT INTO `hotels` (`id`, `name`, `address`, `description`, `rooms`, `floors`, `logo`) VALUES
(1, 'Hampton Inn',  'Jonesboro AR', 'Features newly renovated rooms.',  45, 4,  'https://pbs.twimg.com/profile_images/378800000867904707/AIy8AP0S.jpeg'),
(2, 'Hilton Garden Inn',    'Paragould AR', NULL,   30, 3,  'http://www.gewekehospitality.com/Hilton-Garden-Inn-Logo-Wallpaper.png'),
(3, 'Embassy Suites',   'Memphis TN',   'Brand new location with 24 hour swimming and fitness.',    100,    10, 'https://upload.wikimedia.org/wikipedia/en/thumb/6/67/Embassy_Suites_Hotels.svg/1280px-Embassy_Suites_Hotels.svg.png'),
(4, 'Hampton Inn',  'St. Louis MO', 'Features newly renovated rooms.',  50, 5,  'https://pbs.twimg.com/profile_images/378800000867904707/AIy8AP0S.jpeg'),
(5, 'Hilton Garden Inn',    'Jonesboro AR', NULL,   63, 6,  'http://www.gewekehospitality.com/Hilton-Garden-Inn-Logo-Wallpaper.png');

DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `confirmation` varchar(255) NOT NULL,
  `reservation_id` int(11) NOT NULL,
  `amount` float NOT NULL,
  `type` enum('card','check','cash') DEFAULT NULL,
  `card_type` enum('none','visa','mastercard','amex','discover','other') DEFAULT NULL,
  `card_last_four` varchar(10) DEFAULT NULL,
  `paid_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `confirmation` (`confirmation`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `payments`;
INSERT INTO `payments` (`id`, `confirmation`, `reservation_id`, `amount`, `type`, `card_type`, `card_last_four`, `paid_on`) VALUES
(1, 'accf8b33', 1,  34560,  'card', 'visa', '4351', '2017-10-10 10:59:30'),
(2, 'accf9d45', 2,  194400, 'card', 'amex', '7354', '2017-07-14 11:59:30');

DROP TABLE IF EXISTS `rates`;
CREATE TABLE `rates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `room_id` int(11) NOT NULL,
  `type` enum('veteran discount','member discount','breakfast bundle','standard') NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `rates`;
INSERT INTO `rates` (`id`, `room_id`, `type`, `price`) VALUES
(1, 1,  'standard', 12500),
(2, 1,  'member discount',  11000),
(3, 1,  'breakfast bundle', 14000),
(4, 2,  'standard', 17500),
(5, 2,  'member discount',  16000),
(6, 2,  'breakfast bundle', 19000),
(7, 3,  'standard', 10000),
(8, 3,  'veteran discount', 8000),
(9, 3,  'breakfast bundle', 12000),
(10,    4,  'standard', 13000),
(11,    4,  'member discount',  11000),
(12,    4,  'veteran discount', 7500),
(13,    5,  'standard', 10000),
(14,    5,  'veteran discount', 8000),
(15,    5,  'breakfast bundle', 11000),
(16,    6,  'standard', 30000),
(17,    6,  'member discount',  25000),
(18,    7,  'standard', 26000),
(19,    7,  'member discount',  23000),
(20,    8,  'standard', 12000),
(21,    8,  'member discount',  11000),
(22,    8,  'breakfast bundle', 13000),
(23,    9,  'standard', 25000),
(24,    9,  'member discount',  22500),
(25,    10, 'standard', 15000),
(26,    10, 'member discount',  12500);

DROP TABLE IF EXISTS `reservations`;
CREATE TABLE `reservations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `confirmation` varchar(255) NOT NULL,
  `user_id` int(11) NOT NULL,
  `rate_id` int(11) NOT NULL,
  `check_in` datetime DEFAULT NULL,
  `check_out` datetime DEFAULT NULL,
  `guests` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `confirmation` (`confirmation`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `reservations`;
INSERT INTO `reservations` (`id`, `confirmation`, `user_id`, `rate_id`, `check_in`, `check_out`, `guests`) VALUES
(1, '77f55803', 1,  5,  '2017-10-08 00:00:00',  '2017-10-10 00:00:00',  1),
(2, '77f55806', 2,  16, '2017-07-08 00:00:00',  '2017-07-15 00:00:00',  3);

DROP TABLE IF EXISTS `rooms`;
CREATE TABLE `rooms` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hotel_id` int(11) NOT NULL,
  `floor` int(11) NOT NULL,
  `num` int(11) NOT NULL,
  `type` enum('suite','single','double') NOT NULL,
  `status` enum('vacant','occupied') NOT NULL,
  `smoking` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `rooms`;
INSERT INTO `rooms` (`id`, `hotel_id`, `floor`, `num`, `type`, `status`, `smoking`) VALUES
(1, 1,  3,  301,    'single',   'occupied', 1),
(2, 1,  2,  201,    'double',   'vacant',   0),
(3, 2,  3,  310,    'double',   'vacant',   0),
(4, 2,  1,  100,    'single',   'occupied', 0),
(5, 3,  5,  505,    'double',   'vacant',   1),
(6, 3,  10, 1010,   'suite',    'vacant',   0),
(7, 4,  5,  505,    'suite',    'occupied', 1),
(8, 4,  2,  203,    'single',   'vacant',   0),
(9, 5,  6,  605,    'suite',    'occupied', 0),
(10,    5,  2,  203,    'single',   'vacant',   0);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `street` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `state` varchar(10) DEFAULT NULL,
  `zip` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

TRUNCATE `users`;
INSERT INTO `users` (`id`, `name`, `email`, `password`, `phone`, `street`, `city`, `state`, `zip`) VALUES
(1, 'Mary M. Elias',    'marymelias@rhyta.com', 'c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a', '816-985-3981', '4267 Loving Acres Road',   'Kansas City',  'MO',   '66210'),
(2, 'Mark V. Shirley',  'markvshirley@jourrapride.com', NULL,   NULL,   NULL,   NULL,   NULL,   NULL);

-- 2017-11-14 03:49:07
