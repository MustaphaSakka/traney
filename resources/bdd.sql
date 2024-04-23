CREATE DATABASE traney;
USE traney;

DROP TABLE IF EXISTS `clients`;
CREATE TABLE `clients` (
  `client_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(100) NOT NULL,
  `zipcode` varchar(10) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
INSERT INTO `clients` VALUES
	(1,'Mustapha','1990-01-02','Lyon','69003',1),
	(2,'Pierre','1998-11-21','Villeurbane','69100',1),
  (3,'Isabelle','1987-01-14','Paris','92000',0),
	(4,'Samuel','1985-04-20','Nice','06000',1);


  
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `account_id` int(11) NOT NULL AUTO_INCREMENT,
  `client_id` int(11) NOT NULL,
  `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `account_type` varchar(10) NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`account_id`),
  KEY `accounts_FK` (`client_id`),
  CONSTRAINT `accounts_FK` FOREIGN KEY (`client_id`) REFERENCES `clients` (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;
INSERT INTO `accounts` VALUES
	(1,1,'2024-03-22 10:20:06', 'epargne', 5555.55, 1),
	(2,2,'2024-03-09 10:27:22', 'courant', 4444.44, 1),
  (3,3,'2023-08-09 10:35:22', 'epargne', 3000, 1),
  (4,3,'2023-08-09 10:38:22', 'courant', 2222.22, 1);


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `client_id` int(11) DEFAULT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `users` VALUES
  ('admin','abc123','admin', NULL, '2021-08-09 11:17:22'),
  ('1','abc123','user', 1, '2020-08-09 10:27:22'),
  ('2','abc123','user', 2, '2021-08-09 12:27:22');