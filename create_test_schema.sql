create database if not exists test default charset utf8mb4;
use test;

create table if not exists seq (
  id bigint unsigned default 100,
  primary key(id)
) engine=MyISAM;
delete from seq;
insert into seq(id)values(100);

create table if not exists user (
  id bigint unsigned,
  name varchar(200) not null,
  create_at datetime,
  memo varchar(200),
  use_point bigint unsigned,
  primary key(id)
);

delete from user;
insert into user(id,name,create_at,memo,use_point)values(1,'test',now(),'memo',0);

create table if not exists user_account (
  id bigint unsigned,
  account_type enum('twitter', 'facebook', 'github') not null,
  account_name varchar(100) not null,
  primary key (id, account_type)
);

delete from user_account;
insert into user_account(id, account_type, account_name)values( 1, 'twitter', 'test_twitter'),(1,'facebook','test_fb');
