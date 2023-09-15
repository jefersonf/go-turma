package turma

import (
	"testing"

	"turminha/types"

	"golang.org/x/exp/slices"
)

func TestGeTFirstNames(t *testing.T) {

	turma := Turma{
		ChildrenList: []types.Child{
			{
				FullName: "Maria Silva",
			},
			{
				FullName: "Jose Lima",
			},
			{
				FullName: "Sebastiao de Souza Cariri",
			},
		},
	}

	expected := []string{"Maria", "Jose", "Sebastiao"}
	result := turma.GetFirstNames()
	if slices.Compare(result, expected) != 0 {
		t.Errorf("Got %v, want %v", result, expected)
	}
}

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
	result := turma.GetFirstAndLastNames()
	if slices.Compare(result, expected) != 0 {
		t.Errorf("Got %v, want %v", result, expected)
	}
}
