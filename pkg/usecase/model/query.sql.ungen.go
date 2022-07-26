package db

import (
	"context"
	"fmt"
	"strconv"
)

const getProductsByIDs = `-- name: GetProductsByIDs :many
SELECT id, name, description, price, seller_id 
FROM products 
WHERE id in %s
`

func (q *Queries) GetProductsByIDs(ctx context.Context, arg []int) ([]Product, error) {
	sql := fmt.Sprintf(getProductsByIDs, arrayInt(arg))
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.SellerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderProductByOrderIDs = `-- name: GetOrderProductByOrderIDs :many
SELECT id, order_id, product_id, quantity FROM order_product WHERE order_id in %s
`

func (q *Queries) GetOrderProductByOrderIDs(ctx context.Context, arg []int) ([]OrderProduct, error) {
	sql := fmt.Sprintf(getOrderProductByOrderIDs, arrayInt(arg))
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderProduct
	for rows.Next() {
		var i OrderProduct
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func arrayInt(arg []int) string {
	value := "("
	for k, v := range arg {
		value += strconv.Itoa(v)
		if k != len(arg)-1 {
			value += ","
		}
	}
	value += ")"
	return value
}
