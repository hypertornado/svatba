CREATE DATABASE svatba CHARACTER SET utf8 DEFAULT COLLATE utf8_unicode_ci;
CREATE USER 'svatba'@'localhost' IDENTIFIED BY 'svatba85';
GRANT ALL ON svatba.* TO 'svatba'@'localhost';
FLUSH PRIVILEGES;