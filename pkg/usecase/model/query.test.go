package db

import "context"

func (q *Queries) TruncateData(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, "DELETE FROM users WHERE 1=1")
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, "DELETE FROM products WHERE 1=1")
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, "DELETE FROM orders WHERE 1=1")
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, "DELETE FROM order_product WHERE 1=1")
	return err

}
