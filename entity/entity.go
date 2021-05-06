package main

import "database/sql"

type Asset struct {
	owner     string
	aid       string
	content   []string
	timestamp int
	signature string
}

type Cargo struct {
	DB *sql.DB
}

type Unit struct {
	owner         string
	uid           int
	NFT           []byte
	creation_time int
	timestamp     int
	signature     []byte
}

func (c Cargo) All_Assets() ([]Asset, error) {
	rows, err := c.DB.Query("SELECT * FROM assets")
	if err != nil {
		return nil, err
	}

	var tks []Asset

	for rows.Next() {
		var a Asset

		err := rows.Scan(&a.owner, &a.aid, &a.content, &a.timestamp, &a.signature)
		if err != nil {
			return nil, err
		}

		tks = append(tks, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tks, nil
}

func (c Cargo) All_Units() ([]Unit, error) {
	rows, err := c.DB.Query("SELECT * FROM units")
	if err != nil {
		return nil, err
	}
	var unit []Unit

	for rows.Next() {
		var u Unit

		err := rows.Scan(&u.owner, &u.uid, &u.NFT, &u.creation_time, &u.timestamp, &u.signature)
		if err != nil {
			return nil, err
		}

		unit = append(unit, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return unit, nil
}
