package types

type Turma struct {
	Children []Child
}

type Child struct {
	FullName  string
	BirthDate string
	Genre     string
}
