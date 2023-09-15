package turma

import (
	"fmt"
	"strings"
	"unicode"

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

func (t *Turma) GetFirstNames() []string {
	firstNames := make([]string, t.Size())
	for i, child := range t.ChildrenList {
		firstNames[i] = strings.Split(child.FullName, " ")[0]
	}
	return firstNames
}

func (t *Turma) GetFirstAndLastNames() []string {
	shortNames := make([]string, t.Size())
	for i, child := range t.ChildrenList {
		nameParts := strings.Split(child.FullName, " ")
		firtName := capitalize(nameParts[0])
		lastName := capitalize(nameParts[len(nameParts)-1])
		shortNames[i] = fmt.Sprintf("%s %s", firtName, lastName)
	}
	return shortNames
}

func capitalize(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
