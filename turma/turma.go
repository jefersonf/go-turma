package turma

type Turma struct {
	ChildrenList  []Child
	OldestChild   string
	YoungestChild string
}

type Child struct {
	FullName     string
	Name         string
	Age          int
	AgeInMinutes float64
	BirthDate    string
	Genre        string
}

func New() *Turma {
	return &Turma{}
}

func (t *Turma) AddChild(child Child) {
	t.ChildrenList = append(t.ChildrenList, child)
}

func (t *Turma) Size() int {
	return len(t.ChildrenList)
}
