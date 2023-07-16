CREATE DATABASE IF NOT EXISTS urlshortener;
use urlshortener;
CREATE TABLE `shortenedurls` (
	  `shortcode` int(12) unsigned NOT NULL auto_increment,
	  `longurl` varchar(255) NOT NULL,
	  `active` bool,
	  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	  PRIMARY KEY  (`shortcode`),
	  UNIQUE KEY `long` (`longurl`)
);
