package data

//User Model structs
type User struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Visits []*Visit `json:"visits"`
}

//Visit struct
type Visit struct {
	Id       string `json:"id"`
	Location string `json:"location"`
}

// Mock data
var viewer = &User{
	Id:    "1",
	Name:  "Brooke",
	Email: "brooke@gmail.com",
}
var visits = []*Visit{
	&Visit{"0", "Eureka"},
	&Visit{"1", "Paxtis"},
	&Visit{"2", "Capitol Grill"},
}

//GetUser will return a user given an ID.
func GetUser(id string) *User {
	if id == viewer.Id {
		return viewer
	}
	return nil
}

//GetViewer will get the current user
func GetViewer() *User {
	return viewer
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
