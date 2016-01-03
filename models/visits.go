package models

//Visit struct
type Visit struct {
	Id       string `json:"id"`
	Location string `json:"location"`
}

var visits = []*Visit{
	&Visit{"0", "Eureka"},
	&Visit{"1", "Paxtis"},
	&Visit{"2", "Capitol Grill"},
}

//GetVisit will find a visit for this user by ID
func GetVisit(id string) *Visit {
	for _, visit := range visits {
		if visit.Id == id {
			return visit
		}
	}
	return nil
}

//GetVisits will give all visits
func GetVisits() []*Visit {
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
