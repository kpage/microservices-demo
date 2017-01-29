package models

import (
	"database/sql"
	"log"

	"github.com/nvellon/hal"
)

type (
	// Top-level list type for HAL format
	AccountResponse struct {
	}

	Account struct {
		Id           int64
		Username     string
		PasswordHash string
	}
)

func (a AccountResponse) GetMap() hal.Entry {
	return hal.Entry{}
}

func (a Account) GetMap() hal.Entry {
	return hal.Entry{
		"id":       a.Id,
		"username": a.Username,
	}
}

func (db *DB) CreateAccount(account *Account) (*Account, error) {
	result, err := db.Exec(
		"INSERT INTO account (username, passwordHash) VALUES (?, ?)",
		account.Username,
		account.PasswordHash,
	)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newAccount := Account{Id: lastInsertID, Username: account.Username}
	return &newAccount, nil
}

func (db *DB) GetAccountByUsername(username string) (*Account, error) {
	row := db.QueryRow("SELECT * FROM account WHERE username = ?", username)
	a := new(Account)
	err := row.Scan(&a.Id, &a.Username, &a.PasswordHash)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
		return nil, nil
	case err != nil:
		return nil, err
	default:
	}
	return a, nil
}

func (db *DB) AllAccounts() ([]*Account, error) {
	rows, err := db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]*Account, 0)
	for rows.Next() {
		a := new(Account)
		err := rows.Scan(&a.Id, &a.Username, &a.PasswordHash)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}
