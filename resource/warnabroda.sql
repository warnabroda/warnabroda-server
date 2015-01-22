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
INSERT INTO `contact_types` VALUES (1,'E-mail','pt-br'),(2,'SMS','pt-br'),(3,'WhatsApp','pt-br');
/*!40000 ALTER TABLE `contact_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ignore_list`
--

DROP TABLE IF EXISTS `ignore_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ignore_list` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Contact` varchar(255) DEFAULT NULL,
  `Ip` varchar(255) DEFAULT NULL,
  `Browser` varchar(255) DEFAULT NULL,
  `Operating_system` varchar(255) DEFAULT NULL,
  `Device` varchar(255) DEFAULT NULL,
  `Raw` varchar(255) DEFAULT NULL,
  `Created_by` varchar(255) DEFAULT NULL,
  `Created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_modified_date` timestamp NULL DEFAULT NULL,
  `confirmed` tinyint(1) NOT NULL DEFAULT '0',
  `confirmation_code` varchar(10) DEFAULT NULL,
  `message` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Contact_UNIQUE` (`Contact`),
  UNIQUE KEY `confirmation_code_UNIQUE` (`confirmation_code`)
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (1,'Você está com Mau Hálito','pt-br'),(2,'Você está com odor de suor','pt-br'),(3,'Você tem Sujeira nos dentes','pt-br'),(4,'Sua menstruação vazou na roupa','pt-br'),(5,'O Seu vaso sanitário está sujo de coco','pt-br'),(7,'Você está fazendo barulho incomodo com a boca','pt-br'),(8,'O barulho dos pés e/ou mãos incomodam','pt-br'),(9,'Você tá com um chulezinho eim','pt-br'),(10,'Tua roupa tá do lado avesso','pt-br'),(11,'Teu cofrinho tá aparecendo','pt-br'),(12,'Tens caca no nariz','pt-br'),(13,'Mas que mesa bagunçada eim','pt-br'),(14,'Sou um(a) adimirador(a) secreto(a)','pt-br'),(15,'Relaxa, você está muito estressado(a)','pt-br'),(16,'Você está com o zipper aberto','pt-br'),(17,'Teu som tá muito alto. Poderia diminuir por favor?','pt-br'),(18,'Você tá muito linda, tá judiando do meu coração','pt-br'),(19,'Você tá um gato, to de olho eim se liga','pt-br'),(20,'Your bad breath is noticeable','en'),(21,'Your sweat odor bothers','en'),(22,'You have got something stuck on your teeth','en'),(23,'Your flow leaked in your clothes','en'),(24,'You have some smelly shit stuck in your toilet','en'),(25,'Do you really need to eat so loud?','en'),(26,'Does your hands and/or feet really need to be this loud?','en'),(27,'Your feet are really smelly','en'),(28,'Your clothes are inside out','en'),(29,'Your butt crack is visible','en'),(30,'You have a booger in your nose','en'),(31,'Your desk is a complete mess','en'),(32,'You have a secret admirer, me =D','en'),(33,'Wow, you are too stressed, try to chill out','en'),(34,'Your zipper is down, dont let it run','en'),(35,'Your music/audio is too loud. Please turn it down','en'),(36,'You are so hot you hurt my heart, you are beautiful','en'),(37,'I Cant take my eyes off of you handsome!','en'),(38,'Usted está con mal aliento','es'),(39,'Su olor del sudor me molesta','es'),(40,'¿Tiene algo atorado en los dientes','es'),(41,'Tiene signo de la menstruación en la ropa','es'),(42,'Usted tiene mierda pegado en su sanitário','es'),(43,'Usted está haciendo ruido molesto con la boca','es'),(44,'Usted está haciendo ruido molesto con los pies u manos','es'),(45,'Usted tiene mal olor en los pies','es'),(46,'Usted está con la ropa al revés','es'),(47,'Tu grieta del extremo está muy visible','es'),(48,'Tu tiene un notable mocos en la nariz','es'),(49,'Tu escritorio está un desastre','es'),(50,'Usted tiene un admirador secreto','es'),(51,'Usted está muy estresado, relajarse','es'),(52,'Tu cremallera está abierta','es'),(53,'Su música/audio es demasiado alto. Por favor, baje el volumen','es'),(54,'Usted es tan caliente que daño a mi corazón, eres muy hermosa','es'),(55,'No puedo tomar mis ojos de ti guapo!','es'),(56,'Please respect, no whistle','en'),(57,'Te acho muito metido(a)','pt-br'),(58,'You are so full of yourself','en'),(59,'Eres tan lleno de ti mismo','es'),(60,'Você tem um odiador(a) secreto(a)','pt-br'),(63,'Você não é um bom chefe','pt-br'),(64,'You are not a good boss','en'),(65,'Usted no es un buen jefe','es'),(66,'Não gostei da sua atitude','pt-br'),(67,'I dont like your attitude','en'),(68,'No me gustó su actitud','es'),(69,'Eu sei que você peidou','pt-br'),(70,'I know you have farted','en'),(71,'Sé que usted ha tirado un pedo','es'),(72,'Use este serviço de aviso anônimo','pt-br'),(73,'Use this anonymous warning service','en'),(74,'Utilice este servicio de aviso anónimo','es');
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
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjects`
--

LOCK TABLES `subjects` WRITE;
/*!40000 ALTER TABLE `subjects` DISABLE KEYS */;
INSERT INTO `subjects` VALUES (1,'Pô parceiro(a)','pt-br'),(2,'Um amigo(a) pediu para avisar','pt-br'),(3,'Só um toque','pt-br'),(4,'Ow, se liga ae','pt-br'),(5,'Olá, venho atraves desse informar','pt-br'),(6,'É para o teu bem','pt-br'),(7,'Quem avisa amigo é','pt-br'),(8,'Não é por mal é só um aviso','pt-br'),(9,'Só estás sendo avisado porque te acham legal','pt-br'),(10,'Só estás sendo avisando porque se importam contigo','pt-br'),(11,'Um amigo acaba de lhe dar um toque','pt-br'),(12,'Los verdaderos amigos te abren los ojos','es'),(13,'Un amigo le advierte suavemente','es'),(15,'Para que lo sepas','es'),(16,'Hey man, escucha me','es'),(17,'Hola, creo que esto puede afectar a usted','es'),(18,'Esto es para su propio bien','es'),(19,'Solo los buenos amigos te abren los ojos','es'),(20,'Yo no quiero ser malo, es sólo un advertirle','es'),(21,'Usted ha sido advertido porque alguien pensará que usted es agradable','es'),(22,'Usted ha sido advertido porque alguien se preocupa por ti','es'),(23,'Un amigo le dio un codazo','es'),(24,'Hey Broda, come on','en'),(25,'A friend gently warns you','en'),(26,'Just so you know','en'),(27,'Hey man, listen up','en'),(28,'Hi, I believe this may concern you','en'),(29,'This is for your own good','en'),(30,'Who warns friend is','en'),(31,'I do not want to be mean, it is just a warn','en'),(32,'You have been warned because someone think you are nice','en'),(33,'You have been warned because someone cares about you','en'),(34,'A friend just poked you','en');
/*!40000 ALTER TABLE `subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `last_login` timestamp NULL DEFAULT NULL,
  `authenticated` tinyint(1) NOT NULL DEFAULT '0',
  `user_hole` varchar(10) NOT NULL DEFAULT 'ROLE_USER',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'ademarizu','ademarizu@gmail.com','a8df37059a2e9915141c92730242a18a9d85c06c','Ademar Izu','2015-01-02 23:04:46',0,'ROLE_ADMIN'),(2,'hbt.vieira','hbt.vieira@gmail.com','ce9cdc7e751033862c9a453e931cbdce11c7852c','Herbert Smith','2015-01-22 16:28:01',0,'ROLE_ADMIN'),(3,'raissaesther','raissaesther@gmail.com','a8df37059a2e9915141c92730242a18a9d85c06c','Raissa Esther','2015-01-02 23:08:31',0,'ROLE_ADMIN');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warnings`
--

DROP TABLE IF EXISTS `warnings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warnings` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
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
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Id_UNIQUE` (`Id`),
  KEY `index_message` (`Id_message`),
  KEY `index_contact_type` (`Id_contact_type`)
) ENGINE=InnoDB AUTO_INCREMENT=59578 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-01-22 15:15:29
