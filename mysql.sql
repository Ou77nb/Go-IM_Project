/*
SQLyog Ultimate v10.00 Beta1
MySQL - 5.7.9-log : Database - im_project
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`im_project` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `im_project`;

/*Table structure for table `community` */

DROP TABLE IF EXISTS `community`;

CREATE TABLE `community` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `img` longtext,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_community_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Data for the table `community` */

insert  into `community`(`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`owner_id`,`img`,`desc`) values (1,'2023-03-22 23:00:25.733','2023-03-22 23:00:25.733',NULL,'kk',4,'',''),(2,'2023-03-23 13:22:36.167','2023-03-23 13:22:36.167',NULL,'名侦探柯南',2,'./asset/upload/16795489482019727887.jpg','xxx');

/*Table structure for table `contact` */

DROP TABLE IF EXISTS `contact`;

CREATE TABLE `contact` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `target_id` bigint(20) unsigned DEFAULT NULL,
  `type` bigint(20) DEFAULT NULL,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_contact_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

/*Data for the table `contact` */

insert  into `contact`(`id`,`created_at`,`updated_at`,`deleted_at`,`owner_id`,`target_id`,`type`,`desc`) values (1,'2023-03-22 16:34:48.453','2023-03-22 16:34:48.453',NULL,2,4,1,''),(2,'2023-03-22 16:34:48.472','2023-03-22 16:34:48.472',NULL,4,2,1,''),(3,'2023-03-22 19:35:15.519','2023-03-22 19:35:15.519',NULL,5,4,1,''),(4,'2023-03-22 19:35:15.526','2023-03-22 19:35:15.526',NULL,4,5,1,''),(5,'2023-03-22 23:00:25.785','2023-03-22 23:00:25.785',NULL,4,1,2,''),(6,'2023-03-22 23:28:43.502','2023-03-22 23:28:43.502',NULL,2,1,2,''),(7,'2023-03-22 23:29:00.321','2023-03-22 23:29:00.321',NULL,5,1,2,''),(8,'2023-03-23 13:22:36.185','2023-03-23 13:22:36.185',NULL,2,2,2,'');

/*Table structure for table `group_basic` */

DROP TABLE IF EXISTS `group_basic`;

CREATE TABLE `group_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `icon` longtext,
  `type` bigint(20) DEFAULT NULL,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_group_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `group_basic` */

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `target_id` bigint(20) DEFAULT NULL,
  `type` bigint(20) DEFAULT NULL,
  `media` bigint(20) DEFAULT NULL,
  `content` longtext,
  `create_time` bigint(20) unsigned DEFAULT NULL,
  `read_time` bigint(20) unsigned DEFAULT NULL,
  `pic` longtext,
  `url` longtext,
  `desc` longtext,
  `amount` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_message_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `message` */

/*Table structure for table `user_basics` */

DROP TABLE IF EXISTS `user_basics`;

CREATE TABLE `user_basics` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `pass_word` longtext,
  `phone` longtext,
  `email` longtext,
  `avatar` longtext,
  `identity` longtext,
  `client_ip` longtext,
  `client_port` longtext,
  `salt` longtext,
  `login_time` datetime(3) DEFAULT NULL,
  `heartbeat_time` datetime(3) DEFAULT NULL,
  `login_out_time` datetime(3) DEFAULT NULL,
  `is_logout` tinyint(1) DEFAULT NULL,
  `device_info` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_user_basics_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `user_basics` */

insert  into `user_basics`(`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`pass_word`,`phone`,`email`,`avatar`,`identity`,`client_ip`,`client_port`,`salt`,`login_time`,`heartbeat_time`,`login_out_time`,`is_logout`,`device_info`) values (2,'2023-03-21 22:31:46.659','2023-03-23 15:08:06.692',NULL,'guoguo','b6a01b1423bfb10ace8be6b0d7a31f24','123','','./asset/upload/16795489221298498081.jpg','0650b243185a32c6d33c613101320b8d','','','1298498081','2023-03-21 22:31:46.658','2023-03-21 22:31:46.658','2023-03-21 22:31:46.658',0,''),(3,'2023-03-21 22:38:46.504','2023-03-21 22:38:46.504','2023-03-21 23:07:11.017','guoguo2','1e945f15004d3b2c17a4e2862a3edfbf','','','','','','','2019727887','2023-03-21 22:38:46.504','2023-03-21 22:38:46.504','2023-03-21 22:38:46.504',0,''),(4,'2023-03-22 14:33:37.256','2023-03-23 15:09:27.303',NULL,'guoguo3','b6a01b1423bfb10ace8be6b0d7a31f24','','','','50d96385734bb05de9bd2b3a43276a5e','','','1298498081','2023-03-22 14:33:37.256','2023-03-22 14:33:37.256','2023-03-22 14:33:37.256',0,''),(5,'2023-03-22 19:34:39.997','2023-03-22 19:34:48.456',NULL,'guoguo4','b6a01b1423bfb10ace8be6b0d7a31f24','','','','962c1b908177e9fc06c42dc999e7298b','','','1298498081','2023-03-22 19:34:39.997','2023-03-22 19:34:39.997','2023-03-22 19:34:39.997',0,'');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
