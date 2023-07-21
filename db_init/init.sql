CREATE DATABASE IF NOT EXISTS urlshortener;
use urlshortener;
CREATE TABLE `shortenedurls` (
	  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
	  `shortcode` char(6) NOT NULL,
	  `longurl` varchar(255) NOT NULL,
	  `active` bool,
	  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	  
	  PRIMARY KEY  (`id`),
	  UNIQUE KEY `long` (`longurl`)
	  UNIQUE KEY `short` (`shortcode`)
);
