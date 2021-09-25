-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: localhost    Database: db_commerce
-- ------------------------------------------------------
-- Server version	5.5.5-10.4.18-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `commerce_addresses`
--

DROP TABLE IF EXISTS `commerce_addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_addresses` (
  `address_id` int(11) NOT NULL AUTO_INCREMENT,
  `address_street` varchar(100) DEFAULT NULL,
  `address_province` varchar(50) DEFAULT NULL,
  `address_city` varchar(50) DEFAULT NULL,
  `address_country` varchar(50) DEFAULT NULL,
  `address_postcode` int(11) DEFAULT NULL,
  `address_latitude` varchar(100) DEFAULT NULL,
  `address_longitude` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`address_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_addresses`
--

LOCK TABLES `commerce_addresses` WRITE;
/*!40000 ALTER TABLE `commerce_addresses` DISABLE KEYS */;
INSERT INTO `commerce_addresses` VALUES (4,'Jalan Babakan Irigasi Gg. Laksana RT. 02/RW. 05','Jawa Barat','Bandung','Indonesia',40232,'6.212312312','128.02130923','2021-09-25 17:36:31','2021-09-25 17:48:07'),(5,'Jalan Babakan Irigasi Gg. Laksana','Jawa Barat','Bandung','Indonesia',40232,'6.212312312','128.02130923','2021-09-25 17:50:09','2021-09-25 17:50:09');
/*!40000 ALTER TABLE `commerce_addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_cart_products`
--

DROP TABLE IF EXISTS `commerce_cart_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_cart_products` (
  `cart_product_id` float NOT NULL AUTO_INCREMENT,
  `cart_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `product_count` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`cart_product_id`),
  KEY `commerce_cart_product_FK` (`cart_id`),
  KEY `commerce_cart_product_FK_1` (`product_id`),
  CONSTRAINT `commerce_cart_product_FK` FOREIGN KEY (`cart_id`) REFERENCES `commerce_cart_users` (`cart_id`) ON UPDATE CASCADE,
  CONSTRAINT `commerce_cart_product_FK_1` FOREIGN KEY (`product_id`) REFERENCES `commerce_products` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_cart_products`
--

LOCK TABLES `commerce_cart_products` WRITE;
/*!40000 ALTER TABLE `commerce_cart_products` DISABLE KEYS */;
/*!40000 ALTER TABLE `commerce_cart_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_cart_users`
--

DROP TABLE IF EXISTS `commerce_cart_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_cart_users` (
  `cart_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`cart_id`),
  KEY `commerce_cart_user_FK` (`user_id`),
  CONSTRAINT `commerce_cart_user_FK` FOREIGN KEY (`user_id`) REFERENCES `commerce_users` (`user_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_cart_users`
--

LOCK TABLES `commerce_cart_users` WRITE;
/*!40000 ALTER TABLE `commerce_cart_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `commerce_cart_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_categories`
--

DROP TABLE IF EXISTS `commerce_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_categories` (
  `category_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(100) DEFAULT NULL,
  `category_parent_id` int(11) DEFAULT NULL,
  `category_child_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_categories`
--

LOCK TABLES `commerce_categories` WRITE;
/*!40000 ALTER TABLE `commerce_categories` DISABLE KEYS */;
INSERT INTO `commerce_categories` VALUES (1,'Fashion',0,0,'2021-09-25 18:41:46','2021-09-25 18:51:16'),(3,'Pakaian Pria',1,0,'2021-09-25 19:32:34','2021-09-25 19:32:34');
/*!40000 ALTER TABLE `commerce_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_etalases`
--

DROP TABLE IF EXISTS `commerce_etalases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_etalases` (
  `etalase_id` int(11) NOT NULL AUTO_INCREMENT,
  `etalase_name` varchar(20) DEFAULT NULL,
  `shop_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`etalase_id`),
  KEY `commerce_etalase_FK` (`shop_id`),
  CONSTRAINT `commerce_etalase_FK` FOREIGN KEY (`shop_id`) REFERENCES `commerce_shops` (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_etalases`
--

LOCK TABLES `commerce_etalases` WRITE;
/*!40000 ALTER TABLE `commerce_etalases` DISABLE KEYS */;
INSERT INTO `commerce_etalases` VALUES (3,'Etalase Baru',4,'2021-09-25 14:37:44','2021-09-25 14:37:44');
/*!40000 ALTER TABLE `commerce_etalases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_product_categories`
--

DROP TABLE IF EXISTS `commerce_product_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_product_categories` (
  `product_category_id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`product_category_id`),
  KEY `commerce_product_category_FK` (`product_id`),
  KEY `commerce_product_category_FK_1` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_product_categories`
--

LOCK TABLES `commerce_product_categories` WRITE;
/*!40000 ALTER TABLE `commerce_product_categories` DISABLE KEYS */;
INSERT INTO `commerce_product_categories` VALUES (4,1,1,'2021-09-25 19:55:28','2021-09-25 19:55:28'),(5,1,3,'2021-09-25 19:55:51','2021-09-25 19:55:51');
/*!40000 ALTER TABLE `commerce_product_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_products`
--

DROP TABLE IF EXISTS `commerce_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_products` (
  `product_id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(100) DEFAULT NULL,
  `product_pict` text DEFAULT NULL,
  `product_desc` varchar(300) DEFAULT NULL,
  `product_price` int(11) DEFAULT NULL,
  `product_condition` varchar(20) DEFAULT NULL,
  `product_weight` int(11) DEFAULT NULL,
  `product_dimension` varchar(100) DEFAULT NULL,
  `etalase_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`product_id`),
  KEY `commerce_products_FK` (`etalase_id`),
  CONSTRAINT `commerce_products_FK` FOREIGN KEY (`etalase_id`) REFERENCES `commerce_etalases` (`etalase_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_products`
--

LOCK TABLES `commerce_products` WRITE;
/*!40000 ALTER TABLE `commerce_products` DISABLE KEYS */;
INSERT INTO `commerce_products` VALUES (1,'Dompet Kulit','https://google.co.id/img/dompet.jpeg','Dompet kulit khas dari garut',150000,'Baru',150,'8cmx8cmx1cm',3,'2021-09-25 15:37:15','2021-09-25 15:37:15'),(3,'Dompet Kulit Super22','https://google.co.id/img/dompet2.jpeg','Dompet kulit khas dari garut dengan kualitas super',250000,'Baru',120,'7cmx7cmx1cm',3,'2021-09-25 19:51:31','2021-09-25 19:52:08');
/*!40000 ALTER TABLE `commerce_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_roles`
--

DROP TABLE IF EXISTS `commerce_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_roles` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_roles`
--

LOCK TABLES `commerce_roles` WRITE;
/*!40000 ALTER TABLE `commerce_roles` DISABLE KEYS */;
INSERT INTO `commerce_roles` VALUES (1,'admin','2021-09-25 10:39:01','2021-09-25 10:39:01'),(2,'user','2021-09-25 10:39:01','2021-09-25 10:39:01');
/*!40000 ALTER TABLE `commerce_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_shops`
--

DROP TABLE IF EXISTS `commerce_shops`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_shops` (
  `shop_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `shop_name` varchar(30) DEFAULT NULL,
  `shop_desc` varchar(100) DEFAULT NULL,
  `shop_email` varchar(100) DEFAULT NULL,
  `shop_phone_number` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`shop_id`),
  KEY `commerce_shop_FK` (`user_id`),
  CONSTRAINT `commerce_shop_FK` FOREIGN KEY (`user_id`) REFERENCES `commerce_users` (`user_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_shops`
--

LOCK TABLES `commerce_shops` WRITE;
/*!40000 ALTER TABLE `commerce_shops` DISABLE KEYS */;
INSERT INTO `commerce_shops` VALUES (4,14,'Toko Yudi','Adalah toko bersejarah','tokoyudi@gmail.com','08996863574','2021-09-25 14:26:25','2021-09-25 14:26:25');
/*!40000 ALTER TABLE `commerce_shops` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_user_addresses`
--

DROP TABLE IF EXISTS `commerce_user_addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_user_addresses` (
  `user_address_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `address_id` int(11) DEFAULT NULL,
  `address_status` varchar(50) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`user_address_id`),
  KEY `commerce_user_address_FK` (`user_id`),
  KEY `commerce_user_address_FK_1` (`address_id`),
  CONSTRAINT `commerce_user_address_FK` FOREIGN KEY (`user_id`) REFERENCES `commerce_users` (`user_id`) ON UPDATE CASCADE,
  CONSTRAINT `commerce_user_address_FK_1` FOREIGN KEY (`address_id`) REFERENCES `commerce_addresses` (`address_id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_user_addresses`
--

LOCK TABLES `commerce_user_addresses` WRITE;
/*!40000 ALTER TABLE `commerce_user_addresses` DISABLE KEYS */;
INSERT INTO `commerce_user_addresses` VALUES (1,14,4,'Not Primary','2021-09-25 17:36:31','2021-09-25 17:36:31'),(2,14,5,'Not Primary','2021-09-25 17:50:10','2021-09-25 17:50:10');
/*!40000 ALTER TABLE `commerce_user_addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commerce_users`
--

DROP TABLE IF EXISTS `commerce_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `commerce_users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) DEFAULT NULL,
  `password` text DEFAULT NULL,
  `first_name` varchar(25) DEFAULT NULL,
  `last_name` varchar(25) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone_number` varchar(50) DEFAULT NULL,
  `gender` varchar(5) DEFAULT NULL,
  `profile_pic` text DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  `role_id` int(11) DEFAULT 2,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `commerce_users_un` (`username`),
  KEY `commerce_users_FK` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commerce_users`
--

LOCK TABLES `commerce_users` WRITE;
/*!40000 ALTER TABLE `commerce_users` DISABLE KEYS */;
INSERT INTO `commerce_users` VALUES (13,'admin','$2a$10$jQtmfzYdrUO0Uj4a3yKQf.s8I7g0VX3iD9orV4I877dpSfDpaIoBu','Yudi','Widiawan','yudiwidiawan@gmail.com','08996863574','l','https://google.co.id/img/contoh.png','2021-09-25 12:31:20','2021-09-25 12:31:20',1),(14,'yudiwidiawan','$2a$10$37KtZmQscANPsnYNn.JSvuDCdiuXenVoM.Zaatqskl6Ey69zngwMm','Yudi','Widiawan Laki','yudiwidiawa12n@gmail.com','08996863574','l','https://google.co.id/img/contoh.png','2021-09-25 12:35:12','2021-09-25 13:07:49',2),(17,'asepdarwa','$2a$10$HpbwnHNL1eN36RtZ4SXHNOjgc2E9eztM3QsdbgHJPYTBsW.NHE32m','Asep','Darwa','asepdarwa@gmail.com','08888882123','l','https://google.co.id/img/contoh.png','2021-09-25 12:57:26','2021-09-25 12:57:26',2);
/*!40000 ALTER TABLE `commerce_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'db_commerce'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-25 21:19:49
