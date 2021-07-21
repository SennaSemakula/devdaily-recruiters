package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type User struct {
	FirstName string
	Age       int
	Id        int
	CreatedAt time.Time
}

type UserService interface {
	InsertUser(u *User) error
	GetUser(name string) (*User, error)
	GetAllUsers() ([]User, error)
}

type DB struct {
	*sql.DB
}

var (
	UserNotFoundError  = errors.New("user not found")
	DuplicateUserError = errors.New("user already exists")
)

func (d *DB) InsertUser(u *User) error {
	exists, err := d.GetUser(u.FirstName)

	if exists != nil {
		return DuplicateUserError
	}

	if err != nil {
		return err
	}

	now := time.Now()

	log.Printf("inserting user %s in developers table\n", u.FirstName)
	_, err = d.Exec(`INSERT INTO developers (first_name, age, created_at) VALUES ($1, $2, $3)`, u.FirstName, u.Age, now)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUser(name string) (*User, error) {
	// TODO: Logic should use email address as a conditional check. First names are not uids.
	var u User
	query := `SELECT id, first_name, age, created_at FROM developers where first_name = $1`

	err := d.QueryRow(query, name).Scan(&u.Id, &u.FirstName, &u.Age, &u.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, UserNotFoundError
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *DB) GetAllUsers() ([]User, error) {
	rows, err := d.Query(`SELECT id, first_name, age, created_at FROM developers`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err = rows.Scan(&u.Id, &u.FirstName, &u.Age, &u.CreatedAt)
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
