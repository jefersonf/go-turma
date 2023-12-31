package util

import (
	"bufio"
	"os"
	"strings"
	"time"
	"turminha/turma"
)

func ParseData() *turma.Turma {
	turmaFile, err := os.Open("data/sample.csv")
	checkErr(err)

	turmaScanner := bufio.NewScanner(turmaFile)
	turmaScanner.Split(bufio.ScanLines)

	// Skip first since it is the header
	turmaScanner.Scan()

	t := turma.New()

	for turmaScanner.Scan() {
		s := strings.Split(strings.Trim(turmaScanner.Text(), "\"'"), ",")

		fullname := strings.Trim(s[0], "\"'")
		name, _ := turma.GetFirstAndLastName(s[0])

		b, err := time.Parse("02/01/2006", strings.Trim(s[1], "\"'"))
		if err != nil {
			panic(err)
		}
		childAgeInMinutes := time.Since(b).Minutes()
		childAge := childAgeInMinutes / (60 * 24 * 365)
		birthDate := b.Format("Jan 02, 2006")

		var genre string
		s[3] = strings.Trim(s[3], "\"'")
		if strings.Compare(s[3], "M") == 0 {
			genre = "Boy"
		} else if strings.Compare(s[3], "F") == 0 {
			genre = "Girl"
		}

		t.AddChild(turma.Child{
			FullName:  fullname,
			Name:      name,
			BirthDate: birthDate,
			Age:       int(childAge),
			Genre:     genre,
		})
	}

	return t
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
