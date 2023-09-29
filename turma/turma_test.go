package turma

import (
	"testing"
)

func TestSortByAge(t *testing.T) {

	children := map[int]Child{
		1: {AgeInMinutes: 3.4},
		3: {AgeInMinutes: 4.2},
		0: {AgeInMinutes: 3.1},
		2: {AgeInMinutes: 3.5},
		4: {AgeInMinutes: 4.7},
	}

	turma := New()
	for _, c := range children {
		turma.AddChild(c)
	}
	turma.SortByAge()
	for i := range turma.ChildrenList {
		if children[i].AgeInMinutes != turma.ChildrenList[i].AgeInMinutes {
			t.Errorf("Got %v, want %v", children[i].AgeInMinutes, turma.ChildrenList[i].AgeInMinutes)
		}
	}
}
