package turma

import (
	"fmt"
	"strconv"
	"time"
)

type Child struct {
	FullName  string
	Name      string
	Age       int
	BirthDate string
	Genre     string
}

func (c *Child) AgeInMinutes() float64 {
	b, err := time.Parse("Jan 02, 2006", c.BirthDate)
	if err != nil {
		fmt.Printf("unable to parse birth date: %v\n", err)
	}
	toMinutes := time.Since(b).Minutes()
	ageInMinutes, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", toMinutes), 64)
	return ageInMinutes
}
