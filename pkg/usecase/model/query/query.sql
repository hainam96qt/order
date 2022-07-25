/* name: GetRoleByID :one */
SELECT * FROM user_role where id = ? LIMIT 1;

/* name: GetUserByEmail :one */
SELECT * FROM users where email = ? LIMIT 1;

/* name: GetUserByID :one */
SELECT * FROM users where id = ? LIMIT 1;


/* name: CreateUser :execresult */
INSERT INTO users (
    id,
    email,
    password,
    name,
    address,
    user_role_id
) VALUES (
    ?,?,?,?,?,?
);

create table products (
    id int primary key AUTO_INCREMENT,
    name TEXT not null,
    Description TEXT,
    Price bigint  not null,
    seller_id int  not null,
    FOREIGN KEY (seller_id) REFERENCES users(id)
);


/* name: CreateProduct :execresult */
INSERT INTO products (
    id,
    name,
    Description,
    price,
    seller_id
) VALUES (
    ?,?,?,?,?
);

/* name: CreateOrder :execresult */
INSERT INTO orders (
    id ,
    delivery_source_address,
    delivery_destination_address,
    buyer_id,
    seller_id,
    status,
    total_price
) VALUES (
    ?,?,?,?,?,?,?
);

/* name: CreateOrderProduct :execresult */
INSERT INTO order_product (
    id,
    order_id,
    product_id,
    quantity
) VALUES (
    ?,?,?,?
);

/* name: GetProductBySeller :many */
SELECT * FROM products where seller_id = ? LIMIT ? OFFSET ?;

/* name: GetNewestProduct :one */
SELECT * FROM products where seller_id = ? and name = ? and Price = ? order by id desc;

/* name: GetProductByID :one */
SELECT * FROM products WHERE id = ?;

/* name: GetOrderByID :one */
SELECT * FROM orders WHERE id = ?;

/* name: GetOrderBySellerID :many */
SELECT * FROM orders WHERE seller_id = ? LIMIT ? offset ?;

/* name: GetOrderByBuyerID :many */
SELECT * FROM orders WHERE buyer_id = ? LIMIT ? offset ?;

/* name: UpdateOrderStatus :exec */
UPDATE orders set status = 'accepted' where seller_id = ? and id = ?;
