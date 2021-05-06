/*
Package Entity contains all of the Business Logic for our Database, and how the operations that allow manipulation
User.go Contains all of the types required for user authentication
*/
package entity

//User type contains the information needed for authentication, and logging in
type User struct {
	ID        string
	Email     string
	Password  string
	Username  string
	PublicKey string
}

//Profile type allows the extension of Organizations, and Cryptographic Keys for Organizational Wallets
type Profile struct {
	Org  string
	PKey string
}

func create_profile() (*Profile, error) {
	p := &Profile{
		Org:  "University of Calgary",
		PKey: "hash",
	}
	return p, nil
}

//Creating a user is done by inputting a JSON object, and decoding it into strings before inputting creating a User type in the Database
func create_user() (*User, error) {
	k := &User{
		ID:        "Test",
		Email:     "Example@ucalgary.ca",
		Password:  "Hash",
		Username:  "Donk",
		PublicKey: "12345",
	}
	return k, nil
}

//Creates a schema of the Users for querying password hashes
func (c Cargo) user_base() ([]User, error) {
	rows, err := c.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var usr User

		err := rows.Scan(&usr.ID, &usr.Email, &usr.Password, &usr.Username, &usr.PublicKey)
		if err != nil {
			return nil, err
		}

		users = append(users, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
