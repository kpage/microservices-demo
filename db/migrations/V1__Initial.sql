
create table rborder (
	id bigint(20) not null AUTO_INCREMENT primary key,
	location  int(10),	
	ordered_date  datetime,
	customer_name varchar(255),
	status  int(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table line_item (
	id bigint(20) not null AUTO_INCREMENT primary key,
	version int(10),
	milk int(10),
	name varchar(255),
	price varchar(255),
	quantity int(10) not null,
	size int(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table rborder_line_item (
	rborder_id bigint(20) not null,
	line_item_id bigint(20) not null,
	constraint PK_rborder_line_item primary key (rborder_id, line_item_id),
	foreign key(rborder_id) references rborder(id),
	foreign key(line_item_id) references line_item(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into rborder (location, ordered_date, customer_name, status) values (0, '2017-01-15 10:34:50', 'Steve', 0);
SET @rborder_id = last_insert_id();
insert into line_item (version, milk, name, price, quantity, size) values (0, 1, 'Cappuchino', 'EUR 4.2', 1, 1);
SET @line_item_id = last_insert_id();
insert into rborder_line_item values (@rborder_id, @line_item_id);

insert into rborder (location, ordered_date, customer_name, status) values (0, '2017-01-15 10:36:00', 'Larry', 0);
SET @rborder_id = last_insert_id();
insert into line_item (version, milk, name, price, quantity, size) values (0, 1, 'Cappuchino', 'EUR 4.2', 1, 1);
SET @line_item_id = last_insert_id();
insert into rborder_line_item values (@rborder_id, @line_item_id);