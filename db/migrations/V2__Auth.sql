-- TODO: maybe separate the auth persistence stuff into a separate database?
create table account (
	id bigint(20) not null AUTO_INCREMENT primary key,
	username varchar(255) not null,
    passwordHash varchar(255),
    constraint UQ_account_username unique (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert ignore into account (username, passwordHash) values ('kevin', '$s2$16384$8$1$KqC/qkUqs3DK1yv2PdOPjvr3$p5phFQAVaY9KPlt/QAYbTiOEiSMNLbe7S+0dGjZcEcg=');