package turma

import (
	"strings"
	"turminha/types"
)

type Turma struct {
	ChildrenList []types.Child
}

func (t *Turma) Size() int {
	return len(t.ChildrenList)
}

func (t *Turma) GetFirstNames() []string {
	firstNames := make([]string, t.Size())
	for i, child := range t.ChildrenList {
		firstNames[i] = strings.Split(child.FullName, " ")[0]
	}
	return firstNames
}
