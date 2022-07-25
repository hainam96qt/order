package db

import "context"

const getProductsByIDs = `-- name: GetUserByEmail :many
SELECT id, name, description, price, seller_id FROM products where id in ? 
`

func (q *Queries) GetProductsByIDs(ctx context.Context, arg []int) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByIDs, arg)
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
SELECT id, order_id, product_id, quantity FROM order_product WHERE id in ?
`

func (q *Queries) GetOrderProductByOrderIDs(ctx context.Context, id []int) ([]OrderProduct, error) {
	rows, err := q.db.QueryContext(ctx, getOrderProductByOrderIDs, id)
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
