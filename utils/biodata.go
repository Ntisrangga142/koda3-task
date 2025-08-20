package utils

type School struct {
	Name  string
	Major string
}

type BioData struct {
	Name         string
	Photo        string
	Email        string
	Age          uint
	Phone        string
	MarialStatus bool
	School       [2]School
}
