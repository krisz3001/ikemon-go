package main

type CardType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (ct CardType) String() string {
	return "[ ID: " + ct.Id + ", Type: " + ct.Name + " ]"
}
