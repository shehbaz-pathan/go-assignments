CREATE DATABASE IF NOT EXISTS books;
USE books;
CREATE TABLE IF NOT EXISTS book_info(Isbn int NOT NULL,Title varchar(100) NOT NULL, Synopsis varchar(100) NOT NULL,AuthorName varchar(100) NOT NULL,Price float(6,2) NOT NULL, PRIMARY KEY(Isbn));
CREATE USER 'db-user'@'%' IDENTIFIED BY 'KpkZJvaPFyR13y3';
GRANT ALL PRIVILEGES ON books1.book_info to 'db-user'@'%';