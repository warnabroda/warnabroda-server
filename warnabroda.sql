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
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warnings`
--

LOCK TABLES `warnings` WRITE;
/*!40000 ALTER TABLE `warnings` DISABLE KEYS */;
INSERT INTO warnings VALUES (1,0,1,'',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00','','0000-00-00',''),(2,1,1,'teste',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00','','0000-00-00',''),(3,6,3,'hbtsmith',0,'','179.187.92.218','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','','0000-00-00','','0000-00-00',''),(4,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:04:27.564643','','0000-00-00','br'),(5,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:19:28.745325','','0000-00-00','br'),(6,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:20:17.147294','','0000-00-00','br'),(7,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:24:23.326356','','0000-00-00','br'),(8,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:24:42.248926','','0000-00-00','br'),(9,1,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:57:21.759218','','0000-00-00','br'),(10,2,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 00:59:25.759712','','2014-11-28 00:59:26.253828','br'),(11,5,1,'hbt.vieira@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36','system','2014-11-28 01:02:02.748533','','2014-11-28 01:02:03.166037','br'),(12,1,1,'herbert.silva@jexperts.com.br',0,'','177.16.146.153','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-02 01:37:36.780776','','0000-00-00','br'),(13,3,1,'Adrieli.daltoe@gmail.com',0,'','179.207.178.19','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; GT-I9300 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-02 17:34:01.946108','','0000-00-00','br'),(14,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:08:52.295077','','0000-00-00','br'),(15,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:01.951865','','0000-00-00','br'),(16,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:03.428617','','0000-00-00','br'),(17,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:09:17.3123','','0000-00-00','br'),(18,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:11:27.591173','','0000-00-00','br'),(19,9,1,'ademarizu@gmail.com',0,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:12:52.720343','','0000-00-00','br'),(20,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:11.856095','','2014-12-04 22:14:12.381011','br'),(21,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:13.519499','','2014-12-04 22:14:13.870665','br'),(22,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:22.32111','','2014-12-04 22:14:22.714811','br'),(23,9,1,'ademarizu@gmail.com',1,'','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:14:22.870497','','2014-12-04 22:14:23.265625','br'),(28,4,2,'4891917810',1,'6;36073798','191.191.43.121','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-04 22:28:25.590279','','2014-12-04 22:28:26.089224','br'),(29,6,2,'4896662015',1,'6;36084625','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:13:40.268972','','2014-12-05 12:13:40.940085','br'),(30,6,2,'4891640107',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:14:14.439585','','0000-00-00','br'),(31,6,1,'tjvargas2@gmail.com',1,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:14:45.694648','','2014-12-05 12:14:46.234828','br'),(32,6,2,'4891188080',1,'6;36084637','186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.104 Safari/537.36','system','2014-12-05 12:14:58.548671','','2014-12-05 12:14:58.866108','br'),(33,4,1,'vinifritzen@gmail.com',1,'','192.168.1.81, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:17:15.381622','','2014-12-05 12:17:16.3563','br'),(34,8,2,'4896004128',1,'6;36084674','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:17:53.437374','','2014-12-05 12:17:54.076505','br'),(35,1,1,'ramoncordini@gmail.com',1,'','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:19:28.383591','','2014-12-05 12:19:28.853291','br'),(36,9,1,'graziela.souza@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:20:00.931013','','2014-12-05 12:20:01.358725','br'),(37,3,1,'fasebastiani@gmail.com',1,'','189.90.61.78','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:22:54.138844','','2014-12-05 12:22:54.612216','br'),(38,9,1,'ademarizu@gmail.com',1,'','189.101.247.92','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:24:42.655943','','2014-12-05 12:24:43.092124','br'),(39,6,2,'4884111635',1,'6;36084834','189.101.247.92','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:25:42.563084','','2014-12-05 12:25:42.744259','br'),(40,1,1,'phelipewinter@gmail.com',1,'','186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.104 Safari/537.36','system','2014-12-05 12:27:45.846754','','2014-12-05 12:27:46.312581','br'),(41,3,1,'mayara.souza@jexperts.com.br',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:28:00.632829','','2014-12-05 12:28:00.99445','br'),(42,8,1,'mayara@jexperts.com.bt',1,'','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:28:05.875512','','2014-12-05 12:28:06.295594','br'),(43,8,1,'mayara.souza@jexperts.com.br',1,'','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:29:25.8777','','2014-12-05 12:29:26.314424','br'),(44,2,1,'leonardo@ahgora.com.br',1,'','189.101.215.166','chrome','android','android','Mozilla/5.0 (Linux; Android 4.1.2; XT920 Build/2_330_2009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-05 12:30:40.91983','','2014-12-05 12:30:41.389234','br'),(45,3,2,'4888634777',0,'','192.168.1.216, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36','system','2014-12-05 12:30:45.277817','','0000-00-00','br'),(46,7,2,'4799698083',1,'6;36085135','60.225.193.247','chrome','android','android','Mozilla/5.0 (Linux; Android 5.0; Nexus 5 Build/LRX21O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36','system','2014-12-05 12:31:13.62631','','2014-12-05 12:31:13.942776','br'),(47,9,1,'pauladesouza1604@gmail.com',1,'','192.168.1.113, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:31:44.99735','','2014-12-05 12:31:45.419115','br'),(48,8,1,'fernanda.barbosa@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:33:50.655507','','2014-12-05 12:33:51.113977','br'),(49,8,1,'kamilla.pires@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:34:06.479346','','2014-12-05 12:34:07.03329','br'),(50,9,1,'ricardojp84@gmail.com',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:34:28.189191','','2014-12-05 12:34:28.548707','br'),(51,10,2,'4891651919',1,'6;36085322','192.168.1.80, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:37:20.848532','','2014-12-05 12:37:21.177131','br'),(52,1,1,'daniela.munari@fiesc.com.br',1,'','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:40:49.259897','','2014-12-05 12:40:49.672771','br'),(53,10,1,'sec.rosimeri@gmail.com',1,'','192.168.1.113, 186.215.116.241','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:41:25.513028','','2014-12-05 12:41:25.95858','br'),(54,10,1,'jabez_t.i@hotmail.com',1,'','177.207.221.70','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64; rv:33.0) Gecko/20100101 Firefox/33.0','system','2014-12-05 12:43:34.957766','','2014-12-05 12:43:35.417384','br'),(55,10,1,'daniel.cherem@fiesc.com.br',1,'','177.221.55.254','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; rv:34.0) Gecko/20100101 Firefox/34.0','system','2014-12-05 12:44:00.477659','','2014-12-05 12:44:00.889084','br'),(56,9,2,'4899312065',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:04.519038','','0000-00-00','br'),(57,3,1,'adrieli.daltoe@gmail.com',1,'','192.168.1.195, 186.215.116.241','chrome','linux','unknown','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:13.815392','','2014-12-05 12:44:14.191885','br'),(58,9,2,'4896852019',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:29.348774','','0000-00-00','br'),(59,9,2,'4896852019',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:44:35.116775','','0000-00-00','br'),(60,10,1,'jabez_t.i@hotmail.com',1,'','177.207.221.70','firefox','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64; rv:33.0) Gecko/20100101 Firefox/33.0','system','2014-12-05 12:45:44.335751','','2014-12-05 12:45:44.796275','br'),(61,9,1,'lecogoulart@bol.com.br',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 12:45:54.498188','','2014-12-05 12:45:54.988359','br'),(62,10,2,'4891640107',0,'','177.207.221.70','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:46:09.209635','','0000-00-00','br'),(63,3,1,'rafael.nishihora@gmail.com',1,'','150.162.81.101','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:58:30.134351','','2014-12-05 12:58:30.555199','br'),(64,3,2,'4899211951',1,'6;36087095','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:32.134835','','2014-12-05 12:59:32.586258','br'),(65,3,2,'4891352028',1,'6;36087099','150.162.81.101','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:35.793465','','2014-12-05 12:59:36.089106','br'),(66,2,2,'4899895468',0,'','10.180.208.137, 177.221.48.74','chrome','windows','unknown','Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 12:59:46.561089','','0000-00-00','br'),(67,9,1,'Contepraju@gmail',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 13:22:33.254409','','2014-12-05 13:22:33.538597','br'),(68,4,1,'contepraju@gmail.com',1,'','189.8.253.109','chrome','mac','unknown','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36','system','2014-12-05 13:24:45.564791','','2014-12-05 13:24:46.053973','br'),(69,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:23.57897','','2014-12-05 13:26:23.99214','br'),(70,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:23.818628','','2014-12-05 13:26:24.24068','br'),(71,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:43.337668','','2014-12-05 13:26:43.769158','br'),(72,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:26:43.537871','','2014-12-05 13:26:44.069501','br'),(73,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:27:09.259178','','2014-12-05 13:27:09.671166','br'),(74,3,1,'le_fazzani@hotmail.com',1,'','191.169.252.111','chrome','android','android','Mozilla/5.0 (Linux; Android 4.4.4; XT1058 Build/KXA21.12-L1.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36','system','2014-12-05 13:27:09.537793','','2014-12-05 13:27:10.144381','br');
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

-- Dump completed on 2014-12-03 17:40:28
