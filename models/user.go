package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int
	name string
	pwd string
	email string
}

const QueryEmailStmt = "select * from users where email=$1;"
const InsertNewUserStmt = "insert into users (name, email, pwd) values($1, $2, $3) returning id;"
const UpdateUserStmt = "update users set name=$2, email=$3, pwd=$4 where id=$1 returning id;"

func (u *User) SaveNewUser(db *sql.DB) error {
	row := db.QueryRow(QueryEmailStmt, u.email)
	if row.Scan() != sql.ErrNoRows {
		return errors.New("email already registered")
	}

	//hash new password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.pwd), 10)
	if err != nil {
		return err
	}

	dbErr := db.QueryRow(InsertNewUserStmt, u.name, u.email, hash).Scan(&u.id)
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func (u *User) SaveExistingUser(db *sql.DB) error {
	//hash new password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.pwd), 10)
	if err != nil {
		return err
	}

	dbErr := db.QueryRow(UpdateUserStmt, u.id, u.name, u.email, hash).Scan()
	if dbErr != nil {
		return dbErr
	}
	return nil
}