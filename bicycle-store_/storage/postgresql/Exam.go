package postgresql

import (
	"app/api/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type R_Repo struct {
	db *pgxpool.Pool
}

func NewCodeRepo(db *pgxpool.Pool) *R_Repo {
	return &R_Repo{db: db}
}
func (r R_Repo) GetDate(ctx context.Context, req *models.GiveMe) (res []models.Answer, err error) {
	query := `SELECT
    staffs.first_name || ' ' || staffs.last_name AS "employe",  categories.category_name AS "category",
       products.product_name AS "product",   order_items.quantity AS "quantity",   order_items.list_price * order_items.quantity AS "summ"
FROM orders
         JOIN order_items ON orders.order_id = order_items.order_id
         JOIN products ON order_items.product_id = products.product_id
         JOIN categories ON products.category_id = categories.category_id
         JOIN staffs ON orders.staff_id = staffs.staff_id
WHERE orders.order_date = $1`

	var hh string

	if req.Day == "" {
		dt := time.Now()
		hh = dt.Format("2006-02-01")
	} else {
		hh = req.Day
	}

	fmt.Println("This time -----> ", hh)

	date, error := time.Parse("2006-01-02", hh)
	if error != nil {
		fmt.Println(error)
		return
	}

	rows, err := r.db.Query(ctx, query, date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var s models.Answer
		err = rows.Scan(
			&s.StaffName,
			&s.Category,
			&s.Product,
			&s.Quantity,
			&s.Summ,
		)
		res = append(res, s)
		if err != nil {
			return res, err
		}
	}
	return res, nil

}
