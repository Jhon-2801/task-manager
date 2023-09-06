create database task_manager;
use task_manager;

create table users(
    id int not null,
    name varchar(45) not null,
    mail varchar(255) not null,
    password varchar(255) not null,
    primary key(id)
)

create table tasks(
    id int not null,
    name varchar(45) not null,
    description varchar(255) not null,
    dates date not null,
    user_id int not null,
    primary key(id),
    foreign key (user_id) references users(id)
)