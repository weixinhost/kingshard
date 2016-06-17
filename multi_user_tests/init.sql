
#DROP DATABASE test_1;
#DROP DATABASE test_2;
#DROP DATABASE test_3;
#DROP DATABASE test_4;

CREATE DATABASE test_1;
CREATE DATABASE test_2;
CREATE DATABASE test_3;
CREATE DATABASE test_4;

use test_1;
CREATE TABLE `test_1` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`f1` char(20) DEFAULT NULL,
`f2` char(20) DEFAULT NULL,
`f3` char(20) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


use test_2;
CREATE TABLE `test_2` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`f1` char(20) DEFAULT NULL,
`f2` char(20) DEFAULT NULL,
`f3` char(20) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


use test_3;
CREATE TABLE `test_3` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`f1` char(20) DEFAULT NULL,
`f2` char(20) DEFAULT NULL,
`f3` char(20) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


use test_4;
CREATE TABLE `test_4` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`f1` char(20) DEFAULT NULL,
`f2` char(20) DEFAULT NULL,
`f3` char(20) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

grant all privileges on test_1.* to user_1@'%' identified by  'pass_1';
grant all privileges on test_2.* to user_2@'%' identified by  'pass_2';
grant all privileges on test_3.* to user_3@'%' identified by  'pass_3';
grant all privileges on test_4.* to user_4@'%' identified by  'pass_4';

flush privileges;
