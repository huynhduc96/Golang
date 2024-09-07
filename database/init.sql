
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

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `golang` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `golang`;
DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `address` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,'Marry','marry','123456','London'),
(2,'Bob','bob','123456','Hanoi'),
(3,'Alice','alice','123456','New York'),
(4,'John','john','123456','Paris'),
(5,'Jane','jane','123456','Tokyo'),
(6,'David','david','123456','Sydney'),
(7,'Emily','emily','123456','Berlin'),
(8,'Michael','michael','123456','Rome'),
(9,'Sophia','sophia','123456','Los Angeles'),
(10,'William','william','123456','San Francisco'),
(11,'Olivia','olivia','123456','Chicago'),
(12,'Ethan','ethan','123456','Toronto'),
(13,'Ava','ava','123456','Dubai'),
(14,'Noah','noah','123456','Hong Kong'),
(15,'Isabella','isabella','123456','Madrid'),
(16,'Liam','liam','123456','Seoul'),
(17,'Mia','mia','123456','Mumbai'),
(18,'Jacob','jacob','123456','Istanbul'),
(19,'Charlotte','charlotte','123456','Singapore'),
(20,'Mason','mason','123456','Bangkok'),
(21,'Amelia','amelia','123456','Amsterdam'),
(22,'Aiden','aiden','123456','Melbourne'),
(23,'Harper','harper','123456','Rio de Janeiro'),
(24,'Evelyn','evelyn','123456','Moscow'),
(25,'Benjamin','benjamin','123456','Cairo'),
(26,'Abigail','abigail','123456','Buenos Aires'),
(27,'Lucas','lucas','123456','Kuala Lumpur'),
(28,'Sofia','sofia','123456','Cape Town'),
(29,'Elijah','elijah','123456','Mexico City'),
(30,'Scarlett','scarlett','123456','Jakarta');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

