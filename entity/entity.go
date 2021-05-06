/*
Entity.go contains the types relating to user's asset, and item organization for the main user interface, and our database.
Contain Assets, and Unit types maintain the information that is stored for their assets. The Cargo Type mains the connection between all the information within the Entity structure, and provides an interface to access these resources on by our restapi
*/
package entity

import "database/sql"

//The Asset type defines each asset an owner, and maintains all the information for any user's assets
type Asset struct {
	owner     string
	aid       string
	content   []byte
	id_code   string
	timestamp int
	signature []byte
}

//Cargo maintains the connection to our database, and provides that these functions are well defined in a name space
type Cargo struct {
	DB *sql.DB
}

//Unit type is a single asset object, and contains it's own information that can be updated, and maintain transparency on the user end
type Unit struct {
	owner     string
	subtype   *Asset
	uid       string
	NFT       []byte
	item_code string
	timestamp int
	signature string
}

//All assets provides a Query to the Data base and provides an asset type to be passed back to the user, return all assets under a specific user
func (c Cargo) All_Assets() ([]Asset, error) {
	rows, err := c.DB.Query("SELECT * FROM assets")
	if err != nil {
		return nil, err
	}

	var tks []Asset

	for rows.Next() {
		var a Asset

		err := rows.Scan(&a.owner, &a.aid, &a.content, &a.id_code, &a.timestamp, &a.signature)
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

//All unites provides all of the units under a type of asset, returns all units under an asset
func (c Cargo) All_Units() ([]Unit, error) {
	rows, err := c.DB.Query("SELECT * FROM units")
	if err != nil {
		return nil, err
	}
	var unit []Unit

	for rows.Next() {
		var u Unit

		err := rows.Scan(&u.owner, &u.subtype, &u.uid, &u.NFT, &u.item_code, &u.timestamp, &u.signature)
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
