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