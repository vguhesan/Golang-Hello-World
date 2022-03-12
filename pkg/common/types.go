package common

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Address   struct {
// 		Street1 string
// 		Street2 string
// 		City    string
// 		State   string
// 		Zip5    int
// 	}
// }

// Alternative
type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Zip5    int // Placing the zip up top creates a compact object compared to placing it down below
	Street1 string
	Street2 string
	City    string
	State   string
}
