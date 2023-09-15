package turma

import (
	"testing"

	"turminha/types"

	"golang.org/x/exp/slices"
)

func TestGeTFirstAndLastNames(t *testing.T) {

	turma := Turma{
		ChildrenList: []types.Child{
			{
				FullName: "Maria Nascimento silva",
			},
			{
				FullName: "Jose Lima",
			},
			{
				FullName: "Sebastiao de Souza cariri",
			},
		},
	}

	expected := []string{"Maria Silva", "Jose Lima", "Sebastiao Cariri"}
	result := make([]string, turma.Size())
	for i, child := range turma.ChildrenList {
		result[i], _ = GetFirstAndLastName(child.FullName)
	}
	if slices.Compare(result, expected) != 0 {
		t.Errorf("Got %v, want %v", result, expected)
	}
}
