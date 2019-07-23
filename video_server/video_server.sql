SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";

SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;

/*!40101 SET NAMES utf8 */;


--
-- 数据库: `video_server`
--
DROP DATABASE IF EXISTS video_server;

CREATE DATABASE video_server CHARSET=UTF8;

USE `video_server`;


-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE users( 
  id  INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,  
  login_name VARCHAR(64) UNIQUE KEY,    
  pwd TEXT NOT NULL
);


--
-- 表的结构 `video_info`
--

CREATE TABLE video_info( 
  id  VARCHAR(64) PRIMARY KEY NOT NULL, 
  author_id INT UNSIGNED, 
  name TEXT,
  display_ctime TEXT,
  create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
);


--
-- 表的结构 `comments`
--

CREATE TABLE comments( 
  id  VARCHAR(64) PRIMARY KEY NOT NULL, 
  video_id VARCHAR(64),
  author_id INT UNSIGNED,  
  content TEXT,
  time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--
-- 表的结构 `sessions`
--

CREATE TABLE sessions( 
  session_id VARCHAR(32) PRIMARY KEY NOT NULL,
  TTL TINYTEXT,
  login_name VARCHAR(64)
);

--
-- 表的结构 `video_del_rec`
--


CREATE TABLE video_del_rec( 
  video_id VARCHAR(64) PRIMARY KEY NOT NULL
);