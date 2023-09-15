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
				FullName:  "Maria Silva",
				BirthDate: "10/09/2019",
				Genre:     "F",
			},
			{
				FullName:  "Jose Lima",
				BirthDate: "10/09/2019",
				Genre:     "M",
			},
			{
				FullName:  "Sebastiao de Souza Cariri",
				BirthDate: "10/09/2019",
				Genre:     "M",
			},
		},
	}

	expected := []string{"Maria", "Jose", "Sebastiao"}
	result := turma.GetFirstNames()
	if slices.Compare(result, expected) != 0 {
		t.Errorf("Got %v, want %v", result, expected)
	}
}
