package main

import (
	"bufio"
	"os"
	"strings"
	"turminha/turma"
	"turminha/types"
)

func main() {
	//data := parseData()
}

func parseData() *turma.Turma {
	turmaFile, err := os.Open("data/turma.csv")
	checkErr(err)

	turmaScanner := bufio.NewScanner(turmaFile)
	turmaScanner.Split(bufio.ScanLines)

	// Skip first since it is a header
	turmaScanner.Scan()

	t := turma.New()

	for turmaScanner.Scan() {
		s := strings.Split(turmaScanner.Text(), ",")
		t.AddChild(types.Child{
			FullName:  s[0],
			BirthDate: s[1],
			Genre:     s[3],
		})
	}

	return t
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
