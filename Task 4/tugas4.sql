/*
SQLyog Community v13.1.6 (64 bit)
MySQL - 10.1.36-MariaDB : Database - tugas4
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`tugas4` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `tugas4`;

/*Table structure for table `modification` */

DROP TABLE IF EXISTS `modification`;

CREATE TABLE `modification` (
  `Date` varchar(30) DEFAULT NULL,
  `DateGmt` varchar(30) DEFAULT NULL,
  `RenderedGUID` varchar(30) DEFAULT NULL,
  `ID` int(10) NOT NULL,
  `Link` varchar(40) DEFAULT NULL,
  `Modified` varchar(30) DEFAULT NULL,
  `ModifiedGmt` varchar(30) DEFAULT NULL,
  `Slug` varchar(30) DEFAULT NULL,
  `RenderedTitle` varchar(50) DEFAULT NULL,
  `Type` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `modification` */

insert  into `modification`(`Date`,`DateGmt`,`RenderedGUID`,`ID`,`Link`,`Modified`,`ModifiedGmt`,`Slug`,`RenderedTitle`,`Type`) values 
('','','',0,'','','','','',''),
('2017-07-21T10:30:34','2017-07-21T17:30:34','https://www.sitepoint.com/?p=1',143,'https://www.sitepoint.com/why-the-iot-th','2017-07-23T21:56:35','2017-07-24T04:56:35','why-the-iot-threatens-your-wor','Why the IoT Threatens Your WordPress Site (and How','post');

/*Table structure for table `thumbnail` */

DROP TABLE IF EXISTS `thumbnail`;

CREATE TABLE `thumbnail` (
  `Color` varchar(20) NOT NULL,
  `Category` varchar(30) DEFAULT NULL,
  `Hex` varchar(30) DEFAULT NULL,
  `Rgba` varchar(30) DEFAULT NULL,
  `Type` varchar(20) DEFAULT NULL,
  `URL` varchar(50) DEFAULT NULL,
  `Width` int(10) DEFAULT NULL,
  `Height` int(10) DEFAULT NULL,
  PRIMARY KEY (`Color`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `thumbnail` */

insert  into `thumbnail`(`Color`,`Category`,`Hex`,`Rgba`,`Type`,`URL`,`Width`,`Height`) values 
('','','','','','0',0,0),
('hue','#000','255,255,255,1','black','primary','32',0,32);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
