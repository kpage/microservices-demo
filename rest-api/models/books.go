package models

import "github.com/nvellon/hal"

type (
	// Top-level list type for HAL format
	BookResponse struct {
	}

	Book struct {
		Isbn   string
		Title  string
		Author string
		Price  float32
	}
)

func (p BookResponse) GetMap() hal.Entry {
	return hal.Entry{}
}

func (b Book) GetMap() hal.Entry {
	return hal.Entry{
		"isbn":   b.Isbn,
		"title":  b.Title,
		"author": b.Author,
		"price":  b.Price,
	}
}

func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
