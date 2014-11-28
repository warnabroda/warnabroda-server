-- MySQL dump 10.13  Distrib 5.5.40, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: martini
-- ------------------------------------------------------
-- Server version	5.5.40-0ubuntu0.14.04.1

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
-- Table structure for table `contact_types`
--

DROP TABLE IF EXISTS `contact_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contact_types` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Lang_key` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contact_types`
--

LOCK TABLES `contact_types` WRITE;
/*!40000 ALTER TABLE `contact_types` DISABLE KEYS */;
INSERT INTO `contact_types` VALUES (1,'E-mail','br'),(2,'SMS','br'),(3,'Facebook','br');
/*!40000 ALTER TABLE `contact_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `messages`
--

DROP TABLE IF EXISTS `messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `messages` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Lang_key` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (1,'Está com Mal Hálito','br'),(2,'Está com cheio desagradável de suor','br'),(3,'Tem Sujeira nos dentes','br'),(4,'Tem Sinal de menstruação na roupa','br'),(5,'Tem Sugeira de merda no vaso de casa','br'),(6,'Está sendo traido(a)','br'),(7,'Está Fazendo barulho incomodo com a boca','br'),(8,'Está Fazendo barulho incomodo com pés ou mãos','br'),(9,'Está com chulé','br');
/*!40000 ALTER TABLE `messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subjects`
--

DROP TABLE IF EXISTS `subjects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subjects` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Lang_key` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjects`
--

LOCK TABLES `subjects` WRITE;
/*!40000 ALTER TABLE `subjects` DISABLE KEYS */;
INSERT INTO `subjects` VALUES (1,'Pô parceiro(a)','br'),(2,'Um amigo(a) pediu para avisar','br'),(3,'Só um toque','br'),(4,'Ow, se liga ae','br'),(5,'Ola, venho atraves desse informar','br'),(6,'É para o teu bem','br'),(7,'Quem avisa amigo é','br'),(8,'Não é por mal é só um aviso','br'),(9,'Só tô avisando porque te acho legal','br'),(10,'Só to te avisando porque me importo contigo','br'),(11,'Um amigo acaba de lhe dar um toque','br');
/*!40000 ALTER TABLE `subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warnings`
--

DROP TABLE IF EXISTS `warnings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warnings` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Id_message` bigint(20) DEFAULT NULL,
  `Id_contact_type` bigint(20) DEFAULT NULL,
  `Contact` varchar(255) DEFAULT NULL,
  `Sent` tinyint(1) DEFAULT NULL,
  `Message` varchar(255) DEFAULT NULL,
  `Ip` varchar(255) DEFAULT NULL,
  `Browser` varchar(255) DEFAULT NULL,
  `Operating_system` varchar(255) DEFAULT NULL,
  `Device` varchar(255) DEFAULT NULL,
  `Raw` varchar(255) DEFAULT NULL,
  `Created_by` varchar(255) DEFAULT NULL,
  `Created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Last_modified_by` varchar(255) DEFAULT NULL,
  `Last_modified_date` timestamp NULL DEFAULT NULL,
  `Lang_key` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warnings`
--

LOCK TABLES `warnings` WRITE;
/*!40000 ALTER TABLE `warnings` DISABLE KEYS */;
INSERT INTO `warnings` VALUES (1,0,1,'',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00 00:00:00','','0000-00-00 00:00:00',''),(2,1,1,'teste',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00 00:00:00','','0000-00-00 00:00:00',''),(3,6,3,'hbtsmith',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00 00:00:00','','0000-00-00 00:00:00',''),(4,1,1,'dalkjsdlsajdlkajs',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00',''),(5,9,2,'4896662015',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00',''),(6,1,3,'testeste',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(7,1,1,'teaetafsaf',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(8,2,3,'dasdasdsadas',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(9,1,1,'123456788',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(10,1,1,'poiuy',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(11,1,2,'dasdasdsadas',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 20:47:22','','0000-00-00 00:00:00','br'),(12,1,2,'mnbcamnnb',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 20:53:28','','0000-00-00 00:00:00','br'),(13,7,1,'15978963',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','0000-00-00 00:00:00','','0000-00-00 00:00:00','br'),(14,9,3,'09876543',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 21:38:40','','0000-00-00 00:00:00','br'),(15,1,1,'hbt.vieira@gmail.com',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 21:45:08','','0000-00-00 00:00:00','br'),(16,1,1,'1',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 21:46:12','','0000-00-00 00:00:00','br'),(17,1,1,'aaaaaaaaaa',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 22:09:15','','0000-00-00 00:00:00','br'),(18,9,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 22:16:25','','0000-00-00 00:00:00','br'),(19,1,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 22:33:48','','0000-00-00 00:00:00','br'),(20,3,1,'herbert.silva@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.65 Safari/537.36','system','2014-11-26 22:35:35','','2014-11-26 22:35:36','br'),(21,1,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:22:52','','0000-00-00 00:00:00','br'),(22,1,1,'TESTE1',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:24:17','','2014-11-27 15:24:18','br'),(23,1,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:29:10','','0000-00-00 00:00:00','br'),(24,1,1,'teste2',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:29:43','','0000-00-00 00:00:00','br'),(25,1,1,'teste2',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:30:03','','2014-11-27 15:30:04','br'),(26,1,1,'teste3',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:31:57','','2014-11-27 15:31:57','br'),(27,1,1,'teste4',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:32:55','','0000-00-00 00:00:00','br'),(28,1,1,'teste4',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:34:52','','2014-11-27 15:34:53','br'),(29,1,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:37:42','','0000-00-00 00:00:00','br'),(30,1,1,'hbt.vieira@gmail.com',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:40:40','','0000-00-00 00:00:00','br'),(31,1,1,'hbt.vieira@gmail.com',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:41:19','','2014-11-27 15:41:20','br'),(32,2,1,'teste4',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:43:41','','2014-11-27 15:43:42','br'),(33,3,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:52:29','','0000-00-00 00:00:00','br'),(34,4,1,'herbert.silva@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:53:25','','2014-11-27 15:53:26','br'),(35,5,1,'herbert.silva@jexperts.com.br',0,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:55:22','','0000-00-00 00:00:00','br'),(36,1,1,'herbert.silva@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:55:58','','2014-11-27 15:55:59','br'),(37,8,1,'hbt.vieira@gmail.com',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 15:56:50','','2014-11-27 15:56:50','br'),(38,9,1,'hbt.vieira@gmail.com',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-11-27 18:36:10','','2014-11-27 18:36:11','br');
/*!40000 ALTER TABLE `warnings` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-11-28  9:48:16
