drop database if exists Ryo;

create database Ryo;

Use Ryo;

drop table if exists t_user;
create table t_user(
	userId integer not null auto_increment,
	userAccount varchar(255) not null,
	userPwd varchar(255) not null,
	primary key(userId)
)engine=innodb default charset=utf8;

drop table if exists t_book;

create table t_book(
	bookId integer not null auto_increment,
	bookCode varchar(64) not null,
	bookName varchar(64) not null,
	userId integer not null,
	primary key(bookId)
) engine=innodb default charset=utf8;
	
