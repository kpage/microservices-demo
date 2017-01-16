package models

import (
	"time"

	"github.com/nvellon/hal"
)

type (
	// Top-level list type for HAL format
	OrderResponse struct {
	}

	Order struct {
		Id           int64
		Location     int32 // TODO: use some kind of enum
		OrderedDate  time.Time
		CustomerName string
		Status       int32 // TODO: use some kind of enum
	}
)

func (o OrderResponse) GetMap() hal.Entry {
	return hal.Entry{}
}

func (o Order) GetMap() hal.Entry {
	return hal.Entry{
		"id":           o.Id,
		"location":     o.Location,
		"orderedDate":  o.OrderedDate,
		"customerName": o.CustomerName,
		"status":       o.Status,
	}
}

func (db *DB) AllOrders() ([]*Order, error) {
	rows, err := db.Query("SELECT * FROM rborder")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*Order, 0)
	for rows.Next() {
		o := new(Order)
		err := rows.Scan(&o.Id, &o.Location, &o.OrderedDate, &o.CustomerName, &o.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
