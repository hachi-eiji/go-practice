create database test default charset utf8mb4;
use test;

create table seq (
  id bigint unsigned default 1,
  primary key(id)
) engine=MyISAM;
insert into seq(id)values(100);

create table user (
  id bigint unsigned,
  name varchar(200) not null,
  create_at datetime,
  memo varchar(200),
  use_point bigint unsigned,
  primary key(id)
);

insert into user(id,name,create_at,memo,use_point)values(1,'test',now(),'memo',0);
