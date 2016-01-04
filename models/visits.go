package models

//Visit struct
type Visit struct {
	Id           string `json:"id"`
	LocationName string `json:"location"`
}

//GetVisit will find a visit for this user by ID
func (db *DB) GetVisit(id string) *Visit {
	stmt, err := db.Prepare("SELECT id, location_name FROM visits WHERE id=$1")
	rows, err := stmt.Query(id)
	if err != nil {
		return nil
	}
	defer rows.Close()

	visit := new(Visit)
	for rows.Next() {
		err := rows.Scan(&visit.Id, &visit.LocationName)
		if err != nil {
			return nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return visit
}

//GetVisits will give all visits
func (db *DB) GetVisits() []*Visit {
	//TODO figure how to get this from connection args
	shopperID := "2"

	stmt, err := db.Prepare("SELECT id, location_name FROM visits WHERE shopper_id=$1")
	rows, err := stmt.Query(shopperID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	visits := make([]*Visit, 0)
	for rows.Next() {
		visit := new(Visit)
		err := rows.Scan(&visit.Id, &visit.LocationName)
		if err != nil {
			return nil
		}
		visits = append(visits, visit)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return visits
}

//VisitsToInterfaceSlice casts the visits to interfaces
func VisitsToInterfaceSlice(visits ...*Visit) []interface{} {
	var interfaceSlice = make([]interface{}, len(visits))
	for i, d := range visits {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
