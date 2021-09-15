create table category (
                          id int not null primary key auto_increment,
                          name varchar(100) not null
)engine InnoDB;

create table merchant (
                          id int not null primary key auto_increment,
                          name varchar(100) not null ,
                          email varchar(100) not null,
                          address varchar(200),
                          rating double default 0
)engine InnoDB;

create table customer (
                          id int not null primary key auto_increment,
                          name varchar(100) not null ,
                          email varchar(100) not null,
                          address varchar(200),
                          phone_number varchar(20)
)engine InnoDB;

create table product (
                         id int not null primary key auto_increment,
                         merchant_id int not null,
                         category_id int not null,
                         name varchar(100) not null,
                         images_path varchar(100),
                         rating double default 0,
                         price double default 0,
                         stock double default 0
)engine InnoDB;

create table orders (
                        id int not null primary key auto_increment,
                        customer_id varchar(100) not null ,
                        total double default 0,
                        payment_method varchar(20),
                        payment_status varchar(20),
                        created_at timestamp default current_timestamp(),
                        confirm_at timestamp,
                        shipping_name varchar(100),
                        shipping_at timestamp,
                        shipping_status varchar(20)
)engine InnoDB;

create table order_details (
                               id int not null primary key auto_increment,
                               order_id varchar(100) not null ,
                               product_id varchar(100) not null,
                               merchant_id varchar(200),
                               price double default 0,
                               quantity double default 0,
                               amount double default 0
)engine InnoDB;