package app

type User struct {
	Identifier 	int
	Name       	string
	IsMachine	bool

	AuthDomain 	string
	AuthEntryId int
}