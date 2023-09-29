package turma

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

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

func (t *Turma) SortByAge() {
	slices.SortFunc(t.ChildrenList, func(a, b Child) int {
		if a.AgeInMinutes > b.AgeInMinutes {
			return 1
		}
		return -1
	})
}

func (t *Turma) GetOldest() (*Child, error) {
	child, err := t.getChildAt(t.Size() - 1)
	if err != nil {
		return child, fmt.Errorf("impossible to retrieve oldest child: %w", err)
	}
	return child, nil
}

func (t *Turma) GetYoungest() (*Child, error) {
	child, err := t.getChildAt(0)
	if err != nil {
		return child, fmt.Errorf("impossible to retrieve youngest child: %w", err)
	}
	return child, nil
}

func (t *Turma) getChildAt(index int) (*Child, error) {
	if t.Size() == 0 {
		return &Child{}, errors.New("no children")
	}
	if index >= t.Size() {
		return &Child{}, errors.New("invalid position")
	}
	t.SortByAge()
	child := new(Child)
	child.FullName = t.ChildrenList[index].FullName
	child.Name = t.ChildrenList[index].Name
	child.Age = t.ChildrenList[index].Age
	child.AgeInMinutes = t.ChildrenList[index].AgeInMinutes
	child.BirthDate = t.ChildrenList[index].BirthDate
	child.Genre = t.ChildrenList[index].Genre
	return child, nil
}
