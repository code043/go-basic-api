create table product (
	id serial primary key,
	product_name varchar(50) not null,
	price numeric(10, 2) not null
);

insert into product (product_name, price) values('Coffee', 20);