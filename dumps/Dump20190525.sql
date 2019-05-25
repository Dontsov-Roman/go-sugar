-- MySQL dump 10.13  Distrib 5.7.26, for Linux (x86_64)
--
-- Host: localhost    Database: sugar
-- ------------------------------------------------------
-- Server version	5.7.26-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `auth_session`
--

DROP TABLE IF EXISTS `auth_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auth_session` (
  `user_id` int(10) unsigned NOT NULL,
  `device_id` varchar(45) DEFAULT NULL,
  `token` varchar(512) NOT NULL COMMENT '\n',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`token`),
  KEY `deviceID` (`device_id`),
  KEY `token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_session`
--

LOCK TABLES `auth_session` WRITE;
/*!40000 ALTER TABLE `auth_session` DISABLE KEYS */;
INSERT INTO `auth_session` VALUES (59,'','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjo1OSwiTmFtZSI6IlRmZ2p1cmQiLCJQYXNzd29yZCI6ImVhNzFmOGIxYjYxNzFmMGEyNDMzNDJiNTAwMmU2NGQxYjM1ZmZiOTMiLCJUeXBlIjowLCJTdGF0dXMiOjEsIkVtYWlsIjoiRGZoamJjQGNibm0uY2JoIiwiUGhvbmUiOiI1NTg4ODUiLCJDcmVhdGVkQXQiOm51bGwsIlVwZGF0ZWRBdCI6bnVsbCwiRGVsZXRlZEF0IjpudWxsfSwiZXhwIjoxNTU4ODA5NTg4fQ.r4r7dbhym_ZOLxGkW-rncA8R9hifocKlKafgWyImIzU','2019-05-25 20:39:48','2019-05-25 20:39:48'),(60,'','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjo2MCwiTmFtZSI6IlJ0Y3p6IiwiUGFzc3dvcmQiOiJlYTcxZjhiMWI2MTcxZjBhMjQzMzQyYjUwMDJlNjRkMWIzNWZmYjkzIiwiVHlwZSI6MCwiU3RhdHVzIjoxLCJFbWFpbCI6IkR2aHR0Z0B6Y2IuY3ZiIiwiUGhvbmUiOiIyMjU1ODgiLCJDcmVhdGVkQXQiOm51bGwsIlVwZGF0ZWRBdCI6bnVsbCwiRGVsZXRlZEF0IjpudWxsfSwiZXhwIjoxNTU4ODA5NzAwfQ.HRg2nEQqegM1Qd1YSoG6F0uz0MBEftT1NgMNtvFABrY','2019-05-25 20:41:40','2019-05-25 20:41:40'),(61,'','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjo2MSwiTmFtZSI6IlF3ZXJ0eSIsIlBhc3N3b3JkIjoiZWE3MWY4YjFiNjE3MWYwYTI0MzM0MmI1MDAyZTY0ZDFiMzVmZmI5MyIsIlR5cGUiOjAsIlN0YXR1cyI6MSwiRW1haWwiOiJGdm5oeWdAbWFpbC5jb20iLCJQaG9uZSI6IjI1ODAwIiwiQ3JlYXRlZEF0IjpudWxsLCJVcGRhdGVkQXQiOm51bGwsIkRlbGV0ZWRBdCI6bnVsbH0sImV4cCI6MTU1ODgxMDE2NH0.XRZ6OQjC3yeuqwnjw0k2TJngNq06YkFdI7ccw0jwKB8','2019-05-25 20:49:24','2019-05-25 20:49:24');
/*!40000 ALTER TABLE `auth_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `description` longtext COLLATE utf8_bin,
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int(10) unsigned DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders_prices`
--

DROP TABLE IF EXISTS `orders_prices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders_prices` (
  `order_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `price_id` int(10) unsigned NOT NULL,
  UNIQUE KEY `order_price` (`order_id`,`price_id`) USING BTREE,
  KEY `price_id_idx` (`price_id`),
  KEY `user_id_idx` (`user_id`),
  CONSTRAINT `order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `price_id` FOREIGN KEY (`price_id`) REFERENCES `prices` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders_prices`
--

LOCK TABLES `orders_prices` WRITE;
/*!40000 ALTER TABLE `orders_prices` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders_prices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prices`
--

DROP TABLE IF EXISTS `prices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `prices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) COLLATE utf8_bin NOT NULL,
  `status` int(11) DEFAULT '1',
  `price` int(11) NOT NULL,
  `time` int(11) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prices`
--

LOCK TABLES `prices` WRITE;
/*!40000 ALTER TABLE `prices` DISABLE KEYS */;
INSERT INTO `prices` VALUES (2,'Бедра',3,170,45,'2019-02-14 23:00:03','2019-02-14 23:00:03',NULL),(3,'Глубокое бикини',1,190,30,'2019-02-14 23:00:29','2019-02-14 23:00:29',NULL),(4,'Бикини',1,150,30,'2019-02-14 23:00:46','2019-02-14 23:00:46',NULL),(5,'Ножки полностью',1,300,60,'2019-02-14 23:14:32','2019-02-14 23:19:34',NULL),(6,'усики',1,50,15,'2019-05-19 18:21:54','2019-05-19 18:21:54',NULL),(7,'подмышки',1,100,20,'2019-05-19 18:22:15','2019-05-19 18:22:15',NULL),(8,'голени',1,100,30,'2019-05-19 18:22:28','2019-05-19 18:22:28',NULL),(9,'животик',1,150,30,'2019-05-19 18:23:42','2019-05-19 18:23:42',NULL),(10,'поясница',1,170,40,'2019-05-19 18:23:53','2019-05-19 18:23:53',NULL),(11,'ручки до локтя',1,120,25,'2019-05-19 18:24:07','2019-05-19 18:24:07',NULL),(12,'ручки полностью',1,170,40,'2019-05-19 18:24:36','2019-05-19 18:24:36',NULL);
/*!40000 ALTER TABLE `prices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL DEFAULT '1',
  `email` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `phone` varchar(25) COLLATE utf8_bin NOT NULL,
  `name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `deleted_at` datetime DEFAULT NULL,
  `password` varchar(65) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `phone_UNIQUE` (`phone`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (56,0,'Ffhhrd@dfhj.com','0974885047','Rfghgrd','2019-05-25 20:34:06','2019-05-25 20:34:06',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93'),(57,0,'Riuggf@mail.fcx','500485','Fghtdfg','2019-05-25 20:37:03','2019-05-25 20:37:03',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93'),(58,0,'Qwerrt@cvnn.vcc','5865588','Rtgde','2019-05-25 20:37:59','2019-05-25 20:37:59',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93'),(59,0,'Dfhjbc@cbnm.cbh','558885','Tfgjurd','2019-05-25 20:39:48','2019-05-25 20:39:48',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93'),(60,0,'Dvhttg@zcb.cvb','225588','Rtczz','2019-05-25 20:41:40','2019-05-25 20:41:40',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93'),(61,0,'Fvnhyg@mail.com','25800','Qwerty','2019-05-25 20:49:24','2019-05-25 20:49:24',1,NULL,'ea71f8b1b6171f0a243342b5002e64d1b35ffb93');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-25 20:53:24
