package turma

import (
	"testing"
)

func TestSortByAge(t *testing.T) {

	children := map[int]Child{
		1: {BirthDate: "Jan 12, 2020"},
		2: {BirthDate: "Oct 11, 2019"},
		3: {BirthDate: "Mar 23, 2019"},
		0: {BirthDate: "Sep 11, 2020"},
		4: {BirthDate: "Feb 01, 2019"},
	}

	turma := New()
	for _, c := range children {
		turma.AddChild(c)
	}
	turma.SortByAge()
	for i, c := range turma.ChildrenList {
		if c.AgeInMinutes() != turma.ChildrenList[i].AgeInMinutes() {
			t.Errorf("Got %v, want %v", c.AgeInMinutes(), turma.ChildrenList[i].AgeInMinutes())
		}
	}
}

func TestGetOldest(t *testing.T) {

	children := map[int]Child{
		1: {BirthDate: "Jan 12, 2020"},
		2: {BirthDate: "Oct 11, 2019"},
		3: {BirthDate: "Mar 23, 2019"},
		0: {BirthDate: "Sep 11, 2020"},
		4: {BirthDate: "Feb 01, 2019"},
	}

	turma := New()
	for _, c := range children {
		turma.AddChild(c)
	}

	oldestChild, err := turma.GetOldest()
	if err != nil {
		t.Errorf("Got err, want nil")
	}
	oldestedOnList := children[4]
	if oldestChild.AgeInMinutes() != oldestedOnList.AgeInMinutes() {
		t.Errorf("Got %v, want %v", oldestChild.AgeInMinutes(), oldestedOnList.AgeInMinutes())
	}
}
