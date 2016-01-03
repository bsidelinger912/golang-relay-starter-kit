package models

//Shopper is a mystery shopper
type Shopper struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Visits []*Visit `json:"visits"`
}

//AllShoppers goes to the DB and gets everything
func allShoppers(db *DB) ([]*Shopper, error) {
	rows, err := db.Query("SELECT id, name, email FROM shoppers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shoppers := make([]*Shopper, 0)
	for rows.Next() {
		shopper := new(Shopper)
		err := rows.Scan(&shopper.Id, &shopper.Name, &shopper.Email)
		if err != nil {
			return nil, err
		}
		shoppers = append(shoppers, shopper)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return shoppers, nil
}

//GetShopper returns a shopper by ID
func (db *DB) GetShopper(id string) *Shopper {
	shoppers, _ := allShoppers(db)
	for _, shopper := range shoppers {
		if shopper.Id == id {
			return shopper
		}
	}
	return nil
}
