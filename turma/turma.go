package turma

import (
	"turminha/types"
)

type Turma struct {
	ChildrenList []types.Child
}

func New() *Turma {
	return &Turma{}
}

func (t *Turma) AddChild(child types.Child) {
	t.ChildrenList = append(t.ChildrenList, child)
}

func (t *Turma) Size() int {
	return len(t.ChildrenList)
}
