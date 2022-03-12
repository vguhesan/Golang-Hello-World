package common

type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip5    int
}
