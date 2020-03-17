
CREATE DATABASE `test`;

CREATE TABLE `user` (
  `id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `desc` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

INSERT INTO test.`User` (id,name,created_at,updated_at,deleted_at,`desc`) VALUES
('1213','teamlint','2019-12-13 15:00:37.162','2019-12-13 15:00:37.162',NULL,NULL)
,('1214','ardan','2019-12-13 15:01:09.129','2019-12-13 15:01:09.129','2019-12-13 15:01:09.0',NULL)
;
