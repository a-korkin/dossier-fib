package models

type Person struct {
	ID         uint32 `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Age        uint8  `json:"age"`
}

type PersonDto struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Age        uint8  `json:"age"`
}
