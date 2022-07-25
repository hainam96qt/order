create table user_role (
    id int AUTO_INCREMENT primary key,
    role_name VARCHAR(255)
);

create table users (
    id int primary key AUTO_INCREMENT,
    email varchar(25) not null,
    password varchar(255) not null,
    name varchar(255) not null,
    address varchar(255) not null,
    user_role_id int  not null,
    FOREIGN KEY (user_role_id) REFERENCES user_role(id)
);

CREATE UNIQUE INDEX users_email_idx ON users(email);

create table products (
    id int primary key AUTO_INCREMENT,
    name TEXT not null,
    Description TEXT,
    Price bigint  not null,
    seller_id int  not null,
    FOREIGN KEY (seller_id) REFERENCES users(id)
);

create table orders(
    id int primary key AUTO_INCREMENT,
    delivery_source_address TEXT not null,
    delivery_destination_address TEXT not null,
    buyer_id int not null,
    seller_id int not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status enum('pending', 'accepted')  not null,
    total_price BIGINT  not null,
    FOREIGN KEY (buyer_id) REFERENCES users(id)
);

create table order_product (
    id int primary key AUTO_INCREMENT,
    order_id int not null,
    product_id int not null,
    quantity int not null,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

insert into user_role (role_name) value('seller');
insert into user_role (role_name) value('buyer');
