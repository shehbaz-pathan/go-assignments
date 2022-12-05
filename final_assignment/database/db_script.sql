CREATE DATABASE IF NOT EXISTS books;
USE books;
CREATE TABLE IF NOT EXISTS book_info(Isbn bigint NOT NULL,Title varchar(500) NOT NULL, Synopsis varchar(1000) NOT NULL,AuthorName varchar(100) NOT NULL,Price float(6,2) NOT NULL, PRIMARY KEY(Isbn));
CREATE USER 'db-user'@'%' IDENTIFIED BY 'KpkZJvaPFyR13y3';
GRANT ALL PRIVILEGES ON books.book_info to 'db-user'@'%';
