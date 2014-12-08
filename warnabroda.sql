-- MySQL dump 10.13  Distrib 5.5.40, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: warnabroda
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
INSERT INTO `contact_types` VALUES (1,'E-mail','br'),(2,'SMS','br');
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (1,'Está com Mal Hálito','br'),(2,'Está com odor de suor','br'),(3,'Tem Sujeira nos dentes','br'),(4,'Tem Sinal de menstruação na roupa','br'),(5,'Tem marca de coco no vaso de casa','br'),(7,'Está Fazendo barulho incomodo com a boca','br'),(8,'Está Fazendo barulho incomodo com pés ou mãos','br'),(9,'Está com chulé','br'),(10,'Está com a roupa do lado avesso','br'),(11,'Está com o cofrinho aparecendo','br');
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
) ENGINE=InnoDB AUTO_INCREMENT=170 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warnings`
--

LOCK TABLES `warnings` WRITE;
/*!40000 ALTER TABLE `warnings` DISABLE KEYS */;
INSERT INTO `warnings` VALUES (4,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:04:27','','0000-00-00 00:00:00','br'),(5,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:19:28','','0000-00-00 00:00:00','br'),(6,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:20:17','','0000-00-00 00:00:00','br'),(7,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:24:23','','0000-00-00 00:00:00','br'),(8,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:24:42','','0000-00-00 00:00:00','br'),(9,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:57:21','','0000-00-00 00:00:00','br'),(10,2,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:59:25','','2014-11-28 00:59:26','br'),(11,5,1,'hbt.vieira@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 01:02:02','','2014-11-28 01:02:03','br'),(12,1,1,'herbert.silva@jexperts.com.br',0,'','177.16.146.153','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-02 01:37:36','','0000-00-00 00:00:00','br'),(13,3,1,'Adrieli.daltoe@gmail.com',0,'','179.207.178.19','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; GT-I9300 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-02 17:34:01','','0000-00-00 00:00:00','br'),(14,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:08:52','','0000-00-00 00:00:00','br'),(15,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:01','','0000-00-00 00:00:00','br'),(16,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:03','','0000-00-00 00:00:00','br'),(17,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:17','','0000-00-00 00:00:00','br'),(18,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:11:27','','0000-00-00 00:00:00','br'),(19,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:12:52','','0000-00-00 00:00:00','br'),(20,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:11','','2014-12-04 22:14:12','br'),(21,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:13','','2014-12-04 22:14:13','br'),(22,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:22','','2014-12-04 22:14:22','br'),(23,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:22','','2014-12-04 22:14:23','br'),(28,4,2,'4891917810',1,'6;36073798','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:28:25','','2014-12-04 22:28:26','br'),(29,6,2,'4896662015',1,'6;36084625','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:13:40','','2014-12-05 12:13:40','br'),(30,6,2,'4891640107',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:14:14','','0000-00-00 00:00:00','br'),(31,6,1,'tjvargas2@gmail.com',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:14:45','','2014-12-05 12:14:46','br'),(32,6,2,'4891188080',1,'6;36084637','186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.104 Safari/537.36','system','2014-12-05 12:14:58','','2014-12-05 12:14:58','br'),(33,4,1,'vinifritzen@gmail.com',1,'','192.168.1.81, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:17:15','','2014-12-05 12:17:16','br'),(34,8,2,'4896004128',1,'6;36084674','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:17:53','','2014-12-05 12:17:54','br'),(35,1,1,'ramoncordini@gmail.com',1,'','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:19:28','','2014-12-05 12:19:28','br'),(36,9,1,'graziela.souza@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:20:00','','2014-12-05 12:20:01','br'),(37,3,1,'fasebastiani@gmail.com',1,'','189.90.61.78','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:22:54','','2014-12-05 12:22:54','br'),(38,9,1,'ademarizu@gmail.com',1,'','189.101.247.92','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:24:42','','2014-12-05 12:24:43','br'),(39,6,2,'4884111635',1,'6;36084834','189.101.247.92','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:25:42','','2014-12-05 12:25:42','br'),(40,1,1,'phelipewinter@gmail.com',1,'','186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.104 Safari/537.36','system','2014-12-05 12:27:45','','2014-12-05 12:27:46','br'),(41,3,1,'mayara.souza@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:28:00','','2014-12-05 12:28:00','br'),(42,8,1,'mayara@jexperts.com.bt',1,'','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:28:05','','2014-12-05 12:28:06','br'),(43,8,1,'mayara.souza@jexperts.com.br',1,'','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:29:25','','2014-12-05 12:29:26','br'),(44,2,1,'leonardo@ahgora.com.br',1,'','189.101.215.166','chrome','android','android','Mozilla/5.0 (Linux; Android 4.1.2; XT920 Build/2_330_2009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-05 12:30:40','','2014-12-05 12:30:41','br'),(45,3,2,'4888634777',0,'','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:30:45','','0000-00-00 00:00:00','br'),(46,7,2,'4799698083',1,'6;36085135','60.225.193.247','chrome','android','android','Mozilla/5.0 (Linux; Android 5.0; Nexus 5 Build/LRX21O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-05 12:31:13','','2014-12-05 12:31:13','br'),(47,9,1,'pauladesouza1604@gmail.com',1,'','192.168.1.113, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:31:44','','2014-12-05 12:31:45','br'),(48,8,1,'fernanda.barbosa@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:33:50','','2014-12-05 12:33:51','br'),(49,8,1,'kamilla.pires@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:34:06','','2014-12-05 12:34:07','br'),(50,9,1,'ricardojp84@gmail.com',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:34:28','','2014-12-05 12:34:28','br'),(51,10,2,'4891651919',1,'6;36085322','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:37:20','','2014-12-05 12:37:21','br'),(52,1,1,'daniela.munari@fiesc.com.br',1,'','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:40:49','','2014-12-05 12:40:49','br'),(53,10,1,'sec.rosimeri@gmail.com',1,'','192.168.1.113, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:41:25','','2014-12-05 12:41:25','br'),(54,10,1,'jabez_t.i@hotmail.com',1,'','177.207.221.70','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64; rv:33.0) Gecko/20100101 Firefox/33.0','system','2014-12-05 12:43:34','','2014-12-05 12:43:35','br'),(55,10,1,'daniel.cherem@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:44:00','','2014-12-05 12:44:00','br'),(56,9,2,'4899312065',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:04','','0000-00-00 00:00:00','br'),(57,3,1,'adrieli.daltoe@gmail.com',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:13','','2014-12-05 12:44:14','br'),(58,9,2,'4896852019',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:29','','0000-00-00 00:00:00','br'),(59,9,2,'4896852019',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:35','','0000-00-00 00:00:00','br'),(60,10,1,'jabez_t.i@hotmail.com',1,'','177.207.221.70','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64; rv:33.0) Gecko/20100101 Firefox/33.0','system','2014-12-05 12:45:44','','2014-12-05 12:45:44','br'),(61,9,1,'lecogoulart@bol.com.br',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 12:45:54','','2014-12-05 12:45:54','br'),(62,10,2,'4891640107',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:46:09','','0000-00-00 00:00:00','br'),(63,3,1,'rafael.nishihora@gmail.com',1,'','150.162.81.101','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:58:30','','2014-12-05 12:58:30','br'),(64,3,2,'4899211951',1,'6;36087095','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:32','','2014-12-05 12:59:32','br'),(65,3,2,'4891352028',1,'6;36087099','150.162.81.101','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:35','','2014-12-05 12:59:36','br'),(66,2,2,'4899895468',0,'','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:46','','0000-00-00 00:00:00','br'),(67,9,1,'Contepraju@gmail',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 13:22:33','','2014-12-05 13:22:33','br'),(68,4,1,'contepraju@gmail.com',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 13:24:45','','2014-12-05 13:24:46','br'),(69,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:23','','2014-12-05 13:26:23','br'),(70,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:23','','2014-12-05 13:26:24','br'),(71,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:43','','2014-12-05 13:26:43','br'),(72,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:43','','2014-12-05 13:26:44','br'),(73,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:27:09','','2014-12-05 13:27:09','br'),(74,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:27:09','','2014-12-05 13:27:10','br'),(103,9,1,'altemir@ahgora.com.br',1,'','189.90.61.78','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 19:17:58','','2014-12-05 19:17:58','br'),(104,3,2,'4896270570',1,'6;36186883','189.8.253.109','chrome','linux','unknown','Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 19:19:04','','2014-12-05 19:19:05','br'),(105,7,1,'artcospresentes@live.com',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 19:52:04','','2014-12-05 19:52:05','br'),(106,9,1,'andre.guilherme@live.com',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 19:54:05','','2014-12-05 19:54:06','br'),(107,3,1,'mdiluccio@gmail.com',1,'','150.162.81.59','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:24:43','','2014-12-05 20:24:44','br'),(108,4,1,'javgisele@gmail.com',1,'','150.162.81.59','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:26:02','','2014-12-05 20:26:02','br'),(109,1,1,'nelsonclage@gmail.com',1,'','150.162.81.71','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:26:12','','2014-12-05 20:26:12','br'),(110,1,2,'4896662015',0,'','186.215.116.36','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:27:51','','0000-00-00 00:00:00','br'),(111,9,1,'dirceu.scaratti@gmail.com',1,'','150.162.81.101','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:28:05','','2014-12-05 20:28:06','br'),(112,4,1,'tfurtado@gmail.com',1,'','150.162.81.138','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.104 Safari/537.36','system','2014-12-05 20:32:54','','2014-12-05 20:32:55','br'),(113,9,1,'pedro.vitti@resultadosdigitais.com.br',1,'','187.94.99.90','firefox','linux','unknown','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 20:44:10','','2014-12-05 20:44:10','br'),(114,9,1,'julianapavei@gmail.com',1,'','150.162.81.71','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 20:55:46','','2014-12-05 20:55:46','br'),(115,3,2,'4888246086',1,'6;36233460','189.90.61.78','chrome','linux','unknown','Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36','system','2014-12-05 21:28:02','','2014-12-05 21:28:02','br'),(116,3,2,'4896672721',1,'6;36233908','150.162.148.71','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:45:13','','2014-12-05 21:45:13','br'),(117,3,2,'6699860087',1,'6;36233933','150.162.198.233','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:48:29','','2014-12-05 21:48:29','br'),(118,1,2,'4891762742',0,'','150.162.198.233','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:49:29','','0000-00-00 00:00:00','br'),(119,8,2,'4891762742',0,'','150.162.148.71','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:50:05','','0000-00-00 00:00:00','br'),(120,1,2,'4896044124',0,'','150.162.198.233','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:50:29','','0000-00-00 00:00:00','br'),(121,1,2,'4896044124',0,'','150.162.198.233','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 21:50:36','','0000-00-00 00:00:00','br'),(122,3,1,'barbara@dataa.com.br',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 22:01:15','','2014-12-05 22:01:15','br'),(123,3,1,'gabriel@dataa.com.br',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 22:05:10','','2014-12-05 22:05:11','br'),(124,9,1,'jeffetorquato@gmail.com',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 22:07:44','','2014-12-05 22:07:44','br'),(125,10,2,'4896296178',0,'','189.90.61.78','chrome','linux','unknown','Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36','system','2014-12-05 22:40:31','','0000-00-00 00:00:00','br'),(126,9,1,'tjvargas2@gmail.com',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 22:42:54','','2014-12-05 22:42:55','br'),(127,10,1,'tjvargas2@gmail.com',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 22:47:22','','2014-12-05 22:47:22','br'),(128,2,1,'ademarizu@gmail.com',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 00:00:50','','2014-12-06 00:00:51','br'),(129,8,2,'4891533880',1,'6;36250367','177.97.190.215','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 02:57:31','','2014-12-06 02:57:31','br'),(130,8,2,'4899315933',0,'','177.97.190.215','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 02:59:50','','0000-00-00 00:00:00','br'),(131,5,2,'4899315933',0,'','177.42.51.251','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:04:06','','0000-00-00 00:00:00','br'),(132,5,2,'4899315933',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:07:59','','0000-00-00 00:00:00','br'),(133,5,2,'4891651919',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:09:25','','0000-00-00 00:00:00','br'),(134,5,2,'4899973992',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:10:12','','0000-00-00 00:00:00','br'),(135,5,2,'4884167097',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:10:53','','0000-00-00 00:00:00','br'),(136,5,2,'4899513441',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:11:32','','0000-00-00 00:00:00','br'),(137,5,2,'4888234804',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:12:12','','0000-00-00 00:00:00','br'),(138,5,2,'4888234804',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:12:14','','0000-00-00 00:00:00','br'),(139,5,2,'4896662015',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:12:52','','0000-00-00 00:00:00','br'),(140,5,2,'4896662015',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:58:59','','0000-00-00 00:00:00','br'),(141,5,2,'4899999999',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 03:59:07','','0000-00-00 00:00:00','br'),(142,5,2,'4899973992',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 07:03:01','','0000-00-00 00:00:00','br'),(143,5,2,'4891651919',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 07:03:22','','0000-00-00 00:00:00','br'),(144,5,2,'4899513441',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 07:45:24','','0000-00-00 00:00:00','br'),(145,5,2,'4896662015',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 07:45:36','','0000-00-00 00:00:00','br'),(146,5,2,'4899973992',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 08:53:38','','0000-00-00 00:00:00','br'),(147,5,2,'4899315933',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 11:00:05','','0000-00-00 00:00:00','br'),(148,1,1,'hbt.vieira@gmail.com',1,'','177.41.241.92','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-06 17:47:22','','2014-12-06 17:47:23','br'),(149,9,2,'4896837097',1,'6;36293496','189.4.83.123','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 03:52:22','','2014-12-07 03:52:23','br'),(150,9,2,'4896837097',0,'','189.4.83.123','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 03:52:23','','0000-00-00 00:00:00','br'),(151,9,1,'jose1011_10@hotmail.com',1,'','189.4.83.123','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 03:53:32','','2014-12-07 03:53:32','br'),(152,5,2,'4899315933',0,'','187.58.206.44','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:05:33','','0000-00-00 00:00:00','br'),(153,5,2,'4899315933',0,'','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:10:30','','0000-00-00 00:00:00','br'),(154,5,2,'4899973992',0,'','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:10:47','','0000-00-00 00:00:00','br'),(155,5,2,'4899973992',0,'','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:10:49','','0000-00-00 00:00:00','br'),(156,5,2,'4899973992',0,'','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:11:10','','0000-00-00 00:00:00','br'),(157,5,2,'4896662015',0,'','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:11:50','','0000-00-00 00:00:00','br'),(158,5,2,'4888234804',1,'2;Usuario sem Creditos','177.16.205.107','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:11:58','','2014-12-07 07:11:58','br'),(159,10,2,'4891651919',0,'','187.112.74.68','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 07:19:41','','0000-00-00 00:00:00','br'),(160,5,2,'4899315933',0,'','187.112.74.68','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 08:30:39','','0000-00-00 00:00:00','br'),(161,5,2,'4899973992',0,'','187.112.74.68','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 08:30:46','','0000-00-00 00:00:00','br'),(162,5,2,'4899973992',0,'','187.112.74.68','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 08:30:48','','0000-00-00 00:00:00','br'),(163,3,1,'mdiluccio@gmail.com',1,'','187.5.172.175','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-08 04:18:52','','2014-12-08 04:18:52','br'),(164,10,1,'hbt.vieira@gmail.com',1,'','177.132.169.217','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 23:42:23','','2014-12-07 23:42:23','br'),(165,10,1,'hbt.vieira@gmail.com',0,'','177.132.169.217','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 23:42:34','','0000-00-00 00:00:00','br'),(166,8,2,'4896662015',1,'000:MPG000087788863-73498198','177.132.169.217','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 23:42:47','','2014-12-07 23:42:48','br'),(167,10,2,'4896662015',0,'','177.132.169.217','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-07 23:44:57','','0000-00-00 00:00:00','br'),(168,11,1,'Warnabroda@gmail.com',1,'','177.132.169.217','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; GT-I9300 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-08 00:38:53','','2014-12-08 00:38:54','br'),(169,11,2,'4891640107',1,'000:MPG000087848629-73546066','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-08 18:20:20','','2014-12-08 18:20:20','br');
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

-- Dump completed on 2014-12-08 13:48:27
