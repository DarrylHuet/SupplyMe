/*

 */
package main

import (
	"database/sql"
	"fmt"
)

//The basics of what is collected by the user.
type User struct {
	ID        string
	Email     string
	Password  string
	Username  string
	PublicKey string
}

//Our understanding of what is kept between organizations, and the structures of the required details.
type Profile struct {
	Org  string
	PKey string
}

type User_Cert struct {
	DB *sql.DB
}

func create_profile() (*Profile, error) {
	p := &Profile{
		Org:  "University of Calgary",
		PKey: "hash",
	}
	return p, nil
}

//Our Method for making a new user in our Entity Database. Will take inputs of User information in the form of a JSON file, which we unmarshal and provide
func create_user() (*User, error) {
	k := &User{
		ID:        "Test",
		Email:     "Darryl.huet1@ucalgary.ca",
		Password:  "Hash",
		Username:  "Donk",
		PublicKey: "12345",
	}
	return k, nil
}

func (u User_Cert) user_base() ([]User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user []User

	for rows.Next() {
		var user User

		err := rows.Scan{&user.ID, &user.Email, &user.Password, &user.Username, &user.PublicKey}
		if err != nil {
			return nil, err
		}

		user = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return user, nil

}

func main() {
	a, err := create_user()
	b, err2 := create_profile()
	if err != nil || err2 != nil {
		panic("idk man")
	}
	fmt.Printf("%+v\n", *a)
	fmt.Printf("%+v\n", *b)
}
